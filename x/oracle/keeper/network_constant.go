package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gogotypes "github.com/gogo/protobuf/types"
	"mande-chain/x/oracle/types"
)

// SetNetworkConstantResult saves the NetworkConstant result
func (k Keeper) SetNetworkConstantResult(ctx sdk.Context, requestID types.OracleRequestID, result types.NetworkConstantResult) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.NetworkConstantResultStoreKey(requestID), k.cdc.MustMarshal(&result))
}

// GetNetworkConstantResult returns the NetworkConstant by requestId
func (k Keeper) GetNetworkConstantResult(ctx sdk.Context, id types.OracleRequestID) (types.NetworkConstantResult, error) {
	bz := ctx.KVStore(k.storeKey).Get(types.NetworkConstantResultStoreKey(id))
	if bz == nil {
		return types.NetworkConstantResult{}, sdkerrors.Wrapf(types.ErrSample,
			"GetResult: Result for request ID %d is not available.", id,
		)
	}
	var result types.NetworkConstantResult
	k.cdc.MustUnmarshal(bz, &result)
	return result, nil
}

// GetLastNetworkConstantID return the id from the last NetworkConstant request
func (k Keeper) GetLastNetworkConstantID(ctx sdk.Context) int64 {
	bz := ctx.KVStore(k.storeKey).Get(types.KeyPrefix(types.LastNetworkConstantIDKey))
	intV := gogotypes.Int64Value{}
	k.cdc.MustUnmarshalLengthPrefixed(bz, &intV)
	return intV.GetValue()
}

// SetLastNetworkConstantID saves the id from the last NetworkConstant request
func (k Keeper) SetLastNetworkConstantID(ctx sdk.Context, id types.OracleRequestID) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefix(types.LastNetworkConstantIDKey),
		k.cdc.MustMarshalLengthPrefixed(&gogotypes.Int64Value{Value: int64(id)}))
}
