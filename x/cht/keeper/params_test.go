package keeper_test

import (
	"testing"

	testkeeper "github.com/chronicnetwork/cht/testutil/keeper"
	"github.com/chronicnetwork/cht/x/cht/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ChtKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
