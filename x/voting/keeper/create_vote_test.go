package keeper_test

import (
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"mande-chain/testutil"
	votingkeeper "mande-chain/x/voting/keeper"
	"mande-chain/x/voting/types"
	votingtypes "mande-chain/x/voting/types"
	oracletypes "mande-chain/x/oracle/types"
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

// Cast valid validator votes
func (suite *VotingKeeperTestSuite) Test_Create_Vote_Cast_Validator_Valid_Positive() {
	app, ctx, msgServer, keeper := suite.App, sdk.WrapSDKContext(suite.Ctx), suite.MsgServer, suite.App.VotingKeeper
	balances := sdk.NewCoins(testutil.NewMandCoin(10000))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr1)
	app.AccountKeeper.SetAccount(suite.Ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, suite.Ctx, addr1, balances))

	validator := app.StakingKeeper.GetAllValidators(suite.Ctx)[0]
	valAddr, _ := sdk.ValAddressFromBech32(validator.OperatorAddress)
	addr2, _ := sdk.AccAddressFromHex(hex.EncodeToString(valAddr.Bytes()))

	app.OracleKeeper.SetNetworkConstantResult(suite.Ctx, oracletypes.OracleRequestID(1), oracletypes.NetworkConstantResult{Response: "1.0"})
	app.OracleKeeper.SetLastNetworkConstantID(suite.Ctx, oracletypes.OracleRequestID(1))

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
	aggregateVoteCountCreator, _ := keeper.GetAggregateVotesCasted(suite.Ctx, addr1.String())
	aggregateVoteCountReceiver, _ := keeper.GetAggregateVotesReceived(suite.Ctx, addr2.String())
	suite.Require().Equal(uint64(100), aggregateVoteCountCreator.GetPositive())
	suite.Require().Equal(uint64(100), aggregateVoteCountReceiver.GetPositive())

	// validate vote book
	voteBookEntry, _ := keeper.GetVoteBook(suite.Ctx, types.VoteBookIndex(addr2.String(), addr1.String()))
	suite.validateVoteBook(voteBookEntry, 100, 0)

	credibility, _ := keeper.GetCredibility(suite.Ctx, addr2.String())
	suite.Require().Equal("100.00", credibility.Score) // because X is set to 1
}

func (suite *VotingKeeperTestSuite) Test_Create_Vote_Cast_Validator_Valid_Negative() {
	app, ctx, msgServer, keeper := suite.App, sdk.WrapSDKContext(suite.Ctx), suite.MsgServer, suite.App.VotingKeeper
	balances := sdk.NewCoins(testutil.NewMandCoin(10000))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr1)
	app.AccountKeeper.SetAccount(suite.Ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, suite.Ctx, addr1, balances))

	validator := app.StakingKeeper.GetAllValidators(suite.Ctx)[0]
	valAddr, _ := sdk.ValAddressFromBech32(validator.OperatorAddress)
	addr2, _ := sdk.AccAddressFromHex(hex.EncodeToString(valAddr.Bytes()))

	app.OracleKeeper.SetNetworkConstantResult(suite.Ctx, oracletypes.OracleRequestID(1), oracletypes.NetworkConstantResult{Response: "1.0"})
	app.OracleKeeper.SetLastNetworkConstantID(suite.Ctx, oracletypes.OracleRequestID(1))

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
	aggregateVoteCountCreator, _ := keeper.GetAggregateVotesCasted(suite.Ctx, addr1.String())
	aggregateVoteCountReceiver, _ := keeper.GetAggregateVotesReceived(suite.Ctx, addr2.String())
	suite.Require().Equal(uint64(100), aggregateVoteCountCreator.GetNegative())
	suite.Require().Equal(uint64(100), aggregateVoteCountReceiver.GetNegative())

	// validate vote book
	voteBookEntry, _ := keeper.GetVoteBook(suite.Ctx, types.VoteBookIndex(addr2.String(), addr1.String()))
	suite.validateVoteBook(voteBookEntry, 0, 100)

	credibility, _ := keeper.GetCredibility(suite.Ctx, addr2.String())
	suite.Require().Equal("-100.00", credibility.Score) // because X is set to 1
}

