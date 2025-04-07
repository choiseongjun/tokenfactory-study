package tokentransfer_test

import (
	"testing"

	keepertest "tokenfactory/testutil/keeper"
	"tokenfactory/testutil/nullify"
	tokentransfer "tokenfactory/x/tokentransfer/module"
	"tokenfactory/x/tokentransfer/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TokentransferKeeper(t)
	tokentransfer.InitGenesis(ctx, k, genesisState)
	got := tokentransfer.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
