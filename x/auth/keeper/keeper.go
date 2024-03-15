package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type AccountKeeper struct {
	authkeeper.AccountKeeper
	key storetypes.StoreKey
}

func NewAccountKeeper(
	cdc codec.BinaryCodec, key storetypes.StoreKey, proto func() authtypes.AccountI,
	maccPerms map[string][]string, bech32Prefix string, authority string,
) AccountKeeper {
	return AccountKeeper{
		AccountKeeper: authkeeper.NewAccountKeeper(cdc, key, proto, maccPerms, bech32Prefix, authority),
		key:           key,
	}
}
