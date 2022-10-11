package voting

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mande-labs/mande-chain/x/voting/keeper"
)

// Called every block, burn voting module acc cred coins
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	k.BurnModuleCredCoins(ctx)
}