package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mande-chain/x/voting/types"
)

func (k Keeper) AggregateVoteCountAll(c context.Context, req *types.QueryAllAggregateVoteCountRequest) (*types.QueryAllAggregateVoteCountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var aggregateVoteCounts []types.AggregateVoteCount
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	aggregateVoteCountStore := prefix.NewStore(store, types.KeyPrefix(types.AggregateVoteCountKeyPrefix))

	pageRes, err := query.Paginate(aggregateVoteCountStore, req.Pagination, func(key []byte, value []byte) error {
		var aggregateVoteCount types.AggregateVoteCount
		if err := k.cdc.Unmarshal(value, &aggregateVoteCount); err != nil {
			return err
		}

		aggregateVoteCounts = append(aggregateVoteCounts, aggregateVoteCount)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAggregateVoteCountResponse{AggregateVoteCount: aggregateVoteCounts, Pagination: pageRes}, nil
}

func (k Keeper) AggregateVoteCount(c context.Context, req *types.QueryGetAggregateVoteCountRequest) (*types.QueryGetAggregateVoteCountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetAggregateVoteCount(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetAggregateVoteCountResponse{AggregateVoteCount: val}, nil
}
