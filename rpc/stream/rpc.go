package stream

import (
	"context"
	"fmt"
	"sync"

	"github.com/cometbft/cometbft/libs/log"
	tmquery "github.com/cometbft/cometbft/libs/pubsub/query"
	rpcclient "github.com/cometbft/cometbft/rpc/client"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	tmtypes "github.com/cometbft/cometbft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/evmos/ethermint/rpc/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

const (
	streamSubscriberName = "ethermint-json-rpc"
	subscribBufferSize   = 1024

	headerStreamSegmentSize = 128
	headerStreamCapacity    = 128 * 32
	txStreamSegmentSize     = 1024
	txStreamCapacity        = 1024 * 32
	logStreamSegmentSize    = 2048
	logStreamCapacity       = 2048 * 32
)

var (
	evmEvents = tmquery.MustParse(fmt.Sprintf("%s='%s' AND %s.%s='%s'",
		tmtypes.EventTypeKey,
		tmtypes.EventTx,
		sdk.EventTypeMessage,
		sdk.AttributeKeyModule, evmtypes.ModuleName)).String()
	headerEvents = tmtypes.QueryForEvent(tmtypes.EventNewBlockHeader).String()
	evmTxHashKey = fmt.Sprintf("%s.%s", evmtypes.TypeMsgEthereumTx, evmtypes.AttributeKeyEthereumTxHash)
)

type RPCHeader struct {
	EthHeader *ethtypes.Header
	Hash      common.Hash
}

// RPCStream provides data streams for newHeads, logs, and pendingTransactions.
type RPCStream struct {
	evtClient rpcclient.EventsClient
	logger    log.Logger
	txDecoder sdk.TxDecoder

	headerStream    *Stream[RPCHeader]
	txStream        *Stream[common.Hash]
	pendingTxStream *Stream[common.Hash]
	logStream       *Stream[*ethtypes.Log]

	wg sync.WaitGroup
}

func NewRPCStreams(
	evtClient rpcclient.EventsClient,
	logger log.Logger,
	txDecoder sdk.TxDecoder,
) (*RPCStream, error) {
	s := &RPCStream{
		evtClient: evtClient,
		logger:    logger,
		txDecoder: txDecoder,

		headerStream:    NewStream[RPCHeader](headerStreamSegmentSize, headerStreamCapacity),
		txStream:        NewStream[common.Hash](txStreamSegmentSize, txStreamCapacity),
		pendingTxStream: NewStream[common.Hash](txStreamSegmentSize, txStreamCapacity),
		logStream:       NewStream[*ethtypes.Log](logStreamSegmentSize, logStreamCapacity),
	}

	ctx := context.Background()

	chHeaders, err := s.evtClient.Subscribe(ctx, streamSubscriberName, headerEvents, subscribBufferSize)
	if err != nil {
		return nil, err
	}

	chLogs, err := s.evtClient.Subscribe(ctx, streamSubscriberName, evmEvents, subscribBufferSize)
	if err != nil {
		if err := s.evtClient.UnsubscribeAll(context.Background(), streamSubscriberName); err != nil {
			s.logger.Error("failed to unsubscribe", "err", err)
		}
		return nil, err
	}

	go s.start(&s.wg, chHeaders, chLogs)

	return s, nil
}

func (s *RPCStream) Close() error {
	if err := s.evtClient.UnsubscribeAll(context.Background(), streamSubscriberName); err != nil {
		return err
	}
	s.wg.Wait()
	return nil
}

func (s *RPCStream) HeaderStream() *Stream[RPCHeader] {
	return s.headerStream
}

func (s *RPCStream) PendingTxStream() *Stream[common.Hash] {
	return s.pendingTxStream
}

func (s *RPCStream) LogStream() *Stream[*ethtypes.Log] {
	return s.logStream
}

// ListenPendingTx is a callback passed to application to listen for pending transactions in CheckTx.
func (s *RPCStream) ListenPendingTx(hash common.Hash) {
	s.pendingTxStream.Add(hash)
}

func (s *RPCStream) start(
	wg *sync.WaitGroup,
	chHeaders <-chan coretypes.ResultEvent,
	chLogs <-chan coretypes.ResultEvent,
) {
	wg.Add(1)
	defer func() {
		wg.Done()
		if err := s.evtClient.UnsubscribeAll(context.Background(), streamSubscriberName); err != nil {
			s.logger.Error("failed to unsubscribe", "err", err)
		}
	}()

	for {
		select {
		case ev, ok := <-chHeaders:
			if !ok {
				chHeaders = nil
				break
			}

			data, ok := ev.Data.(tmtypes.EventDataNewBlockHeader)
			if !ok {
				s.logger.Error("event data type mismatch", "type", fmt.Sprintf("%T", ev.Data))
				continue
			}

			baseFee := types.BaseFeeFromEvents(data.ResultBeginBlock.Events)

			// TODO: fetch bloom from events
			header := types.EthHeaderFromTendermint(data.Header, ethtypes.Bloom{}, baseFee)
			s.headerStream.Add(RPCHeader{EthHeader: header, Hash: common.BytesToHash(data.Header.Hash())})
		case ev, ok := <-chLogs:
			if !ok {
				chLogs = nil
				break
			}

			if _, ok := ev.Events[evmTxHashKey]; !ok {
				// ignore transaction as it's not from the evm module
				continue
			}

			// get transaction result data
			dataTx, ok := ev.Data.(tmtypes.EventDataTx)
			if !ok {
				s.logger.Error("event data type mismatch", "type", fmt.Sprintf("%T", ev.Data))
				continue
			}
			txLogs, err := evmtypes.DecodeTxLogsFromEvents(dataTx.TxResult.Result.Data, dataTx.TxResult.Result.Events, uint64(dataTx.TxResult.Height))
			if err != nil {
				s.logger.Error("fail to decode evm tx response", "error", err.Error())
				continue
			}

			s.logStream.Add(txLogs...)
		}

		if chHeaders == nil && chLogs == nil {
			break
		}
	}
}
