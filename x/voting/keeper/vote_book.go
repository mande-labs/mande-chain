package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"mande-chain/x/voting/types"
)

// SetVoteBook set a specific voteBook in the store from its index
func (k Keeper) SetVoteBook(ctx sdk.Context, voteBook types.VoteBook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteBookKeyPrefix))
	b := k.cdc.MustMarshal(&voteBook)
	store.Set(types.VoteBookKey(
		voteBook.Index,
	), b)
}

// GetVoteBook returns a voteBook from its index
func (k Keeper) GetVoteBook(
	ctx sdk.Context,
	index string,

) (val types.VoteBook, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteBookKeyPrefix))

	b := store.Get(types.VoteBookKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVoteBook removes a voteBook from the store
func (k Keeper) RemoveVoteBook(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteBookKeyPrefix))
	store.Delete(types.VoteBookKey(
		index,
	))
}

// GetAllVoteBook returns all voteBook
func (k Keeper) GetAllVoteBook(ctx sdk.Context) (list []types.VoteBook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteBookKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VoteBook
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) GetVoteBookIteratorByPrefix(ctx sdk.Context, vbPrefix string) storetypes.Iterator {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteBookKeyPrefix))
	return sdk.KVStorePrefixIterator(store, []byte(vbPrefix))
}

func (k Keeper) UnmarshalVoteBook(bz []byte, voteBook *types.VoteBook) error {
	err := k.cdc.Unmarshal(bz, voteBook)
	if err != nil {
		return err
	}
	return nil
}
