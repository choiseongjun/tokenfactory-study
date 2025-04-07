package keeper

import (
	"context"

	"tokenfactory/x/crud/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetPost set a specific post in the store from its index
func (k Keeper) SetPost(ctx context.Context, post types.Post) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostKeyPrefix))
	b := k.cdc.MustMarshal(&post)
	store.Set(types.PostKey(
		post.Index,
	), b)
}

// GetPost returns a post from its index
func (k Keeper) GetPost(
	ctx context.Context,
	index string,

) (val types.Post, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostKeyPrefix))

	b := store.Get(types.PostKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePost removes a post from the store
func (k Keeper) RemovePost(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostKeyPrefix))
	store.Delete(types.PostKey(
		index,
	))
}

// GetAllPost returns all post
func (k Keeper) GetAllPost(ctx context.Context) (list []types.Post) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Post
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