func (suite *VotingKeeperTestSuite) Test_Create_Vote_Cast_Non_Validator_Valid_Positive() {
	app, ctx, msgServer, keeper := suite.App, sdk.WrapSDKContext(suite.Ctx), suite.MsgServer, suite.App.VotingKeeper
	balances := sdk.NewCoins(testutil.NewMandCoin(10000))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr1)
	app.AccountKeeper.SetAccount(suite.Ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, suite.Ctx, addr1, balances))

	addr2 := sdk.AccAddress("addr2_______________")
	acc2 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr2)
	app.AccountKeeper.SetAccount(suite.Ctx, acc2)

	app.OracleKeeper.SetNetworkConstantResult(suite.Ctx, oracletypes.OracleRequestID(1), oracletypes.NetworkConstantResult{Response: "1.0"})
	app.OracleKeeper.SetLastNetworkConstantID(suite.Ctx, oracletypes.OracleRequestID(1))

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
	aggregateVoteCountCreator, _ := keeper.GetAggregateVotesCasted(suite.Ctx, addr1.String())
	aggregateVoteCountReceiver, _ := keeper.GetAggregateVotesReceived(suite.Ctx, addr2.String())
	suite.Require().Equal(uint64(100), aggregateVoteCountCreator.GetPositive())
	suite.Require().Equal(uint64(100), aggregateVoteCountReceiver.GetPositive())

	// validate vote book
	voteBookEntry, _ := keeper.GetVoteBook(suite.Ctx, types.VoteBookIndex(addr2.String(), addr1.String()))
	suite.validateVoteBook(voteBookEntry, 100, 0)

	credibility, _ := keeper.GetCredibility(suite.Ctx, addr2.String())
	suite.Require().Equal("100.00", credibility.Score) // because X is set to 1
}

func (suite *VotingKeeperTestSuite) Test_Create_Vote_Cast_Non_Validator_Valid_Negative() {
	app, ctx, msgServer, keeper := suite.App, sdk.WrapSDKContext(suite.Ctx), suite.MsgServer, suite.App.VotingKeeper
	balances := sdk.NewCoins(testutil.NewMandCoin(10000))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr1)
	app.AccountKeeper.SetAccount(suite.Ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, suite.Ctx, addr1, balances))

	addr2 := sdk.AccAddress("addr2_______________")
	acc2 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr2)
	app.AccountKeeper.SetAccount(suite.Ctx, acc2)

	app.OracleKeeper.SetNetworkConstantResult(suite.Ctx, oracletypes.OracleRequestID(1), oracletypes.NetworkConstantResult{Response: "1.0"})
	app.OracleKeeper.SetLastNetworkConstantID(suite.Ctx, oracletypes.OracleRequestID(1))

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
	aggregateVoteCountCreator, _ := keeper.GetAggregateVotesCasted(suite.Ctx, addr1.String())
	aggregateVoteCountReceiver, _ := keeper.GetAggregateVotesReceived(suite.Ctx, addr2.String())
	suite.Require().Equal(uint64(100), aggregateVoteCountCreator.GetNegative())
	suite.Require().Equal(uint64(100), aggregateVoteCountReceiver.GetNegative())

	// validate vote book
	voteBookEntry, _ := keeper.GetVoteBook(suite.Ctx, types.VoteBookIndex(addr2.String(), addr1.String()))
	suite.validateVoteBook(voteBookEntry, 0, 100)

	credibility, _ := keeper.GetCredibility(suite.Ctx, addr2.String())
	suite.Require().Equal("-100.00", credibility.Score) // because X is set to 1
}

