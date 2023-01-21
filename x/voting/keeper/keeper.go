package keeper

import (
	"encoding/hex"
	"fmt"
	"github.com/tendermint/tendermint/libs/log"
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	oracletypes "mande-chain/x/oracle/types"
	"mande-chain/x/voting/types"
	"github.com/ignite/cli/ignite/pkg/cosmosibckeeper"
)

type (
	Keeper struct {
		*cosmosibckeeper.Keeper
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
		stakingKeeper types.StakingKeeper
		oracleKeeper  types.OracleKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	channelKeeper cosmosibckeeper.ChannelKeeper,
	portKeeper cosmosibckeeper.PortKeeper,
	scopedKeeper cosmosibckeeper.ScopedKeeper,

	accountKeeper types.AccountKeeper, bankKeeper types.BankKeeper, stakingKeeper types.StakingKeeper, oracleKeeper types.OracleKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		Keeper: cosmosibckeeper.NewKeeper(
                types.PortKey,
                storeKey,
                channelKeeper,
                portKeeper,
                scopedKeeper,
        ),

		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		accountKeeper: accountKeeper, bankKeeper: bankKeeper, stakingKeeper: stakingKeeper, oracleKeeper: oracleKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) BurnModuleCredCoins(ctx sdk.Context) {
	acc := k.accountKeeper.GetModuleAddress(types.ModuleName)
	coinsToBurn := k.bankKeeper.GetBalance(ctx, acc, "cred")
	if !coinsToBurn.Amount.IsZero() {
		k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.Coins{sdk.NewCoin("cred", coinsToBurn.Amount)})
	}
}

func (k Keeper) UpdateValidatorsCredibility(ctx sdk.Context) {
	rawAppliedX, _ := k.GetAppliedX(ctx)
	appliedX := rawAppliedX.X
	lastReqId := k.oracleKeeper.GetLastNetworkConstantID(ctx)
	xResult, _ := k.oracleKeeper.GetNetworkConstantResult(ctx, oracletypes.OracleRequestID(lastReqId))
	currentX := xResult.Response

	if appliedX == currentX {
		ctx.Logger().Info(fmt.Sprintf("Credibility update not required"))
		return
	}

	validators := k.stakingKeeper.GetBondedValidatorsByPower(ctx)
	for _, validator := range validators {
		valAddr, _ := sdk.ValAddressFromBech32(validator.OperatorAddress)
		valAccAddr, _ := sdk.AccAddressFromHex(hex.EncodeToString(valAddr.Bytes()))
		valAccAddrStr := valAccAddr.String()
		credibility, _ := k.GetCredibility(ctx, valAccAddrStr)
		beforeCred, _ := strconv.ParseFloat(credibility.Score, 64)
		k.UpdateCredibility(ctx, valAccAddrStr, &credibility)
		afterCred, _ := strconv.ParseFloat(credibility.Score, 64)
		k.delegationHandler(ctx, valAccAddrStr, int64(beforeCred), int64(afterCred))
		ctx.Logger().Info(fmt.Sprintf("Credibility updated for: %s", valAccAddrStr))
	}

	rawAppliedX.X = xResult.Response
	k.SetAppliedX(ctx, rawAppliedX)
}
