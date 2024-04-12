package auth

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	ethauthkeeper "github.com/xpladev/ethermint/x/auth/keeper"
)

type AppModule struct {
	auth.AppModule

	accountKeeper ethauthkeeper.AccountKeeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, accountKeeper ethauthkeeper.AccountKeeper, randGenAccountsFn authtypes.RandomGenesisAccountsFn, ss exported.Subspace) AppModule {
	return AppModule{
		AppModule:     auth.NewAppModule(cdc, accountKeeper.AccountKeeper, randGenAccountsFn, ss),
		accountKeeper: accountKeeper,
	}
}
