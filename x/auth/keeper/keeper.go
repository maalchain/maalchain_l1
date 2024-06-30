package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type AccountKeeper struct {
	authkeeper.AccountKeeper
	storeKey storetypes.StoreKey
}

func NewAccountKeeper(
	cdc codec.BinaryCodec, storeKey storetypes.StoreKey, proto func() authtypes.AccountI,
	maccPerms map[string][]string, bech32Prefix string, authority string,
) AccountKeeper {
	return AccountKeeper{
		AccountKeeper: authkeeper.NewAccountKeeper(cdc, storeKey, proto, maccPerms, bech32Prefix, authority),
		storeKey:      storeKey,
	}
}

func (ak AccountKeeper) decodeAccount(bz []byte) authtypes.AccountI {
	acc, err := ak.AccountKeeper.UnmarshalAccount(bz)
	if err != nil {
		panic(err)
	}

	return acc
}
