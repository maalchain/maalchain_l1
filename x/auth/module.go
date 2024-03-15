package auth

import (
	"github.com/cosmos/cosmos-sdk/codec"
	auth "github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	types "github.com/cosmos/cosmos-sdk/x/auth/types"
	keeper "github.com/xpladev/ethermint/x/auth/keeper"
)

type AppModule struct {
	auth.AppModule

	accountKeeper keeper.AccountKeeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, accountKeeper keeper.AccountKeeper, randGenAccountsFn types.RandomGenesisAccountsFn, ss exported.Subspace) AppModule {
	return AppModule{
		AppModule:     auth.NewAppModule(cdc, accountKeeper.AccountKeeper, randGenAccountsFn, ss),
		accountKeeper: accountKeeper,
	}
}
