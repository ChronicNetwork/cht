package keeper

import (
	"github.com/chronicnetwork/cht/x/cht/types"
)

var _ types.QueryServer = Keeper{}
