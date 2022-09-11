package keeper_test

import (
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

	addr2 := sdk.AccAddress("addr2_______________")
	acc2 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr2)
	app.AccountKeeper.SetAccount(suite.Ctx, acc2)

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

	addr2 := sdk.AccAddress("addr2_______________")
	acc2 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr2)
	app.AccountKeeper.SetAccount(suite.Ctx, acc2)

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

	addr2 := sdk.AccAddress("addr2_______________")
	acc2 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr2)
	app.AccountKeeper.SetAccount(suite.Ctx, acc2)

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

	addr2 := sdk.AccAddress("addr2_______________")
	acc2 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr2)
	app.AccountKeeper.SetAccount(suite.Ctx, acc2)

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
