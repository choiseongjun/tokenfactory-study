package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "tokenfactory/testutil/keeper"
	"tokenfactory/x/crud/keeper"
	"tokenfactory/x/crud/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestPostMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.CrudKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreatePost{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreatePost(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetPost(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestPostMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdatePost
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdatePost{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdatePost{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdatePost{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.CrudKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreatePost{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreatePost(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdatePost(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetPost(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestPostMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeletePost
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeletePost{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeletePost{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeletePost{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.CrudKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreatePost(ctx, &types.MsgCreatePost{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeletePost(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetPost(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
