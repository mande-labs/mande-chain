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

func createNAggregateVoteCount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AggregateVoteCount {
	items := make([]types.AggregateVoteCount, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetAggregateVoteCount(ctx, items[i])
	}
	return items
}

func TestAggregateVoteCountGet(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	items := createNAggregateVoteCount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAggregateVoteCount(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAggregateVoteCountRemove(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	items := createNAggregateVoteCount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAggregateVoteCount(ctx,
			item.Index,
		)
		_, found := keeper.GetAggregateVoteCount(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestAggregateVoteCountGetAll(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	items := createNAggregateVoteCount(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAggregateVoteCount(ctx)),
	)
}
