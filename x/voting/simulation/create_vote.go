package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"mande-chain/x/voting/keeper"
	"mande-chain/x/voting/types"
)

func SimulateMsgCreateVote(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateVote{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateVote simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateVote simulation not implemented"), nil, nil
	}
}
