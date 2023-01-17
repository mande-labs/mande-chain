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

func createNAggregateVotesReceived(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AggregateVotesReceived {
	items := make([]types.AggregateVotesReceived, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetAggregateVotesReceived(ctx, items[i])
	}
	return items
}

func TestAggregateVotesReceivedGet(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	items := createNAggregateVotesReceived(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAggregateVotesReceived(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAggregateVotesReceivedRemove(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	items := createNAggregateVotesReceived(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAggregateVotesReceived(ctx,
			item.Index,
		)
		_, found := keeper.GetAggregateVotesReceived(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestAggregateVotesReceivedGetAll(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	items := createNAggregateVotesReceived(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAggregateVotesReceived(ctx)),
	)
}
