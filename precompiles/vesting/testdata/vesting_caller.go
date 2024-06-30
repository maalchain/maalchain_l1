// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package testdata

import (
	contractutils "github.com/maalchain/maalchain_l1/contracts/utils"
	evmtypes "github.com/maalchain/maalchain_l1/x/evm/types"
)

func LoadVestingCallerContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("VestingCaller.json")
}
