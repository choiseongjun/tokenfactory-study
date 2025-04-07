package crud_test

import (
	"testing"

	keepertest "tokenfactory/testutil/keeper"
	"tokenfactory/testutil/nullify"
	crud "tokenfactory/x/crud/module"
	"tokenfactory/x/crud/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PostList: []types.Post{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CrudKeeper(t)
	crud.InitGenesis(ctx, k, genesisState)
	got := crud.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PostList, got.PostList)
	// this line is used by starport scaffolding # genesis/test/assert
}
