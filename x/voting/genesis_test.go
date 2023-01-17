package voting_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "mande-chain/testutil/keeper"
	"mande-chain/testutil/nullify"
	"mande-chain/x/voting"
	"mande-chain/x/voting/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		VoteBookList: []types.VoteBook{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		AggregateVotesCastedList: []types.AggregateVotesCasted{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		AggregateVotesReceivedList: []types.AggregateVotesReceived{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VotingKeeper(t)
	voting.InitGenesis(ctx, *k, genesisState)
	got := voting.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.VoteBookList, got.VoteBookList)
	require.ElementsMatch(t, genesisState.AggregateVotesCastedList, got.AggregateVotesCastedList)
	require.ElementsMatch(t, genesisState.AggregateVotesReceivedList, got.AggregateVotesReceivedList)
	// this line is used by starport scaffolding # genesis/test/assert
}
