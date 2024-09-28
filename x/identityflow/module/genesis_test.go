package identityflow_test

import (
	"testing"

	keepertest "identityflow/testutil/keeper"
	"identityflow/testutil/nullify"
	identityflow "identityflow/x/identityflow/module"
	"identityflow/x/identityflow/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IdentityflowKeeper(t)
	identityflow.InitGenesis(ctx, k, genesisState)
	got := identityflow.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
