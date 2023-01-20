package keeper

import (
	"fmt"
	"math"
	"strconv"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"mande-chain/x/voting/types"
	oracletypes "mande-chain/x/oracle/types"
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
	ctx.Logger().Info(fmt.Sprintf("entered UpdateCredibility"))
	lastReqId := k.oracleKeeper.GetLastNetworkConstantID(ctx)
	ctx.Logger().Info(fmt.Sprintf("check lastReqId here: %d", lastReqId))
	xResult, _ := k.oracleKeeper.GetNetworkConstantResult(ctx, oracletypes.OracleRequestID(lastReqId))
	ctx.Logger().Info(fmt.Sprintf("check xResult here: %s", xResult.Response))

	x, _ := strconv.ParseFloat(xResult.Response, 64)
	var credScore float64

	// TODO: Implement some caching mechanism to cache cred scores for votes and only compute for the current vote
	iterator := k.GetVoteBookIteratorByPrefix(ctx, receiver)
	defer iterator.Close()
	ctx.Logger().Info(fmt.Sprintf("iteraction started"))
	for ; iterator.Valid(); iterator.Next() {
		var voteBook types.VoteBook
		err := k.UnmarshalVoteBook(iterator.Value(), &voteBook)
		if err != nil {
			panic(err)
		}
		ctx.Logger().Info(fmt.Sprintf("computing x for votebook idx: %s", voteBook.Index))
		ballot := int64(voteBook.Positive - voteBook.Negative)
		ctx.Logger().Info(fmt.Sprintf("ballot count: %d", ballot))
		voteCount := float64(intAbs2(ballot))
		ctx.Logger().Info(fmt.Sprintf("vote count: %f", voteCount))
		if ballot < 0 {
			credScore = credScore + math.Pow(voteCount, x) * -1.0
		} else {
			credScore = credScore + math.Pow(voteCount, x)
		}

		ctx.Logger().Info(fmt.Sprintf("cred score after computation: %.2f", credScore))
	}
	ctx.Logger().Info(fmt.Sprintf("iteraction ended"))

	ctx.Logger().Info(fmt.Sprintf("final cred score after computation: %.2f", credScore))
	credibility.Score = fmt.Sprintf("%.2f", credScore)
	credibility.ForX = xResult.Response

	k.SetCredibility(ctx, *credibility)

	ctx.Logger().Info(fmt.Sprintf("check x here: %s", xResult))
	ctx.Logger().Info(fmt.Sprintf("check xDec here: %.2f", x))
	ctx.Logger().Info(fmt.Sprintf("check score here: %s", credibility.Score))
	ctx.Logger().Info(fmt.Sprintf("exiting UpdateCredibility"))
	return
}