func (suite *VotingKeeperTestSuite) Test_Create_Vote_Self_Cast_UnCast_Validator_Valid_Positive() {
	app, ctx, msgServer, keeper := suite.App, sdk.WrapSDKContext(suite.Ctx), suite.MsgServer, suite.App.VotingKeeper
	balances := sdk.NewCoins(testutil.NewMandCoin(10000))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr1)
	app.AccountKeeper.SetAccount(suite.Ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, suite.Ctx, addr1, balances))

	app.OracleKeeper.SetNetworkConstantResult(suite.Ctx, oracletypes.OracleRequestID(1), oracletypes.NetworkConstantResult{Response: "1.0"})
	app.OracleKeeper.SetLastNetworkConstantID(suite.Ctx, oracletypes.OracleRequestID(1))

	// ============= addr1 votes addr2 ===============
	msgCreateVote := &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr1.String(),
		Count:    100,
		Mode:     1, // cast
	}
	_, err := msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().NoError(err)

	// validate aggregate vote counts after tx
	aggregateVoteCountCreator, _ := keeper.GetAggregateVotesCasted(suite.Ctx, addr1.String())
	aggregateVoteCountReceiver, _ := keeper.GetAggregateVotesReceived(suite.Ctx, addr1.String())
	suite.Require().Equal(uint64(100), aggregateVoteCountCreator.GetPositive())
	suite.Require().Equal(uint64(100), aggregateVoteCountReceiver.GetPositive())

	// ============= addr1 votes addr2 ===============
	msgCreateVote = &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr1.String(),
		Count:    100,
		Mode:     1, // cast
	}
	_, err = msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().NoError(err)

	// validate aggregate vote counts after tx
	aggregateVoteCountCreator, _ = keeper.GetAggregateVotesCasted(suite.Ctx, addr1.String())
	aggregateVoteCountReceiver, _ = keeper.GetAggregateVotesReceived(suite.Ctx, addr1.String())
	suite.Require().Equal(uint64(200), aggregateVoteCountCreator.GetPositive())
	suite.Require().Equal(uint64(200), aggregateVoteCountReceiver.GetPositive())

	// validate vote book
	voteBookEntry, _ := keeper.GetVoteBook(suite.Ctx, types.VoteBookIndex(addr1.String(), addr1.String()))
	suite.validateVoteBook(voteBookEntry, 200, 0)

	credibility, _ := keeper.GetCredibility(suite.Ctx, addr1.String())
	suite.Require().Equal("200.00", credibility.Score) // because X is set to 1

	// ============= addr1 un-votes addr2 ===============
	msgCreateVote = &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr1.String(),
		Count:    100,
		Mode:     0, // un-cast
	}
	_, err = msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().NoError(err)

	// validate aggregate vote counts after tx
	aggregateVoteCountCreator, _ = keeper.GetAggregateVotesCasted(suite.Ctx, addr1.String())
	aggregateVoteCountReceiver, _ = keeper.GetAggregateVotesReceived(suite.Ctx, addr1.String())
	suite.Require().Equal(uint64(100), aggregateVoteCountCreator.GetPositive())
	suite.Require().Equal(uint64(100), aggregateVoteCountReceiver.GetPositive())

	credibility, _ = keeper.GetCredibility(suite.Ctx, addr1.String())
	suite.Require().Equal("100.00", credibility.Score) // because X is set to 1
}

