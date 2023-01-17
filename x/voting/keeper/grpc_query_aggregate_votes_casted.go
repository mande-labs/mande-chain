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

func (k Keeper) AggregateVotesCastedAll(c context.Context, req *types.QueryAllAggregateVotesCastedRequest) (*types.QueryAllAggregateVotesCastedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var aggregateVotesCasteds []types.AggregateVotesCasted
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	aggregateVotesCastedStore := prefix.NewStore(store, types.KeyPrefix(types.AggregateVotesCastedKeyPrefix))

	pageRes, err := query.Paginate(aggregateVotesCastedStore, req.Pagination, func(key []byte, value []byte) error {
		var aggregateVotesCasted types.AggregateVotesCasted
		if err := k.cdc.Unmarshal(value, &aggregateVotesCasted); err != nil {
			return err
		}

		aggregateVotesCasteds = append(aggregateVotesCasteds, aggregateVotesCasted)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAggregateVotesCastedResponse{AggregateVotesCasted: aggregateVotesCasteds, Pagination: pageRes}, nil
}

func (k Keeper) AggregateVotesCasted(c context.Context, req *types.QueryGetAggregateVotesCastedRequest) (*types.QueryGetAggregateVotesCastedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetAggregateVotesCasted(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetAggregateVotesCastedResponse{AggregateVotesCasted: val}, nil
}
