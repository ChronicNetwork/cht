package cht_test

import (
	"testing"

	keepertest "github.com/chronicnetwork/cht/testutil/keeper"
	"github.com/chronicnetwork/cht/testutil/nullify"
	"github.com/chronicnetwork/cht/x/cht"
	"github.com/chronicnetwork/cht/x/cht/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ChtKeeper(t)
	cht.InitGenesis(ctx, *k, genesisState)
	got := cht.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
