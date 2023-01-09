package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"mande-chain/x/oracle/types"
)

// NetworkConstantResult returns the NetworkConstant result by RequestId
func (k Keeper) NetworkConstantResult(c context.Context, req *types.QueryNetworkConstantRequest) (*types.QueryNetworkConstantResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	result, err := k.GetNetworkConstantResult(ctx, types.OracleRequestID(req.RequestId))
	if err != nil {
		return nil, err
	}
	return &types.QueryNetworkConstantResponse{Result: &result}, nil
}

// LastNetworkConstantId returns the last NetworkConstant request Id
func (k Keeper) LastNetworkConstantId(c context.Context, req *types.QueryLastNetworkConstantIdRequest) (*types.QueryLastNetworkConstantIdResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	id := k.GetLastNetworkConstantID(ctx)
	return &types.QueryLastNetworkConstantIdResponse{RequestId: id}, nil
}
