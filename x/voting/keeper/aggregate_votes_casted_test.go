package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "mande-chain/testutil/keeper"
	"mande-chain/testutil/nullify"
	"mande-chain/x/voting/keeper"
	"mande-chain/x/voting/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNAggregateVotesCasted(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AggregateVotesCasted {
	items := make([]types.AggregateVotesCasted, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetAggregateVotesCasted(ctx, items[i])
	}
	return items
}

func TestAggregateVotesCastedGet(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	items := createNAggregateVotesCasted(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAggregateVotesCasted(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAggregateVotesCastedRemove(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	items := createNAggregateVotesCasted(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAggregateVotesCasted(ctx,
			item.Index,
		)
		_, found := keeper.GetAggregateVotesCasted(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestAggregateVotesCastedGetAll(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	items := createNAggregateVotesCasted(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAggregateVotesCasted(ctx)),
	)
}
