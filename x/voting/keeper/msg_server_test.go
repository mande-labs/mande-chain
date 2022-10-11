package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/mande-labs/mande-chain/testutil/keeper"
	"github.com/mande-labs/mande-chain/x/voting/keeper"
	"github.com/mande-labs/mande-chain/x/voting/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.VotingKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
