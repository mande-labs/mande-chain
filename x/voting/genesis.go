package voting

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"mande-chain/x/voting/keeper"
	"mande-chain/x/voting/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the voteBook
	for _, elem := range genState.VoteBookList {
		k.SetVoteBook(ctx, elem)
	}
	// Set all the aggregateVotesCasted
	for _, elem := range genState.AggregateVotesCastedList {
		k.SetAggregateVotesCasted(ctx, elem)
	}
	// Set all the aggregateVotesReceived
	for _, elem := range genState.AggregateVotesReceivedList {
		k.SetAggregateVotesReceived(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.VoteBookList = k.GetAllVoteBook(ctx)
	genesis.AggregateVotesCastedList = k.GetAllAggregateVotesCasted(ctx)
	genesis.AggregateVotesReceivedList = k.GetAllAggregateVotesReceived(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
