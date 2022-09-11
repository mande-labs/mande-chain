package keeper

import (
	"mande-chain/x/voting/types"
)

var _ types.QueryServer = Keeper{}
