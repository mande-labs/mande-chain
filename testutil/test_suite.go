package testutil

import (
	"time"

	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"mande-chain/app"
)

const (
	mandDenom = "mand"
)

func NewMandCoin(amt int64) sdk.Coin {
	return sdk.NewInt64Coin(mandDenom, amt)
}

type IntegrationTestSuite struct {
	suite.Suite

	App         *app.App
	Ctx         sdk.Context
	queryClient banktypes.QueryClient
}

func (suite *IntegrationTestSuite) SetupTest() {
	app := app.Setup(suite.T(), false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Time: time.Now()})

	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
	app.BankKeeper.SetParams(ctx, banktypes.DefaultParams())

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	banktypes.RegisterQueryServer(queryHelper, app.BankKeeper)
	queryClient := banktypes.NewQueryClient(queryHelper)

	suite.App = app
	suite.Ctx = ctx
	suite.queryClient = queryClient
}
