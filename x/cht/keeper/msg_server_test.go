package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/chronicnetwork/cht/testutil/keeper"
	"github.com/chronicnetwork/cht/x/cht/keeper"
	"github.com/chronicnetwork/cht/x/cht/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ChtKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
