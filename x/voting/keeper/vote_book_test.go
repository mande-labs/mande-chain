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

func createNVoteBook(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.VoteBook {
	items := make([]types.VoteBook, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetVoteBook(ctx, items[i])
	}
	return items
}

func TestVoteBookGet(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	items := createNVoteBook(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetVoteBook(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestVoteBookRemove(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	items := createNVoteBook(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVoteBook(ctx,
			item.Index,
		)
		_, found := keeper.GetVoteBook(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestVoteBookGetAll(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	items := createNVoteBook(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVoteBook(ctx)),
	)
}
