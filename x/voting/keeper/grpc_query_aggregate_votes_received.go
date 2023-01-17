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

func (k Keeper) AggregateVotesReceivedAll(c context.Context, req *types.QueryAllAggregateVotesReceivedRequest) (*types.QueryAllAggregateVotesReceivedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var aggregateVotesReceiveds []types.AggregateVotesReceived
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	aggregateVotesReceivedStore := prefix.NewStore(store, types.KeyPrefix(types.AggregateVotesReceivedKeyPrefix))

	pageRes, err := query.Paginate(aggregateVotesReceivedStore, req.Pagination, func(key []byte, value []byte) error {
		var aggregateVotesReceived types.AggregateVotesReceived
		if err := k.cdc.Unmarshal(value, &aggregateVotesReceived); err != nil {
			return err
		}

		aggregateVotesReceiveds = append(aggregateVotesReceiveds, aggregateVotesReceived)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAggregateVotesReceivedResponse{AggregateVotesReceived: aggregateVotesReceiveds, Pagination: pageRes}, nil
}

func (k Keeper) AggregateVotesReceived(c context.Context, req *types.QueryGetAggregateVotesReceivedRequest) (*types.QueryGetAggregateVotesReceivedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetAggregateVotesReceived(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetAggregateVotesReceivedResponse{AggregateVotesReceived: val}, nil
}
