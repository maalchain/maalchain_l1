package keeper_test

import (
	"encoding/json"
	"math"
	"math/big"
	"strconv"
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	ibcgotesting "github.com/cosmos/ibc-go/v7/testing"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
	"github.com/xpladev/ethermint/app"
	"github.com/xpladev/ethermint/contracts"
	"github.com/xpladev/ethermint/crypto/ethsecp256k1"
	"github.com/xpladev/ethermint/ibc"
	ibctesting "github.com/xpladev/ethermint/ibc/testing"
	"github.com/xpladev/ethermint/server/config"
	"github.com/xpladev/ethermint/testutil"
	utiltx "github.com/xpladev/ethermint/testutil/tx"
	"github.com/xpladev/ethermint/x/erc20/types"
	"github.com/xpladev/ethermint/x/evm/statedb"
	evm "github.com/xpladev/ethermint/x/evm/types"
	feemarkettypes "github.com/xpladev/ethermint/x/feemarket/types"
)

func CreatePacket(amount, denom, sender, receiver, srcPort, srcChannel, dstPort, dstChannel string, seq, timeout uint64) channeltypes.Packet {
	transfer := transfertypes.FungibleTokenPacketData{
		Amount:   amount,
		Denom:    denom,
		Receiver: sender,
		Sender:   receiver,
	}
	return channeltypes.NewPacket(
		transfer.GetBytes(),
		seq,
		srcPort,
		srcChannel,
		dstPort,
		dstChannel,
		clienttypes.ZeroHeight(), // timeout height disabled
		timeout,
	)
}

func (suite *KeeperTestSuite) DoSetupTest(t require.TestingT) {
	// account key
	priv, err := ethsecp256k1.GenerateKey()
	require.NoError(t, err)
	suite.priv = priv
	suite.address = common.BytesToAddress(priv.PubKey().Address().Bytes())
	suite.signer = utiltx.NewSigner(priv)

	// consensus key
	privCons, err := ethsecp256k1.GenerateKey()
	require.NoError(t, err)
	consAddress := sdk.ConsAddress(privCons.PubKey().Address())
	suite.consAddress = consAddress

	// init app
	feemarketGenesis := feemarkettypes.DefaultGenesisState()
	suite.app = app.Setup(false, feemarketGenesis)
	header := testutil.NewHeader(
		1, time.Now().UTC(), "maalchain_7862-1", consAddress, nil, nil,
	)
	suite.ctx = suite.app.BaseApp.NewContext(false, header)

	// query clients
	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, suite.app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, suite.app.Erc20Keeper)
	suite.queryClient = types.NewQueryClient(queryHelper)

	queryHelperEvm := baseapp.NewQueryServerTestHelper(suite.ctx, suite.app.InterfaceRegistry())
	evm.RegisterQueryServer(queryHelperEvm, suite.app.EvmKeeper)
	suite.queryClientEvm = evm.NewQueryClient(queryHelperEvm)

	// bond denom
	stakingParams := suite.app.StakingKeeper.GetParams(suite.ctx)
	stakingParams.BondDenom = testutil.BaseDenom
	suite.app.StakingKeeper.SetParams(suite.ctx, stakingParams)

	evmParams := suite.app.EvmKeeper.GetParams(suite.ctx)
	evmParams.EvmDenom = testutil.BaseDenom
	err = suite.app.EvmKeeper.SetParams(suite.ctx, evmParams)
	require.NoError(t, err)

	// Set Validator
	valAddr := sdk.ValAddress(suite.address.Bytes())
	validator, err := stakingtypes.NewValidator(valAddr, privCons.PubKey(), stakingtypes.Description{})
	require.NoError(t, err)
	validator = stakingkeeper.TestingUpdateValidator(suite.app.StakingKeeper, suite.ctx, validator, true)
	err = suite.app.StakingKeeper.Hooks().AfterValidatorCreated(suite.ctx, validator.GetOperator())
	require.NoError(t, err)
	err = suite.app.StakingKeeper.SetValidatorByConsAddr(suite.ctx, validator)
	require.NoError(t, err)

	// fund signer acc to pay for tx fees
	amt := sdkmath.NewInt(int64(math.Pow10(18) * 2))
	err = testutil.FundAccount(
		suite.app.BankKeeper,
		suite.ctx,
		suite.priv.PubKey().Address().Bytes(),
		sdk.NewCoins(sdk.NewCoin(testutil.BaseDenom, amt)),
	)
	suite.Require().NoError(err)

	// TODO change to setup with 1 validator
	validators := s.app.StakingKeeper.GetValidators(s.ctx, 2)
	// set a bonded validator that takes part in consensus
	if validators[0].Status == stakingtypes.Bonded {
		suite.validator = validators[0]
	} else {
		suite.validator = validators[1]
	}

	suite.ethSigner = ethtypes.LatestSignerForChainID(s.app.EvmKeeper.ChainID())

	if suite.suiteIBCTesting {
		suite.SetupIBCTest()
	}
}