func (suite *VotingKeeperTestSuite) Test_Create_Vote_Self_Cast_UnCast_Validator_Valid_Negative() {
	app, ctx, msgServer, keeper := suite.App, sdk.WrapSDKContext(suite.Ctx), suite.MsgServer, suite.App.VotingKeeper
	balances := sdk.NewCoins(testutil.NewMandCoin(10000))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr1)
	app.AccountKeeper.SetAccount(suite.Ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, suite.Ctx, addr1, balances))

	app.OracleKeeper.SetNetworkConstantResult(suite.Ctx, oracletypes.OracleRequestID(1), oracletypes.NetworkConstantResult{Response: "1.0"})
	app.OracleKeeper.SetLastNetworkConstantID(suite.Ctx, oracletypes.OracleRequestID(1))

	// ============= addr1 votes addr2 ===============
	msgCreateVote := &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr1.String(),
		Count:    -100,
		Mode:     1, // cast
	}
	_, err := msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().NoError(err)

	// validate aggregate vote counts after tx
	aggregateVoteCountCreator, _ := keeper.GetAggregateVotesCasted(suite.Ctx, addr1.String())
	aggregateVoteCountReceiver, _ := keeper.GetAggregateVotesReceived(suite.Ctx, addr1.String())
	suite.Require().Equal(uint64(100), aggregateVoteCountCreator.GetNegative())
	suite.Require().Equal(uint64(100), aggregateVoteCountReceiver.GetNegative())

	// ============= addr1 votes addr2 ===============
	msgCreateVote = &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr1.String(),
		Count:    -100,
		Mode:     1, // cast
	}
	_, err = msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().NoError(err)

	// validate aggregate vote counts after tx
	aggregateVoteCountCreator, _ = keeper.GetAggregateVotesCasted(suite.Ctx, addr1.String())
	aggregateVoteCountReceiver, _ = keeper.GetAggregateVotesReceived(suite.Ctx, addr1.String())
	suite.Require().Equal(uint64(200), aggregateVoteCountCreator.GetNegative())
	suite.Require().Equal(uint64(200), aggregateVoteCountReceiver.GetNegative())

	// validate vote book
	voteBookEntry, _ := keeper.GetVoteBook(suite.Ctx, types.VoteBookIndex(addr1.String(), addr1.String()))
	suite.validateVoteBook(voteBookEntry, 0, 200)

	credibility, _ := keeper.GetCredibility(suite.Ctx, addr1.String())
	suite.Require().Equal("-200.00", credibility.Score) // because X is set to 1

	// ============= addr1 votes addr2 (positive votes which will act as un-vote -ve votes) ===============
	msgCreateVote = &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr1.String(),
		Count:    100,
		Mode:     1, // cast
	}
	_, err = msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().NoError(err)

	// validate aggregate vote counts after tx
	aggregateVoteCountCreator, _ = keeper.GetAggregateVotesCasted(suite.Ctx, addr1.String())
	aggregateVoteCountReceiver, _ = keeper.GetAggregateVotesReceived(suite.Ctx, addr1.String())
	suite.Require().Equal(uint64(100), aggregateVoteCountCreator.GetPositive())
	suite.Require().Equal(uint64(100), aggregateVoteCountReceiver.GetPositive())

	voteBookEntry, _ = keeper.GetVoteBook(suite.Ctx, types.VoteBookIndex(addr1.String(), addr1.String()))
	suite.validateVoteBook(voteBookEntry, 100, 200)

	credibility, _ = keeper.GetCredibility(suite.Ctx, addr1.String())
	suite.Require().Equal("-100.00", credibility.Score) // because X is set to 1
}

func (suite *VotingKeeperTestSuite) Test_Create_Vote_Cast_Invalid_Balance() {
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

func (suite *VotingKeeperTestSuite) Test_Create_Vote_UnCast_Validator_Valid_Positive() {
	app, ctx, msgServer, keeper := suite.App, sdk.WrapSDKContext(suite.Ctx), suite.MsgServer, suite.App.VotingKeeper
	balances := sdk.NewCoins(testutil.NewMandCoin(10000))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr1)
	app.AccountKeeper.SetAccount(suite.Ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, suite.Ctx, addr1, balances))

	validator := app.StakingKeeper.GetAllValidators(suite.Ctx)[0]
	valAddr, _ := sdk.ValAddressFromBech32(validator.OperatorAddress)
	addr2, _ := sdk.AccAddressFromHex(hex.EncodeToString(valAddr.Bytes()))

	app.OracleKeeper.SetNetworkConstantResult(suite.Ctx, oracletypes.OracleRequestID(1), oracletypes.NetworkConstantResult{Response: "1.0"})
	app.OracleKeeper.SetLastNetworkConstantID(suite.Ctx, oracletypes.OracleRequestID(1))

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
	aggregateVoteCountCreator, _ := keeper.GetAggregateVotesCasted(suite.Ctx, addr1.String())
	aggregateVoteCountReceiver, _ := keeper.GetAggregateVotesReceived(suite.Ctx, addr2.String())
	suite.Require().Equal(uint64(100), aggregateVoteCountCreator.GetPositive())
	suite.Require().Equal(uint64(100), aggregateVoteCountReceiver.GetPositive())

	credibility, _ := keeper.GetCredibility(suite.Ctx, addr2.String())
	suite.Require().Equal("100.00", credibility.Score) // because X is set to 1

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

	aggregateVoteCountCreator, _ = keeper.GetAggregateVotesCasted(suite.Ctx, addr1.String())
	aggregateVoteCountReceiver, _ = keeper.GetAggregateVotesReceived(suite.Ctx, addr2.String())
	suite.Require().Equal(uint64(50), aggregateVoteCountCreator.GetPositive())
	suite.Require().Equal(uint64(50), aggregateVoteCountReceiver.GetPositive())

	// validate vote book
	voteBookEntry, _ := keeper.GetVoteBook(suite.Ctx, types.VoteBookIndex(addr2.String(), addr1.String()))
	suite.validateVoteBook(voteBookEntry, 50, 0)

	credibility, _ = keeper.GetCredibility(suite.Ctx, addr2.String())
	suite.Require().Equal("50.00", credibility.Score) // because X is set to 1
}

