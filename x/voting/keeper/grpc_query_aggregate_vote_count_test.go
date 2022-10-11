package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/mande-labs/mande-chain/testutil/keeper"
	"github.com/mande-labs/mande-chain/testutil/nullify"
	"github.com/mande-labs/mande-chain/x/voting/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestAggregateVoteCountQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAggregateVoteCount(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetAggregateVoteCountRequest
		response *types.QueryGetAggregateVoteCountResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetAggregateVoteCountRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetAggregateVoteCountResponse{AggregateVoteCount: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetAggregateVoteCountRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetAggregateVoteCountResponse{AggregateVoteCount: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetAggregateVoteCountRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.AggregateVoteCount(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestAggregateVoteCountQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.VotingKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAggregateVoteCount(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllAggregateVoteCountRequest {
		return &types.QueryAllAggregateVoteCountRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AggregateVoteCountAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.AggregateVoteCount), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.AggregateVoteCount),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AggregateVoteCountAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.AggregateVoteCount), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.AggregateVoteCount),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.AggregateVoteCountAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.AggregateVoteCount),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.AggregateVoteCountAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
