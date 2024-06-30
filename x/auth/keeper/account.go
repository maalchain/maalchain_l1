package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// GetAccount implements AccountKeeperI.
func (ak AccountKeeper) GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI {
	store := ctx.KVStore(ak.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, authtypes.AddressStoreKey(addr))
	defer iterator.Close()
	if !iterator.Valid() {
		return nil
	}

	return ak.decodeAccount(iterator.Value())
}