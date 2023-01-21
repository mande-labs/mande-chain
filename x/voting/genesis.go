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
	// Set all the credibility
	for _, elem := range genState.CredibilityList {
		k.SetCredibility(ctx, elem)
	}
	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, genState.PortId) {
	        // module binds to the port on InitChain
	        // and claims the returned capability
	        err := k.BindPort(ctx, genState.PortId)
	        if err != nil {
	                panic("could not claim port capability: " + err.Error())
	        }
	}

	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.PortId = k.GetPort(ctx)

	genesis.VoteBookList = k.GetAllVoteBook(ctx)
	genesis.AggregateVotesCastedList = k.GetAllAggregateVotesCasted(ctx)
	genesis.AggregateVotesReceivedList = k.GetAllAggregateVotesReceived(ctx)
	genesis.CredibilityList = k.GetAllCredibility(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
