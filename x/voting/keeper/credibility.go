package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	oracletypes "mande-chain/x/oracle/types"
	"mande-chain/x/voting/types"
	"math"
	"strconv"
)

// SetCredibility set a specific credibility in the store from its index
func (k Keeper) SetCredibility(ctx sdk.Context, credibility types.Credibility) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CredibilityKeyPrefix))
	b := k.cdc.MustMarshal(&credibility)
	store.Set(types.CredibilityKey(
		credibility.Index,
	), b)
}

// SetCredibility set a specific credibility in the store from its index
func (k Keeper) SetAppliedX(ctx sdk.Context, appliedX types.AppliedX) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AppliedXKeyPrefix))
	b := k.cdc.MustMarshal(&appliedX)
	store.Set(types.AppliedXKey(), b)
}

// GetCredibility returns a credibility from its index
func (k Keeper) GetCredibility(
	ctx sdk.Context,
	index string,

) (val types.Credibility, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CredibilityKeyPrefix))

	b := store.Get(types.CredibilityKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetCredibility returns a credibility from its index
func (k Keeper) GetAppliedX(
	ctx sdk.Context,

) (val types.AppliedX, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AppliedXKeyPrefix))

	b := store.Get(types.AppliedXKey())
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCredibility removes a credibility from the store
func (k Keeper) RemoveCredibility(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CredibilityKeyPrefix))
	store.Delete(types.CredibilityKey(
		index,
	))
}

// GetAllCredibility returns all credibility
func (k Keeper) GetAllCredibility(ctx sdk.Context) (list []types.Credibility) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CredibilityKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Credibility
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) UpdateCredibility(ctx sdk.Context, receiver string, credibility *types.Credibility) {
	lastReqId := k.oracleKeeper.GetLastNetworkConstantID(ctx)
	xResult, _ := k.oracleKeeper.GetNetworkConstantResult(ctx, oracletypes.OracleRequestID(lastReqId))

	x, _ := strconv.ParseFloat(xResult.Response, 64)
	var credScore float64

	// TODO: Implement some caching mechanism to cache cred scores for votes and only compute for the current vote
	iterator := k.GetVoteBookIteratorByPrefix(ctx, receiver)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var voteBook types.VoteBook
		err := k.UnmarshalVoteBook(iterator.Value(), &voteBook)
		if err != nil {
			panic(err)
		}
		ballot := int64(voteBook.Positive - voteBook.Negative)
		voteCount := float64(intAbs2(ballot))
		if ballot < 0 {
			credScore = credScore + math.Pow(voteCount, x)*-1.0
		} else {
			credScore = credScore + math.Pow(voteCount, x)
		}
	}
	credibility.Score = fmt.Sprintf("%.2f", credScore * (math.Pow(1000000,(1-x))))
	credibility.ForX = xResult.Response

	k.SetCredibility(ctx, *credibility)

	return
}