func (suite *KeeperTestSuite) SetupIBCTest() {
	// initializes 3 test chains
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 1, 2)
	suite.EthermintChain = suite.coordinator.GetChain(ibcgotesting.GetChainID(1))
	suite.IBCOsmosisChain = suite.coordinator.GetChain(ibcgotesting.GetChainID(2))
	suite.IBCCosmosChain = suite.coordinator.GetChain(ibcgotesting.GetChainID(3))
	suite.coordinator.CommitNBlocks(suite.EthermintChain, 2)
	suite.coordinator.CommitNBlocks(suite.IBCOsmosisChain, 2)
	suite.coordinator.CommitNBlocks(suite.IBCCosmosChain, 2)

	s.app = suite.EthermintChain.App.(*app.EthermintApp)
	evmParams := s.app.EvmKeeper.GetParams(s.EthermintChain.GetContext())
	evmParams.EvmDenom = testutil.BaseDenom
	err := s.app.EvmKeeper.SetParams(s.EthermintChain.GetContext(), evmParams)
	suite.Require().NoError(err)

	// Set block proposer once, so its carried over on the ibc-go-testing suite
	validators := s.app.StakingKeeper.GetValidators(suite.EthermintChain.GetContext(), 2)
	cons, err := validators[0].GetConsAddr()
	suite.Require().NoError(err)
	suite.EthermintChain.CurrentHeader.ProposerAddress = cons.Bytes()

	err = s.app.StakingKeeper.SetValidatorByConsAddr(suite.EthermintChain.GetContext(), validators[0])
	suite.Require().NoError(err)

	_, err = s.app.EvmKeeper.GetCoinbaseAddress(suite.EthermintChain.GetContext(), sdk.ConsAddress(suite.EthermintChain.CurrentHeader.ProposerAddress))
	suite.Require().NoError(err)
	// Mint coins locked on the ethermint account generated with secp.
	amt, ok := sdk.NewIntFromString("1000000000000000000000")
	suite.Require().True(ok)
	coinEthermint := sdk.NewCoin(testutil.BaseDenom, amt)
	coins := sdk.NewCoins(coinEthermint)
	err = s.app.BankKeeper.MintCoins(suite.EthermintChain.GetContext(), minttypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = s.app.BankKeeper.SendCoinsFromModuleToAccount(suite.EthermintChain.GetContext(), minttypes.ModuleName, suite.EthermintChain.SenderAccount.GetAddress(), coins)
	suite.Require().NoError(err)

	// we need some coins in the bankkeeper to be able to register the coins later
	coins = sdk.NewCoins(sdk.NewCoin(ibc.UosmoIbcdenom, sdkmath.NewInt(100)))
	err = s.app.BankKeeper.MintCoins(s.EthermintChain.GetContext(), types.ModuleName, coins)
	s.Require().NoError(err)
	coins = sdk.NewCoins(sdk.NewCoin(ibc.UatomIbcdenom, sdkmath.NewInt(100)))
	err = s.app.BankKeeper.MintCoins(s.EthermintChain.GetContext(), types.ModuleName, coins)
	s.Require().NoError(err)

	// Mint coins on the osmosis side which we'll use to unlock our aphoton
	coinOsmo := sdk.NewCoin("uosmo", sdkmath.NewInt(10000000))
	coins = sdk.NewCoins(coinOsmo)
	err = suite.IBCOsmosisChain.GetSimApp().BankKeeper.MintCoins(suite.IBCOsmosisChain.GetContext(), minttypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.IBCOsmosisChain.GetSimApp().BankKeeper.SendCoinsFromModuleToAccount(suite.IBCOsmosisChain.GetContext(), minttypes.ModuleName, suite.IBCOsmosisChain.SenderAccount.GetAddress(), coins)
	suite.Require().NoError(err)

	// Mint coins on the cosmos side which we'll use to unlock our aphoton
	coinAtom := sdk.NewCoin("uatom", sdkmath.NewInt(10))
	coins = sdk.NewCoins(coinAtom)
	err = suite.IBCCosmosChain.GetSimApp().BankKeeper.MintCoins(suite.IBCCosmosChain.GetContext(), minttypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.IBCCosmosChain.GetSimApp().BankKeeper.SendCoinsFromModuleToAccount(suite.IBCCosmosChain.GetContext(), minttypes.ModuleName, suite.IBCCosmosChain.SenderAccount.GetAddress(), coins)
	suite.Require().NoError(err)

	// Mint coins for IBC tx fee on Osmosis and Cosmos chains
	stkCoin := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, amt))

	err = suite.IBCOsmosisChain.GetSimApp().BankKeeper.MintCoins(suite.IBCOsmosisChain.GetContext(), minttypes.ModuleName, stkCoin)
	suite.Require().NoError(err)
	err = suite.IBCOsmosisChain.GetSimApp().BankKeeper.SendCoinsFromModuleToAccount(suite.IBCOsmosisChain.GetContext(), minttypes.ModuleName, suite.IBCOsmosisChain.SenderAccount.GetAddress(), stkCoin)
	suite.Require().NoError(err)

	err = suite.IBCCosmosChain.GetSimApp().BankKeeper.MintCoins(suite.IBCCosmosChain.GetContext(), minttypes.ModuleName, stkCoin)
	suite.Require().NoError(err)
	err = suite.IBCCosmosChain.GetSimApp().BankKeeper.SendCoinsFromModuleToAccount(suite.IBCCosmosChain.GetContext(), minttypes.ModuleName, suite.IBCCosmosChain.SenderAccount.GetAddress(), stkCoin)
	suite.Require().NoError(err)

	params := types.DefaultParams()
	params.EnableErc20 = true
	err = s.app.Erc20Keeper.SetParams(suite.EthermintChain.GetContext(), params)
	suite.Require().NoError(err)

	suite.pathOsmosisEthermint = ibctesting.NewTransferPath(suite.IBCOsmosisChain, suite.EthermintChain) // clientID, connectionID, channelID empty
	suite.pathCosmosEthermint = ibctesting.NewTransferPath(suite.IBCCosmosChain, suite.EthermintChain)
	suite.pathOsmosisCosmos = ibctesting.NewTransferPath(suite.IBCCosmosChain, suite.IBCOsmosisChain)
	ibctesting.SetupPath(suite.coordinator, suite.pathOsmosisEthermint) // clientID, connectionID, channelID filled
	ibctesting.SetupPath(suite.coordinator, suite.pathCosmosEthermint)
	ibctesting.SetupPath(suite.coordinator, suite.pathOsmosisCosmos)
	suite.Require().Equal("07-tendermint-0", suite.pathOsmosisEthermint.EndpointA.ClientID)
	suite.Require().Equal("connection-0", suite.pathOsmosisEthermint.EndpointA.ConnectionID)
	suite.Require().Equal("channel-0", suite.pathOsmosisEthermint.EndpointA.ChannelID)

	coinEthermint = sdk.NewCoin(testutil.BaseDenom, sdkmath.NewInt(1000000000000000000))
	coins = sdk.NewCoins(coinEthermint)
	err = s.app.BankKeeper.MintCoins(suite.EthermintChain.GetContext(), types.ModuleName, coins)
	suite.Require().NoError(err)
	err = s.app.BankKeeper.SendCoinsFromModuleToModule(suite.EthermintChain.GetContext(), types.ModuleName, authtypes.FeeCollectorName, coins)
	suite.Require().NoError(err)
}

