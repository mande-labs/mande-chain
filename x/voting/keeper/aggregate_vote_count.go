package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"mande-chain/x/voting/types"
)

// SetAggregateVoteCount set a specific aggregateVoteCount in the store from its index
func (k Keeper) SetAggregateVoteCount(ctx sdk.Context, aggregateVoteCount types.AggregateVoteCount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AggregateVoteCountKeyPrefix))
	b := k.cdc.MustMarshal(&aggregateVoteCount)
	store.Set(types.AggregateVoteCountKey(
		aggregateVoteCount.Index,
	), b)
}

// GetAggregateVoteCount returns a aggregateVoteCount from its index
func (k Keeper) GetAggregateVoteCount(
	ctx sdk.Context,
	index string,

) (val types.AggregateVoteCount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AggregateVoteCountKeyPrefix))

	b := store.Get(types.AggregateVoteCountKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAggregateVoteCount removes a aggregateVoteCount from the store
func (k Keeper) RemoveAggregateVoteCount(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AggregateVoteCountKeyPrefix))
	store.Delete(types.AggregateVoteCountKey(
		index,
	))
}

// GetAllAggregateVoteCount returns all aggregateVoteCount
func (k Keeper) GetAllAggregateVoteCount(ctx sdk.Context) (list []types.AggregateVoteCount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AggregateVoteCountKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AggregateVoteCount
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) ReconcileAggregatedVotes(msg *types.MsgCreateVote, aggregateVoteCreatorCount *types.AggregateVoteCount, aggregateVoteReceiverCount *types.AggregateVoteCount) {
	voteCount := intAbs(msg.Count)
	switch msg.Mode {
	case 0: // uncast
		aggregateVoteCreatorCount.Casted -= voteCount
		if msg.Count < 0 {
			aggregateVoteReceiverCount.NegativeReceived -= voteCount
		} else {
			aggregateVoteReceiverCount.PositiveReceived -= voteCount
		}
	case 1: // cast
		aggregateVoteCreatorCount.Casted += voteCount
		if msg.Count < 0 {
			aggregateVoteReceiverCount.NegativeReceived += voteCount
		} else {
			aggregateVoteReceiverCount.PositiveReceived += voteCount
		}
	}
}
