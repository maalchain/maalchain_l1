package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// GetAccount implements AccountKeeperI.
func (ak AccountKeeper) GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI {
	store := ctx.KVStore(ak.key)
	iterator := sdk.KVStorePrefixIterator(store, types.AddressStoreKey(addr))
	defer iterator.Close()
	if !iterator.Valid() {
		return nil
	}

	return ak.decodeAccount(iterator.Value())
}

func (ak AccountKeeper) decodeAccount(bz []byte) types.AccountI {
	acc, err := ak.AccountKeeper.UnmarshalAccount(bz)
	if err != nil {
		panic(err)
	}

	return acc
}
