package keeper_test

import (
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"mande-chain/testutil"
	votingkeeper "mande-chain/x/voting/keeper"
	"mande-chain/x/voting/types"
	votingtypes "mande-chain/x/voting/types"
	"testing"
)

type VotingKeeperTestSuite struct {
	testutil.IntegrationTestSuite
	MsgServer votingtypes.MsgServer
}

func (suite *VotingKeeperTestSuite) SetupTest() {
	suite.IntegrationTestSuite.SetupTest()
	suite.MsgServer = votingkeeper.NewMsgServerImpl(suite.App.VotingKeeper)
}

func (suite *VotingKeeperTestSuite) TestCreateVote_Cast_Positive_Valid() {
	app, ctx, msgServer, keeper := suite.App, sdk.WrapSDKContext(suite.Ctx), suite.MsgServer, suite.App.VotingKeeper
	balances := sdk.NewCoins(testutil.NewMandCoin(10000))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr1)
	app.AccountKeeper.SetAccount(suite.Ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, suite.Ctx, addr1, balances))

	validator := app.StakingKeeper.GetAllValidators(suite.Ctx)[0]
	valAddr, _ := sdk.ValAddressFromBech32(validator.OperatorAddress)
	addr2, _ := sdk.AccAddressFromHex(hex.EncodeToString(valAddr.Bytes()))

	// ============= addr1 votes addr2 ===============
	msgCreateVote := &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    100,
		Mode:     1, // cast
	}
	_, err := msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().NoError(err)

	// validate aggregate vote counts after tx
	aggregateVoteCountCreator, _ := keeper.GetAggregateVoteCount(suite.Ctx, addr1.String())
	aggregateVoteCountReceiver, _ := keeper.GetAggregateVoteCount(suite.Ctx, addr2.String())
	suite.validateAggregateVoteCount(aggregateVoteCountCreator, 100, 0, 0)
	suite.validateAggregateVoteCount(aggregateVoteCountReceiver, 0, 100, 0)

	// validate vote book
	voteBookEntry, _ := keeper.GetVoteBook(suite.Ctx, types.VoteBookIndex(addr1.String(), addr2.String()))
	suite.validateVoteBook(voteBookEntry, 100, 0)
}

func (suite *VotingKeeperTestSuite) TestCreateVote_Cast_Negative_Valid() {
	app, ctx, msgServer, keeper := suite.App, sdk.WrapSDKContext(suite.Ctx), suite.MsgServer, suite.App.VotingKeeper
	balances := sdk.NewCoins(testutil.NewMandCoin(10000))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr1)
	app.AccountKeeper.SetAccount(suite.Ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, suite.Ctx, addr1, balances))

	validator := app.StakingKeeper.GetAllValidators(suite.Ctx)[0]
	valAddr, _ := sdk.ValAddressFromBech32(validator.OperatorAddress)
	addr2, _ := sdk.AccAddressFromHex(hex.EncodeToString(valAddr.Bytes()))

	// ============= addr1 votes addr2 ===============
	msgCreateVote := &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    -100,
		Mode:     1, // cast
	}
	_, err := msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().NoError(err)

	// validate aggregate vote counts after tx
	aggregateVoteCountCreator, _ := keeper.GetAggregateVoteCount(suite.Ctx, addr1.String())
	aggregateVoteCountReceiver, _ := keeper.GetAggregateVoteCount(suite.Ctx, addr2.String())
	suite.validateAggregateVoteCount(aggregateVoteCountCreator, 100, 0, 0)
	suite.validateAggregateVoteCount(aggregateVoteCountReceiver, 0, 0, 100)

	// validate vote book
	voteBookEntry, _ := keeper.GetVoteBook(suite.Ctx, types.VoteBookIndex(addr1.String(), addr2.String()))
	suite.validateVoteBook(voteBookEntry, 0, 100)
}

