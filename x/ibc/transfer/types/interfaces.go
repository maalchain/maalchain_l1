// Copyright 2022 Evmos Foundation
// This file is part of the Ethermint Network packages.
//
// Ethermint is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Ethermint packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Ethermint packages. If not, see https://github.com/maalchain/maalchain_l1/blob/main/LICENSE

package types

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"

	erc20types "github.com/maalchain/maalchain_l1/x/erc20/types"
)

// AccountKeeper defines the expected interface needed to retrieve account info.
type AccountKeeper interface {
	transfertypes.AccountKeeper
	GetAccount(sdk.Context, sdk.AccAddress) authtypes.AccountI
}

// BankKeeper defines the expected interface needed to check balances and send coins.
type BankKeeper interface {
	transfertypes.BankKeeper
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
}

// ERC20Keeper defines the expected ERC20 keeper interface for supporting
// ERC20 token transfers via IBC.
type ERC20Keeper interface {
	IsERC20Enabled(ctx sdk.Context) bool
	IsERC20Registered(ctx sdk.Context, contractAddr common.Address) bool
	GetTokenPairID(ctx sdk.Context, token string) []byte
	GetTokenPair(ctx sdk.Context, id []byte) (erc20types.TokenPair, bool)
	ConvertERC20(ctx context.Context, msg *erc20types.MsgConvertERC20) (*erc20types.MsgConvertERC20Response, error)
}
