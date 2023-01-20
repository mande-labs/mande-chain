package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"mande-chain/x/voting/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
				CredibilityList: []types.Credibility{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated voteBook",
			genState: &types.GenesisState{
				VoteBookList: []types.VoteBook{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated aggregateVotesCasted",
			genState: &types.GenesisState{
				AggregateVotesCastedList: []types.AggregateVotesCasted{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated aggregateVotesReceived",
			genState: &types.GenesisState{
				AggregateVotesReceivedList: []types.AggregateVotesReceived{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated credibility",
			genState: &types.GenesisState{
				CredibilityList: []types.Credibility{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
