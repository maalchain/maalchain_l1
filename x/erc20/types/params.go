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
	fmt "fmt"
)

// Parameter store key
var (
	ParamStoreKeyEnableErc20   = []byte("EnableErc20")
	ParamStoreKeyEnableEVMHook = []byte("EnableEVMHook")
)

// NewParams creates a new Params object
func NewParams(
	enableErc20 bool,
	enableEVMHook bool,
) Params {
	return Params{
		EnableErc20:   enableErc20,
		EnableEVMHook: enableEVMHook,
	}
}

func DefaultParams() Params {
	return Params{
		EnableErc20:   true,
		EnableEVMHook: true,
	}
}

func ValidateBool(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func (p Params) Validate() error {
	if err := ValidateBool(p.EnableEVMHook); err != nil {
		return err
	}

	return ValidateBool(p.EnableErc20)
}