func (suite *VotingKeeperTestSuite) Test_Create_Vote_UnCast_Validator_Valid_Negative() {
	app, ctx, msgServer, keeper := suite.App, sdk.WrapSDKContext(suite.Ctx), suite.MsgServer, suite.App.VotingKeeper
	balances := sdk.NewCoins(testutil.NewMandCoin(10000))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr1)
	app.AccountKeeper.SetAccount(suite.Ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, suite.Ctx, addr1, balances))

	validator := app.StakingKeeper.GetAllValidators(suite.Ctx)[0]
	valAddr, _ := sdk.ValAddressFromBech32(validator.OperatorAddress)
	addr2, _ := sdk.AccAddressFromHex(hex.EncodeToString(valAddr.Bytes()))

	app.OracleKeeper.SetNetworkConstantResult(suite.Ctx, oracletypes.OracleRequestID(1), oracletypes.NetworkConstantResult{Response: "1.0"})
	app.OracleKeeper.SetLastNetworkConstantID(suite.Ctx, oracletypes.OracleRequestID(1))

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
	aggregateVoteCountCreator, _ := keeper.GetAggregateVotesCasted(suite.Ctx, addr1.String())
	aggregateVoteCountReceiver, _ := keeper.GetAggregateVotesReceived(suite.Ctx, addr2.String())
	suite.Require().Equal(uint64(100), aggregateVoteCountCreator.GetNegative())
	suite.Require().Equal(uint64(100), aggregateVoteCountReceiver.GetNegative())

	credibility, _ := keeper.GetCredibility(suite.Ctx, addr2.String())
	suite.Require().Equal("-100.00", credibility.Score) // because X is set to 1

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
	aggregateVoteCountCreator, _ = keeper.GetAggregateVotesCasted(suite.Ctx, addr1.String())
	aggregateVoteCountReceiver, _ = keeper.GetAggregateVotesReceived(suite.Ctx, addr2.String())
	suite.Require().Equal(uint64(50), aggregateVoteCountCreator.GetNegative())
	suite.Require().Equal(uint64(50), aggregateVoteCountReceiver.GetNegative())

	// validate vote book
	voteBookEntry, _ := keeper.GetVoteBook(suite.Ctx, types.VoteBookIndex(addr2.String(), addr1.String()))
	suite.validateVoteBook(voteBookEntry, 0, 50)

	credibility, _ = keeper.GetCredibility(suite.Ctx, addr2.String())
	suite.Require().Equal("-50.00", credibility.Score) // because X is set to 1
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