var timeoutHeight = clienttypes.NewHeight(1000, 1000)

func (suite *KeeperTestSuite) StateDB() *statedb.StateDB {
	return statedb.New(suite.ctx, suite.app.EvmKeeper, statedb.NewEmptyTxConfig(common.BytesToHash(suite.ctx.HeaderHash().Bytes())))
}

func (suite *KeeperTestSuite) MintFeeCollector(coins sdk.Coins) {
	err := suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToModule(suite.ctx, types.ModuleName, authtypes.FeeCollectorName, coins)
	suite.Require().NoError(err)
}

func (suite *KeeperTestSuite) sendTx(contractAddr, from common.Address, transferData []byte) *evm.MsgEthereumTx {
	ctx := sdk.WrapSDKContext(suite.ctx)
	chainID := suite.app.EvmKeeper.ChainID()

	args, err := json.Marshal(&evm.TransactionArgs{To: &contractAddr, From: &from, Data: (*hexutil.Bytes)(&transferData)})
	suite.Require().NoError(err)
	res, err := suite.queryClientEvm.EstimateGas(ctx, &evm.EthCallRequest{
		Args:   args,
		GasCap: config.DefaultGasCap,
	})
	suite.Require().NoError(err)

	nonce := suite.app.EvmKeeper.GetNonce(suite.ctx, suite.address)

	// Mint the max gas to the FeeCollector to ensure balance in case of refund
	evmParams := suite.app.EvmKeeper.GetParams(suite.ctx)
	suite.MintFeeCollector(sdk.NewCoins(sdk.NewCoin(evmParams.EvmDenom, sdkmath.NewInt(suite.app.FeeMarketKeeper.GetBaseFee(suite.ctx).Int64()*int64(res.Gas)))))
	ercTransferTxParams := &evm.EvmTxArgs{
		ChainID:   chainID,
		Nonce:     nonce,
		To:        &contractAddr,
		GasLimit:  res.Gas,
		GasFeeCap: suite.app.FeeMarketKeeper.GetBaseFee(suite.ctx),
		GasTipCap: big.NewInt(1),
		Input:     transferData,
		Accesses:  &ethtypes.AccessList{},
	}
	ercTransferTx := evm.NewTx(ercTransferTxParams)

	ercTransferTx.From = suite.address.Hex()
	err = ercTransferTx.Sign(ethtypes.LatestSignerForChainID(chainID), suite.signer)
	suite.Require().NoError(err)
	rsp, err := suite.app.EvmKeeper.EthereumTx(ctx, ercTransferTx)
	suite.Require().NoError(err)
	suite.Require().Empty(rsp.VmError)
	return ercTransferTx
}

