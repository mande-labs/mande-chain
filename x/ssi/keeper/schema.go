package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"mande-chain/x/ssi/types"
)

func (k Keeper) GetSchemaCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey and SchemaCountKey (which is "Schema-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SchemaCountKey))
	// Convert the SchemaCountKey to bytes
	byteKey := []byte(types.SchemaCountKey)
	// Get the value of the count
	bz := store.Get(byteKey)
	// Return zero if the count value is not found
	if bz == nil {
		return 0
	}
	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

// Check whether the given Schema already exists in the store
func (k Keeper) HasSchema(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SchemaKey))
	return store.Has([]byte(id))
}

func (k Keeper) SetSchemaCount(ctx sdk.Context, count uint64) {
	// Get the store using storeKey and SchemaCountKey (which is "Schema-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SchemaCountKey))
	// Convert the SchemaCountKey to bytes
	byteKey := []byte(types.SchemaCountKey)
	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	// Set the value of Schema-count- to count
	store.Set(byteKey, bz)
}

func (k Keeper) RegisterSchemaInStore(ctx sdk.Context, schema types.Schema) uint64 {
	// Get the current number of Schemas in the store
	count := k.GetSchemaCount(ctx)
	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SchemaKey))
	// Marshal the Schema into bytes
	schemaBytes := k.cdc.MustMarshal(&schema)
	store.Set([]byte(schema.Id), schemaBytes)
	// Update the Schema count
	k.SetSchemaCount(ctx, count+1)
	return count
}

// Get the schema from store
func (k Keeper) GetSchemaFromStore(ctx sdk.Context, querySchemaId string) []*types.Schema {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SchemaKey))
	var versionNumLengthWithColon int = 4
	var schemas []*types.Schema
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	for ; iterator.Valid(); iterator.Next() {
		var schema types.Schema
		k.cdc.MustUnmarshal(iterator.Value(), &schema)

		if querySchemaId == schema.Id[0:len(schema.Id)-versionNumLengthWithColon] || querySchemaId == schema.Id {
			schemas = append(schemas, &schema)
		}
	}

	return schemas
}
