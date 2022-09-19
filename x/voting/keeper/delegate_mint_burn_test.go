package keeper_test

import (
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
	"mande-chain/testutil"
	"mande-chain/x/voting/types"
)

func (suite *VotingKeeperTestSuite) TestLockMandAndDelegateStake() {
	app, ctx, msgServer := suite.App, sdk.WrapSDKContext(suite.Ctx), suite.MsgServer
	balances := sdk.NewCoins(testutil.NewMandCoin(500))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr1)
	app.AccountKeeper.SetAccount(suite.Ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, suite.Ctx, addr1, balances))

	validator := app.StakingKeeper.GetAllValidators(suite.Ctx)[0]
	initialValidatorStake := validator.GetTokens()
	valAddr, _ := sdk.ValAddressFromBech32(validator.OperatorAddress)
	addr2, _ := sdk.AccAddressFromHex(hex.EncodeToString(valAddr.Bytes()))

	votingModuleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))

	acc1Balance := app.BankKeeper.GetBalance(suite.Ctx, addr1, "mand").Amount.Uint64()
	suite.Require().Equal(uint64(500), acc1Balance)

	// ============= addr1 votes addr2 ===============
	msgCreateVote := &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    100,
		Mode:     1, // cast
	}
	_, err := msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().NoError(err)

	// validate mand transferred to voting module
	acc1Balance = app.BankKeeper.GetBalance(suite.Ctx, addr1, "mand").Amount.Uint64()
	suite.Require().Equal(uint64(400), acc1Balance)
	votingModuleMandBalance := app.BankKeeper.GetBalance(suite.Ctx, votingModuleAcct, "mand").Amount.Uint64()
	suite.Require().Equal(uint64(100), votingModuleMandBalance)

	// validate stake tokens delegated to receiver
	validator, _ = app.StakingKeeper.GetValidator(suite.Ctx, valAddr)
	suite.Require().Equal(validator.GetTokens().Int64(), initialValidatorStake.AddRaw(100).Int64())
}

func (suite *VotingKeeperTestSuite) TestUndelegateStakeAndUnlockMand() {
	app, ctx, msgServer := suite.App, sdk.WrapSDKContext(suite.Ctx), suite.MsgServer
	balances := sdk.NewCoins(testutil.NewMandCoin(500))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr1)
	app.AccountKeeper.SetAccount(suite.Ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, suite.Ctx, addr1, balances))

	validator := app.StakingKeeper.GetAllValidators(suite.Ctx)[0]
	initialValidatorStake := validator.GetTokens()
	valAddr, _ := sdk.ValAddressFromBech32(validator.OperatorAddress)
	addr2, _ := sdk.AccAddressFromHex(hex.EncodeToString(valAddr.Bytes()))

	votingModuleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))

	acc1Balance := app.BankKeeper.GetBalance(suite.Ctx, addr1, "mand").Amount.Uint64()
	suite.Require().Equal(uint64(500), acc1Balance)

	// ============= addr1 casts vote to addr2 ===============
	msgCreateVote := &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    100,
		Mode:     1, // cast
	}
	_, err := msgServer.CreateVote(ctx, msgCreateVote)

	// ============= addr1 un-casts vote from addr2 ===============
	msgCreateVote = &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    50,
		Mode:     0, // un-cast
	}
	_, err = msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().NoError(err)

	// validate mand transferred back from voting module
	acc1Balance = app.BankKeeper.GetBalance(suite.Ctx, addr1, "mand").Amount.Uint64()
	suite.Require().Equal(uint64(450), acc1Balance)
	votingModuleMandBalance := app.BankKeeper.GetBalance(suite.Ctx, votingModuleAcct, "mand").Amount.Uint64()
	suite.Require().Equal(uint64(50), votingModuleMandBalance)

	// validate stake tokens un-delegated from receiver
	validator, _ = app.StakingKeeper.GetValidator(suite.Ctx, valAddr)
	suite.Require().Equal(validator.GetTokens().Int64(), initialValidatorStake.AddRaw(50).Int64())
}