// Commit commits and starts a new block with an updated context.
func (suite *KeeperTestSuite) Commit() {
	suite.CommitAndBeginBlockAfter(time.Hour * 1)
}

// Commit commits a block at a given time. Reminder: At the end of each
// Tendermint Consensus round the following methods are run
//  1. BeginBlock
//  2. DeliverTx
//  3. EndBlock
//  4. Commit
func (suite *KeeperTestSuite) CommitAndBeginBlockAfter(t time.Duration) {
	var err error
	suite.ctx, err = testutil.Commit(suite.ctx, suite.app, t, nil)
	suite.Require().NoError(err)

	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, suite.app.InterfaceRegistry())
	evm.RegisterQueryServer(queryHelper, suite.app.EvmKeeper)
	suite.queryClientEvm = evm.NewQueryClient(queryHelper)
}

// DeployContract deploys the ERC20MinterBurnerDecimalsContract.
func (suite *KeeperTestSuite) DeployContract(name, symbol string, decimals uint8) (common.Address, error) {
	suite.Commit()
	addr, err := testutil.DeployContract(
		suite.ctx,
		suite.app,
		suite.priv,
		suite.queryClientEvm,
		contracts.ERC20MinterBurnerDecimalsContract,
		name, symbol, decimals,
	)
	suite.Commit()
	return addr, err
}

