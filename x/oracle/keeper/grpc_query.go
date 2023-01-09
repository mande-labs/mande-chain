package keeper

import (
	"mande-chain/x/oracle/types"
)

var _ types.QueryServer = Keeper{}
