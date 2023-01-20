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

func createNCredibility(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Credibility {
	items := make([]types.Credibility, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetCredibility(ctx, items[i])
	}
	return items
}

func TestCredibilityGet(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	items := createNCredibility(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCredibility(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestCredibilityRemove(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	items := createNCredibility(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCredibility(ctx,
			item.Index,
		)
		_, found := keeper.GetCredibility(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestCredibilityGetAll(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	items := createNCredibility(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCredibility(ctx)),
	)
}