func (suite *KeeperTestSuite) DeployContractMaliciousDelayed() (common.Address, error) {
	suite.Commit()
	addr, err := testutil.DeployContract(
		suite.ctx,
		suite.app,
		suite.priv,
		suite.queryClientEvm,
		contracts.ERC20MaliciousDelayedContract,
		big.NewInt(1000000000000000000),
	)
	suite.Commit()
	return addr, err
}

func (suite *KeeperTestSuite) DeployContractDirectBalanceManipulation() (common.Address, error) {
	suite.Commit()
	addr, err := testutil.DeployContract(
		suite.ctx,
		suite.app,
		suite.priv,
		suite.queryClientEvm,
		contracts.ERC20DirectBalanceManipulationContract,
		big.NewInt(1000000000000000000),
	)
	suite.Commit()
	return addr, err
}

// DeployContractToChain deploys the ERC20MinterBurnerDecimalsContract
// to the Ethermint chain (used on IBC tests)
func (suite *KeeperTestSuite) DeployContractToChain(name, symbol string, decimals uint8) (common.Address, error) {
	return testutil.DeployContract(
		s.EthermintChain.GetContext(),
		s.EthermintChain.App.(*app.EthermintApp),
		suite.EthermintChain.SenderPrivKey,
		suite.queryClientEvm,
		contracts.ERC20MinterBurnerDecimalsContract,
		name, symbol, decimals,
	)
}

func (suite *KeeperTestSuite) sendAndReceiveMessage(
	path *ibctesting.Path,
	originEndpoint *ibctesting.Endpoint,
	destEndpoint *ibctesting.Endpoint,
	originChain *ibcgotesting.TestChain,
	coin string,
	amount int64,
	sender, receiver string,
	seq uint64,
	ibcCoinMetadata string,
) {
	transferMsg := transfertypes.NewMsgTransfer(originEndpoint.ChannelConfig.PortID, originEndpoint.ChannelID, sdk.NewCoin(coin, sdkmath.NewInt(amount)), sender, receiver, timeoutHeight, 0, "")
	_, err := ibctesting.SendMsgs(originChain, ibctesting.DefaultFeeAmt, transferMsg)
	suite.Require().NoError(err) // message committed
	// Recreate the packet that was sent
	var transfer transfertypes.FungibleTokenPacketData
	if ibcCoinMetadata == "" {
		transfer = transfertypes.NewFungibleTokenPacketData(coin, strconv.Itoa(int(amount)), sender, receiver, "")
	} else {
		transfer = transfertypes.NewFungibleTokenPacketData(ibcCoinMetadata, strconv.Itoa(int(amount)), sender, receiver, "")
	}
	packet := channeltypes.NewPacket(transfer.GetBytes(), seq, originEndpoint.ChannelConfig.PortID, originEndpoint.ChannelID, destEndpoint.ChannelConfig.PortID, destEndpoint.ChannelID, timeoutHeight, 0)
	// Receive message on the counterparty side, and send ack
	err = path.RelayPacket(packet)
	suite.Require().NoError(err)
}

func (suite *KeeperTestSuite) SendAndReceiveMessage(path *ibctesting.Path, origin *ibcgotesting.TestChain, coin string, amount int64, sender, receiver string, seq uint64, ibcCoinMetadata string) {
	// Send coin from A to B
	suite.sendAndReceiveMessage(path, path.EndpointA, path.EndpointB, origin, coin, amount, sender, receiver, seq, ibcCoinMetadata)
}

// Send back coins (from path endpoint B to A). In case of IBC coins need to provide ibcCoinMetadata (<port>/<channel>/<denom>, e.g.: "transfer/channel-0/aphoton") as input parameter.
// We need this to instantiate properly a FungibleTokenPacketData https://github.com/cosmos/ibc-go/blob/main/docs/architecture/adr-001-coin-source-tracing.md
func (suite *KeeperTestSuite) SendBackCoins(path *ibctesting.Path, origin *ibcgotesting.TestChain, coin string, amount int64, sender, receiver string, seq uint64, ibcCoinMetadata string) {
	// Send coin from B to A
	suite.sendAndReceiveMessage(path, path.EndpointB, path.EndpointA, origin, coin, amount, sender, receiver, seq, ibcCoinMetadata)
}
