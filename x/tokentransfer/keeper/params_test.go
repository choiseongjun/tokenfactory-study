package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "tokenfactory/testutil/keeper"
	"tokenfactory/x/tokentransfer/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.TokentransferKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
