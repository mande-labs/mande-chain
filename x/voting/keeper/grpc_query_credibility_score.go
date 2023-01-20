package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mande-chain/x/voting/types"
	oracletypes "mande-chain/x/oracle/types"
)

func (k Keeper) CredibilityScore(goCtx context.Context, req *types.QueryCredibilityScoreRequest) (*types.QueryCredibilityScoreResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	credibility, found := k.GetCredibility(ctx, req.Address)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	rawAppliedX, _ := k.GetAppliedX(ctx)
	appliedX := rawAppliedX.X
	lastReqId := k.oracleKeeper.GetLastNetworkConstantID(ctx)
	xResult, _ := k.oracleKeeper.GetNetworkConstantResult(ctx, oracletypes.OracleRequestID(lastReqId))
	currentX := xResult.Response

	updated := false
	if appliedX == currentX {
		updated = true
	}

	return &types.QueryCredibilityScoreResponse{Score: credibility.Score, Updated: updated}, nil
}
