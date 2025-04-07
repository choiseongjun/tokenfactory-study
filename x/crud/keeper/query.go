package keeper

import (
	"context"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"tokenfactory/x/crud/types"
)

var _ types.QueryServer = Keeper{}

// PostHistory implements the Query/PostHistory RPC method
func (k Keeper) PostHistory(ctx context.Context, req *types.QueryPostHistoryRequest) (*types.QueryPostHistoryResponse, error) {
	if req == nil {
		return nil, types.ErrInvalidRequest
	}

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte(types.PostKeyPrefix))

	var posts []types.Post
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var post types.Post
		if err := k.cdc.Unmarshal(value, &post); err != nil {
			return err
		}
		if post.Index == req.Index {
			posts = append(posts, post)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &types.QueryPostHistoryResponse{
		Post:       posts,
		Pagination: pageRes,
	}, nil
}
