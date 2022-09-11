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

func (k Keeper) VoteBookAll(c context.Context, req *types.QueryAllVoteBookRequest) (*types.QueryAllVoteBookResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var voteBooks []types.VoteBook
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	voteBookStore := prefix.NewStore(store, types.KeyPrefix(types.VoteBookKeyPrefix))

	pageRes, err := query.Paginate(voteBookStore, req.Pagination, func(key []byte, value []byte) error {
		var voteBook types.VoteBook
		if err := k.cdc.Unmarshal(value, &voteBook); err != nil {
			return err
		}

		voteBooks = append(voteBooks, voteBook)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVoteBookResponse{VoteBook: voteBooks, Pagination: pageRes}, nil
}

func (k Keeper) VoteBook(c context.Context, req *types.QueryGetVoteBookRequest) (*types.QueryGetVoteBookResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetVoteBook(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetVoteBookResponse{VoteBook: val}, nil
}