func (suite *VotingKeeperTestSuite) TestCreateVote_UnCast_Positive_Valid() {
	app, ctx, msgServer, keeper := suite.App, sdk.WrapSDKContext(suite.Ctx), suite.MsgServer, suite.App.VotingKeeper
	balances := sdk.NewCoins(testutil.NewMandCoin(10000))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr1)
	app.AccountKeeper.SetAccount(suite.Ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, suite.Ctx, addr1, balances))

	validator := app.StakingKeeper.GetAllValidators(suite.Ctx)[0]
	valAddr, _ := sdk.ValAddressFromBech32(validator.OperatorAddress)
	addr2, _ := sdk.AccAddressFromHex(hex.EncodeToString(valAddr.Bytes()))

	// ============= addr1 votes addr2 ===============
	msgCreateVote := &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    100,
		Mode:     1, // cast
	}
	_, err := msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().NoError(err)

	// validate aggregate vote counts after cast tx
	aggregateVoteCountCreator, _ := keeper.GetAggregateVoteCount(suite.Ctx, addr1.String())
	aggregateVoteCountReceiver, _ := keeper.GetAggregateVoteCount(suite.Ctx, addr2.String())
	suite.validateAggregateVoteCount(aggregateVoteCountCreator, 100, 0, 0)
	suite.validateAggregateVoteCount(aggregateVoteCountReceiver, 0, 100, 0)

	// ============= addr1 votes addr2 ===============
	msgCreateVote = &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    50,
		Mode:     0, // un-cast
	}
	_, err = msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().NoError(err)

	// validate aggregate vote counts after un-cast tx
	aggregateVoteCountCreator, _ = keeper.GetAggregateVoteCount(suite.Ctx, addr1.String())
	aggregateVoteCountReceiver, _ = keeper.GetAggregateVoteCount(suite.Ctx, addr2.String())
	suite.validateAggregateVoteCount(aggregateVoteCountCreator, 50, 0, 0)
	suite.validateAggregateVoteCount(aggregateVoteCountReceiver, 0, 50, 0)

	// validate vote book
	voteBookEntry, _ := keeper.GetVoteBook(suite.Ctx, types.VoteBookIndex(addr1.String(), addr2.String()))
	suite.validateVoteBook(voteBookEntry, 50, 0)
}

func (suite *VotingKeeperTestSuite) TestCreateVote_UnCast_Negative_Valid() {
	app, ctx, msgServer, keeper := suite.App, sdk.WrapSDKContext(suite.Ctx), suite.MsgServer, suite.App.VotingKeeper
	balances := sdk.NewCoins(testutil.NewMandCoin(10000))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr1)
	app.AccountKeeper.SetAccount(suite.Ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, suite.Ctx, addr1, balances))

	validator := app.StakingKeeper.GetAllValidators(suite.Ctx)[0]
	valAddr, _ := sdk.ValAddressFromBech32(validator.OperatorAddress)
	addr2, _ := sdk.AccAddressFromHex(hex.EncodeToString(valAddr.Bytes()))

	// ============= addr1 votes addr2 ===============
	msgCreateVote := &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    -100,
		Mode:     1, // cast
	}
	_, err := msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().NoError(err)

	// validate aggregate vote counts after cast tx
	aggregateVoteCountCreator, _ := keeper.GetAggregateVoteCount(suite.Ctx, addr1.String())
	aggregateVoteCountReceiver, _ := keeper.GetAggregateVoteCount(suite.Ctx, addr2.String())
	suite.validateAggregateVoteCount(aggregateVoteCountCreator, 100, 0, 0)
	suite.validateAggregateVoteCount(aggregateVoteCountReceiver, 0, 0, 100)

	// ============= addr1 votes addr2 ===============
	msgCreateVote = &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    -50,
		Mode:     0, // un-cast
	}
	_, err = msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().NoError(err)

	// validate aggregate vote counts after un-cast tx
	aggregateVoteCountCreator, _ = keeper.GetAggregateVoteCount(suite.Ctx, addr1.String())
	aggregateVoteCountReceiver, _ = keeper.GetAggregateVoteCount(suite.Ctx, addr2.String())
	suite.validateAggregateVoteCount(aggregateVoteCountCreator, 50, 0, 0)
	suite.validateAggregateVoteCount(aggregateVoteCountReceiver, 0, 0, 50)

	// validate vote book
	voteBookEntry, _ := keeper.GetVoteBook(suite.Ctx, types.VoteBookIndex(addr1.String(), addr2.String()))
	suite.validateVoteBook(voteBookEntry, 0, 50)
}

func (suite *VotingKeeperTestSuite) TestValidation_VoteExceedsBalance() {
	app, ctx, msgServer := suite.App, sdk.WrapSDKContext(suite.Ctx), suite.MsgServer
	balances := sdk.NewCoins(testutil.NewMandCoin(500))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr1)
	app.AccountKeeper.SetAccount(suite.Ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, suite.Ctx, addr1, balances))

	validator := app.StakingKeeper.GetAllValidators(suite.Ctx)[0]
	valAddr, _ := sdk.ValAddressFromBech32(validator.OperatorAddress)
	addr2, _ := sdk.AccAddressFromHex(hex.EncodeToString(valAddr.Bytes()))

	acc1Balance := app.BankKeeper.GetBalance(suite.Ctx, addr1, "mand").Amount.Uint64()
	suite.Require().Equal(uint64(500), acc1Balance)

	// ============= addr1 votes addr2 ===============
	msgCreateVote := &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    600,
		Mode:     1, // cast
	}
	_, err := msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().ErrorIs(err, types.ErrNotEnoughMand)
}