func (suite *VotingKeeperTestSuite) Test_X_Value_Change_Non_Validator() {
	app, ctx, msgServer, keeper := suite.App, sdk.WrapSDKContext(suite.Ctx), suite.MsgServer, suite.App.VotingKeeper
	balances := sdk.NewCoins(testutil.NewMandCoin(1000000000))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr1)
	app.AccountKeeper.SetAccount(suite.Ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, suite.Ctx, addr1, balances))

	addr2 := sdk.AccAddress("addr2_______________")
	acc2 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr2)
	app.AccountKeeper.SetAccount(suite.Ctx, acc2)

	app.OracleKeeper.SetNetworkConstantResult(suite.Ctx, oracletypes.OracleRequestID(1), oracletypes.NetworkConstantResult{Response: "1.0"})
	app.OracleKeeper.SetLastNetworkConstantID(suite.Ctx, oracletypes.OracleRequestID(1))

	// ============= addr1 votes addr2 ===============
	msgCreateVote := &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    100000000,
		Mode:     1, // cast
	}
	_, err := msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().NoError(err)

	// validate aggregate vote counts after tx
	aggregateVoteCountCreator, _ := keeper.GetAggregateVotesCasted(suite.Ctx, addr1.String())
	aggregateVoteCountReceiver, _ := keeper.GetAggregateVotesReceived(suite.Ctx, addr2.String())
	suite.Require().Equal(uint64(100000000), aggregateVoteCountCreator.GetPositive())
	suite.Require().Equal(uint64(100000000), aggregateVoteCountReceiver.GetPositive())

	// validate vote book
	voteBookEntry, _ := keeper.GetVoteBook(suite.Ctx, types.VoteBookIndex(addr2.String(), addr1.String()))
	suite.validateVoteBook(voteBookEntry, 100000000, 0)

	credibility, _ := keeper.GetCredibility(suite.Ctx, addr2.String())
	suite.Require().Equal("100000000.00", credibility.Score) // because X is set to 1

	app.OracleKeeper.SetNetworkConstantResult(suite.Ctx, oracletypes.OracleRequestID(2), oracletypes.NetworkConstantResult{Response: "0.5"})
	app.OracleKeeper.SetLastNetworkConstantID(suite.Ctx, oracletypes.OracleRequestID(2))

	keeper.UpdateValidatorsCredibility(suite.Ctx)

	credibility, _ = keeper.GetCredibility(suite.Ctx, addr2.String())
	suite.Require().Equal("100000000.00", credibility.Score)

	addr3 := sdk.AccAddress("addr3_______________")
	acc3 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr3)
	app.AccountKeeper.SetAccount(suite.Ctx, acc3)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, suite.Ctx, addr3, balances))

	// ============= addr3 votes addr2 ===============
	msgCreateVote = &types.MsgCreateVote{
		Creator:  addr3.String(),
		Receiver: addr2.String(),
		Count:    100000000,
		Mode:     1, // cast
	}
	_, err = msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().NoError(err)

	// validate aggregate vote counts after tx
	aggregateVoteCountCreator, _ = keeper.GetAggregateVotesCasted(suite.Ctx, addr1.String())
	aggregateVoteCountCreator1, _ := keeper.GetAggregateVotesCasted(suite.Ctx, addr3.String())
	aggregateVoteCountReceiver, _ = keeper.GetAggregateVotesReceived(suite.Ctx, addr2.String())
	suite.Require().Equal(uint64(100000000), aggregateVoteCountCreator.GetPositive())
	suite.Require().Equal(uint64(100000000), aggregateVoteCountCreator1.GetPositive())
	suite.Require().Equal(uint64(200000000), aggregateVoteCountReceiver.GetPositive())

	// validate vote book
	voteBookEntry, _ = keeper.GetVoteBook(suite.Ctx, types.VoteBookIndex(addr2.String(), addr1.String()))
	suite.validateVoteBook(voteBookEntry, 100000000, 0)

	voteBookEntry, _ = keeper.GetVoteBook(suite.Ctx, types.VoteBookIndex(addr2.String(), addr3.String()))
	suite.validateVoteBook(voteBookEntry, 100000000, 0)

	credibility, _ = keeper.GetCredibility(suite.Ctx, addr2.String())
	suite.Require().Equal("20000000.00", credibility.Score) // because X is set to 1
}

