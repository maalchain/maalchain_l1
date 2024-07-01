package keeper_test

import (
	"fmt"
	"testing"

	"github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/ethermint/testutil"
	"github.com/stretchr/testify/suite"
)

type ABCITestSuite struct {
	testutil.BaseTestSuite
}

func TestABCITestSuite(t *testing.T) {
	suite.Run(t, new(ABCITestSuite))
}

func (suite *ABCITestSuite) TestEndBlock() {
	testCases := []struct {
		name         string
		NoBaseFee    bool
		malleate     func()
		expGasWanted uint64
	}{
		{
			"baseFee nil",
			true,
			func() {},
			uint64(0),
		},
		{
			"pass",
			false,
			func() {
				meter := sdk.NewGasMeter(uint64(1000000000))
				suite.Ctx = suite.Ctx.WithBlockGasMeter(meter)
				suite.App.FeeMarketKeeper.SetTransientBlockGasWanted(suite.Ctx, 5000000)
			},
			uint64(2500000),
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset
			params := suite.App.FeeMarketKeeper.GetParams(suite.Ctx)
			params.NoBaseFee = tc.NoBaseFee
			suite.App.FeeMarketKeeper.SetParams(suite.Ctx, params)

			tc.malleate()
			suite.App.FeeMarketKeeper.EndBlock(suite.Ctx, types.RequestEndBlock{Height: 1})
			gasWanted := suite.App.FeeMarketKeeper.GetBlockGasWanted(suite.Ctx)
			suite.Require().Equal(tc.expGasWanted, gasWanted, tc.name)
		})
	}
}
