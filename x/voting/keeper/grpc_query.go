package keeper

import (
	"github.com/mande-labs/mande-chain/x/voting/types"
)

var _ types.QueryServer = Keeper{}
