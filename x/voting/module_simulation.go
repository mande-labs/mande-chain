package voting

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"mande-chain/testutil/sample"
	votingsimulation "mande-chain/x/voting/simulation"
	"mande-chain/x/voting/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = votingsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateVote = "op_weight_msg_create_vote"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateVote int = 100

	opWeightMsgUpdateCredibility = "op_weight_msg_update_credibility"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateCredibility int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	votingGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&votingGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateVote int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateVote, &weightMsgCreateVote, nil,
		func(_ *rand.Rand) {
			weightMsgCreateVote = defaultWeightMsgCreateVote
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateVote,
		votingsimulation.SimulateMsgCreateVote(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateCredibility int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateCredibility, &weightMsgUpdateCredibility, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateCredibility = defaultWeightMsgUpdateCredibility
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateCredibility,
		votingsimulation.SimulateMsgUpdateCredibility(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
