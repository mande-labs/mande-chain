package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"mande-chain/x/voting/types"
)

// SetAggregateVotesReceived set a specific aggregateVotesReceived in the store from its index
func (k Keeper) SetAggregateVotesReceived(ctx sdk.Context, aggregateVotesReceived types.AggregateVotesReceived) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AggregateVotesReceivedKeyPrefix))
	b := k.cdc.MustMarshal(&aggregateVotesReceived)
	store.Set(types.AggregateVotesReceivedKey(
		aggregateVotesReceived.Index,
	), b)
}

// GetAggregateVotesReceived returns a aggregateVotesReceived from its index
func (k Keeper) GetAggregateVotesReceived(
	ctx sdk.Context,
	index string,

) (val types.AggregateVotesReceived, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AggregateVotesReceivedKeyPrefix))

	b := store.Get(types.AggregateVotesReceivedKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAggregateVotesReceived removes a aggregateVotesReceived from the store
func (k Keeper) RemoveAggregateVotesReceived(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AggregateVotesReceivedKeyPrefix))
	store.Delete(types.AggregateVotesReceivedKey(
		index,
	))
}

// GetAllAggregateVotesReceived returns all aggregateVotesReceived
func (k Keeper) GetAllAggregateVotesReceived(ctx sdk.Context) (list []types.AggregateVotesReceived) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AggregateVotesReceivedKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AggregateVotesReceived
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) ReconcileReceiverAggregatedVotes(msg *types.MsgCreateVote, aggregateVoteReceiverCount *types.AggregateVotesReceived) {
	voteCount := intAbs(msg.Count)
	switch msg.Mode {
	case 0: // uncast
		if msg.Count < 0 {
			aggregateVoteReceiverCount.Negative -= voteCount
		} else {
			aggregateVoteReceiverCount.Positive -= voteCount
		}
	case 1: // cast
		if msg.Count < 0 {
			aggregateVoteReceiverCount.Negative += voteCount
		} else {
			aggregateVoteReceiverCount.Positive += voteCount
		}
	}
}
