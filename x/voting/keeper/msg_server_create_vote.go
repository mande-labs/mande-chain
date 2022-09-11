package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"mande-chain/x/voting/types"
)

func (k msgServer) CreateVote(goCtx context.Context, msg *types.MsgCreateVote) (*types.MsgCreateVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return &types.MsgCreateVoteResponse{}, err
	}

	if err := k.Keeper.CreateVote(ctx, msg); err != nil {
		return &types.MsgCreateVoteResponse{}, err
	}

	return &types.MsgCreateVoteResponse{}, nil
}