func (suite *VotingKeeperTestSuite) TestValidation_NoVotingRecord() {
	app, ctx, msgServer := suite.App, sdk.WrapSDKContext(suite.Ctx), suite.MsgServer
	balances := sdk.NewCoins(testutil.NewMandCoin(500))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr1)
	app.AccountKeeper.SetAccount(suite.Ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, suite.Ctx, addr1, balances))

	validator := app.StakingKeeper.GetAllValidators(suite.Ctx)[0]
	valAddr, _ := sdk.ValAddressFromBech32(validator.OperatorAddress)
	addr2, _ := sdk.AccAddressFromHex(hex.EncodeToString(valAddr.Bytes()))

	acc1Balance := app.BankKeeper.GetBalance(suite.Ctx, addr1, "mand").Amount.Uint64()
	suite.Require().Equal(uint64(500), acc1Balance)

	// ============= addr1 votes addr2 ===============
	msgCreateVote := &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    100,
		Mode:     0, // un-cast
	}
	_, err := msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().ErrorIs(err, types.ErrNoVoteRecord)
}

func (suite *VotingKeeperTestSuite) TestValidation_PositiveVotesCastedLimit() {
	app, ctx, msgServer := suite.App, sdk.WrapSDKContext(suite.Ctx), suite.MsgServer
	balances := sdk.NewCoins(testutil.NewMandCoin(500))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr1)
	app.AccountKeeper.SetAccount(suite.Ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, suite.Ctx, addr1, balances))

	validator := app.StakingKeeper.GetAllValidators(suite.Ctx)[0]
	valAddr, _ := sdk.ValAddressFromBech32(validator.OperatorAddress)
	addr2, _ := sdk.AccAddressFromHex(hex.EncodeToString(valAddr.Bytes()))

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

	// ============= addr1 un-votes addr2 ===============
	msgCreateVote = &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    200,
		Mode:     0, // un-cast
	}
	_, err = msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().ErrorIs(err, types.ErrPositiveVotesCastedLimit)
}

func (suite *VotingKeeperTestSuite) TestValidation_NegativeVotesCastedLimit() {
	app, ctx, msgServer := suite.App, sdk.WrapSDKContext(suite.Ctx), suite.MsgServer
	balances := sdk.NewCoins(testutil.NewMandCoin(500))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr1)
	app.AccountKeeper.SetAccount(suite.Ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, suite.Ctx, addr1, balances))

	validator := app.StakingKeeper.GetAllValidators(suite.Ctx)[0]
	valAddr, _ := sdk.ValAddressFromBech32(validator.OperatorAddress)
	addr2, _ := sdk.AccAddressFromHex(hex.EncodeToString(valAddr.Bytes()))

	acc1Balance := app.BankKeeper.GetBalance(suite.Ctx, addr1, "mand").Amount.Uint64()
	suite.Require().Equal(uint64(500), acc1Balance)

	// ============= addr1 votes addr2 ===============
	msgCreateVote := &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    -100,
		Mode:     1, // cast
	}
	_, err := msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().NoError(err)

	// ============= addr1 un-votes addr2 ===============
	msgCreateVote = &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    -200,
		Mode:     0, // un-cast
	}
	_, err = msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().ErrorIs(err, types.ErrNegativeVotesCastedLimit)
}

func (suite *VotingKeeperTestSuite) validateAggregateVoteCount(aggregateVoteCount types.AggregateVoteCount, casted uint64, positiveReceived uint64, negativeReceived uint64) {
	suite.Require().Equal(casted, aggregateVoteCount.GetCasted())
	suite.Require().Equal(positiveReceived, aggregateVoteCount.GetPositiveReceived())
	suite.Require().Equal(negativeReceived, aggregateVoteCount.GetNegativeReceived())
}

func (suite *VotingKeeperTestSuite) validateVoteBook(voteBookEntry types.VoteBook, positive uint64, negative uint64) {
	suite.Require().Equal(positive, voteBookEntry.GetPositive())
	suite.Require().Equal(negative, voteBookEntry.GetNegative())
}

func TestVotingKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(VotingKeeperTestSuite))
}
