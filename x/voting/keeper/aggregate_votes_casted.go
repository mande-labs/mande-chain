package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"mande-chain/x/voting/types"
)

// SetAggregateVotesCasted set a specific aggregateVotesCasted in the store from its index
func (k Keeper) SetAggregateVotesCasted(ctx sdk.Context, aggregateVotesCasted types.AggregateVotesCasted) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AggregateVotesCastedKeyPrefix))
	b := k.cdc.MustMarshal(&aggregateVotesCasted)
	store.Set(types.AggregateVotesCastedKey(
		aggregateVotesCasted.Index,
	), b)
}

// GetAggregateVotesCasted returns a aggregateVotesCasted from its index
func (k Keeper) GetAggregateVotesCasted(
	ctx sdk.Context,
	index string,

) (val types.AggregateVotesCasted, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AggregateVotesCastedKeyPrefix))

	b := store.Get(types.AggregateVotesCastedKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAggregateVotesCasted removes a aggregateVotesCasted from the store
func (k Keeper) RemoveAggregateVotesCasted(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AggregateVotesCastedKeyPrefix))
	store.Delete(types.AggregateVotesCastedKey(
		index,
	))
}

// GetAllAggregateVotesCasted returns all aggregateVotesCasted
func (k Keeper) GetAllAggregateVotesCasted(ctx sdk.Context) (list []types.AggregateVotesCasted) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AggregateVotesCastedKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AggregateVotesCasted
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) ReconcileCreatorAggregatedVotes(msg *types.MsgCreateVote, aggregateVoteCreatorCount *types.AggregateVotesCasted) {
	voteCount := intAbs(msg.Count)
	switch msg.Mode {
	case 0: // uncast
		if msg.Count < 0 {
			aggregateVoteCreatorCount.Negative -= voteCount
		} else {
			aggregateVoteCreatorCount.Positive -= voteCount
		}
	case 1: // cast
		if msg.Count < 0 {
			aggregateVoteCreatorCount.Negative += voteCount
		} else {
			aggregateVoteCreatorCount.Positive += voteCount
		}
	}
}