func (suite *VotingKeeperTestSuite) Test_X_Value_Change_Validator() {
	app, ctx, msgServer, keeper := suite.App, sdk.WrapSDKContext(suite.Ctx), suite.MsgServer, suite.App.VotingKeeper
	balances := sdk.NewCoins(testutil.NewMandCoin(10000000000))

	addr1 := sdk.AccAddress("addr1_______________")
	acc1 := app.AccountKeeper.NewAccountWithAddress(suite.Ctx, addr1)
	app.AccountKeeper.SetAccount(suite.Ctx, acc1)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, suite.Ctx, addr1, balances))

	validator := app.StakingKeeper.GetAllValidators(suite.Ctx)[0]
	valAddr, _ := sdk.ValAddressFromBech32(validator.OperatorAddress)
	addr2, _ := sdk.AccAddressFromHex(hex.EncodeToString(valAddr.Bytes()))

	app.OracleKeeper.SetNetworkConstantResult(suite.Ctx, oracletypes.OracleRequestID(1), oracletypes.NetworkConstantResult{Response: "1.0"})
	app.OracleKeeper.SetLastNetworkConstantID(suite.Ctx, oracletypes.OracleRequestID(1))

	// ============= addr1 votes addr2 ===============
	msgCreateVote := &types.MsgCreateVote{
		Creator:  addr1.String(),
		Receiver: addr2.String(),
		Count:    100000000,
		Mode:     1, // cast
	}
	_, err := msgServer.CreateVote(ctx, msgCreateVote)
	suite.Require().NoError(err)

	// validate aggregate vote counts after tx
	aggregateVoteCountCreator, _ := keeper.GetAggregateVotesCasted(suite.Ctx, addr1.String())
	aggregateVoteCountReceiver, _ := keeper.GetAggregateVotesReceived(suite.Ctx, addr2.String())
	suite.Require().Equal(uint64(100000000), aggregateVoteCountCreator.GetPositive())
	suite.Require().Equal(uint64(100000000), aggregateVoteCountReceiver.GetPositive())

	// validate vote book
	voteBookEntry, _ := keeper.GetVoteBook(suite.Ctx, types.VoteBookIndex(addr2.String(), addr1.String()))
	suite.validateVoteBook(voteBookEntry, 100000000, 0)

	credibility, _ := keeper.GetCredibility(suite.Ctx, addr2.String())
	suite.Require().Equal("100000000.00", credibility.Score)

	app.OracleKeeper.SetNetworkConstantResult(suite.Ctx, oracletypes.OracleRequestID(2), oracletypes.NetworkConstantResult{Response: "0.5"})
	app.OracleKeeper.SetLastNetworkConstantID(suite.Ctx, oracletypes.OracleRequestID(2))

	credibility, _ = keeper.GetCredibility(suite.Ctx, addr2.String())
	suite.Require().Equal("100000000.00", credibility.Score) // because X is set to 1

	keeper.UpdateValidatorsCredibility(suite.Ctx)

	credibility, _ = keeper.GetCredibility(suite.Ctx, addr2.String())
	suite.Require().Equal("10000000.00", credibility.Score) // because X is set to 1

	app.OracleKeeper.SetNetworkConstantResult(suite.Ctx, oracletypes.OracleRequestID(3), oracletypes.NetworkConstantResult{Response: "0.1"})
	app.OracleKeeper.SetLastNetworkConstantID(suite.Ctx, oracletypes.OracleRequestID(3))

	keeper.UpdateValidatorsCredibility(suite.Ctx)

	credibility, _ = keeper.GetCredibility(suite.Ctx, addr2.String())
	suite.Require().Equal("1584893.19", credibility.Score) // because X is set to 1
}

func (suite *VotingKeeperTestSuite) validateVoteBook(voteBookEntry types.VoteBook, positive uint64, negative uint64) {
	suite.Require().Equal(positive, voteBookEntry.GetPositive())
	suite.Require().Equal(negative, voteBookEntry.GetNegative())
}

func TestVotingKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(VotingKeeperTestSuite))
}
