package ssi

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"mande-chain/x/ssi/keeper"
	"mande-chain/x/ssi/types"
)

// InitGenesis initializes the ssi module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetChainNamespace(&ctx, genState.ChainNamespace)
}

// ExportGenesis returns the ssi module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.ChainNamespace = k.GetChainNamespace(&ctx)

	return genesis
}
