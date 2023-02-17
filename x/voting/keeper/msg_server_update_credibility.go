package keeper

import (
	"context"
	"strconv"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	sdk "github.com/cosmos/cosmos-sdk/types"
	oracletypes "mande-chain/x/oracle/types"
	"mande-chain/x/voting/types"
)

func (k msgServer) UpdateCredibility(goCtx context.Context, msg *types.MsgUpdateCredibility) (*types.MsgUpdateCredibilityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return &types.MsgUpdateCredibilityResponse{}, err
	}

	credibility, found := k.Keeper.GetCredibility(ctx, msg.Address)
	if !found {
		return &types.MsgUpdateCredibilityResponse{}, status.Error(codes.NotFound, "not found")
	}

	lastReqId := k.Keeper.oracleKeeper.GetLastNetworkConstantID(ctx)
	xResult, _ := k.Keeper.oracleKeeper.GetNetworkConstantResult(ctx, oracletypes.OracleRequestID(lastReqId))
	currentX := xResult.Response

	if credibility.ForX == currentX {
		return &types.MsgUpdateCredibilityResponse{}, nil
	}

	beforeCred, _ := strconv.ParseFloat(credibility.Score, 64)
	k.Keeper.UpdateCredibility(ctx, msg.Address, &credibility)
	afterCred, _ := strconv.ParseFloat(credibility.Score, 64)
	k.Keeper.delegationHandler(ctx, msg.Address, int64(beforeCred), int64(afterCred))

	return &types.MsgUpdateCredibilityResponse{}, nil
}
