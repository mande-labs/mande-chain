package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"mande-chain/x/voting/types"
	"strconv"
)

func (k Keeper) CreateVote(ctx sdk.Context, msg *types.MsgCreateVote) error {
	aggregateVoteCreatorCount, found := k.GetAggregateVotesCasted(ctx, msg.Creator)
	if !found {
		aggregateVoteCreatorCount.Index = msg.Creator
		aggregateVoteCreatorCount.Positive = 0
		aggregateVoteCreatorCount.Negative = 0
	}

	credibility, found := k.GetCredibility(ctx, msg.Receiver)
	if !found {
		credibility.Index = msg.Receiver
		credibility.Score = "0"
		credibility.ForX = "0"
	}

	switch msg.Mode {
	case 0:
		if err := k.uncastVote(ctx, msg, &aggregateVoteCreatorCount, credibility); err != nil {
			return err
		}
	case 1:
		if err := k.castVote(ctx, msg, &aggregateVoteCreatorCount, credibility); err != nil {
			return err
		}
	default:
		return sdkerrors.Wrapf(types.ErrInvalidVotingMode, "mode - (%d)", msg.Mode)
	}

	return nil
}

func (k Keeper) uncastVote(ctx sdk.Context, msg *types.MsgCreateVote, aggregateVotesCreatorCount *types.AggregateVotesCasted, credibility types.Credibility) error {
	voteBookIndex := types.VoteBookIndex(msg.Receiver, msg.Creator)
	voteBookEntry, found := k.GetVoteBook(ctx, voteBookIndex)
	if !found {
		return sdkerrors.Wrap(types.ErrNoVoteRecord, msg.Receiver)
	}

	aggregateVotesReceiverCount, found := k.GetAggregateVotesReceived(ctx, msg.Receiver)
	if !found {
		return sdkerrors.Wrap(types.ErrNoVotesCasted, msg.Receiver)
	}

	voteCount := intAbs(msg.Count)
	if msg.Count < 0 {
		if voteBookEntry.Negative >= voteCount {
			voteBookEntry.Negative -= voteCount
		} else {
			return sdkerrors.Wrapf(types.ErrNegativeVotesCastedLimit, "count - (%d)", msg.Count)
		}
	} else {
		if voteBookEntry.Positive >= voteCount {
			voteBookEntry.Positive -= voteCount
		} else {
			return sdkerrors.Wrapf(types.ErrPositiveVotesCastedLimit, "count = (%d)", msg.Count)
		}
	}

	k.SetVoteBook(ctx, voteBookEntry)

	k.ReconcileCreatorAggregatedVotes(msg, aggregateVotesCreatorCount)
	k.ReconcileReceiverAggregatedVotes(msg, &aggregateVotesReceiverCount)
	k.SetAggregateVotesCasted(ctx, *aggregateVotesCreatorCount)
	k.SetAggregateVotesReceived(ctx, aggregateVotesReceiverCount)

	beforeCred, _ := strconv.ParseFloat(credibility.Score, 64)
	k.UpdateCredibility(ctx, aggregateVotesReceiverCount, &credibility)
	afterCred, _ := strconv.ParseFloat(credibility.Score, 64)

	err := k.undelegateStakeAndUnlockMand(ctx, msg, int64(beforeCred), int64(afterCred))
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeVoteUncasted,
			sdk.NewAttribute(types.AttributeKeyCaster, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
			sdk.NewAttribute(sdk.AttributeKeyAmount, strconv.Itoa(int(msg.Count))),
		),
	)

	return nil
}

func (k Keeper) castVote(ctx sdk.Context, msg *types.MsgCreateVote, aggregateVotesCreatorCount *types.AggregateVotesCasted, credibility types.Credibility) error {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	mandBalance := k.bankKeeper.GetBalance(ctx, creator, "mand").Amount.Uint64()

	if mandBalance < intAbs(msg.Count) {
		return sdkerrors.Wrapf(types.ErrNotEnoughMand, "count - (%d)", msg.Count)
	}

	voteBookIndex := types.VoteBookIndex(msg.Receiver, msg.Creator)
	voteBookEntry, found := k.GetVoteBook(ctx, voteBookIndex)
	if !found {
		voteBookEntry.Index = voteBookIndex
		voteBookEntry.Caster = msg.Creator
		voteBookEntry.Receiver = msg.Receiver
		voteBookEntry.Positive = 0
		voteBookEntry.Negative = 0
	}

	voteCount := intAbs(msg.Count)
	if msg.Count < 0 {
		voteBookEntry.Negative += voteCount
	} else {
		voteBookEntry.Positive += voteCount
	}

	aggregateVotesReceiverCount, found := k.GetAggregateVotesReceived(ctx, msg.Receiver)
	if !found {
		aggregateVotesReceiverCount.Index = msg.Receiver
		aggregateVotesReceiverCount.Positive = 0
		aggregateVotesReceiverCount.Negative = 0
	}

	k.SetVoteBook(ctx, voteBookEntry)

	k.ReconcileCreatorAggregatedVotes(msg, aggregateVotesCreatorCount)
	k.ReconcileReceiverAggregatedVotes(msg, &aggregateVotesReceiverCount)
	k.SetAggregateVotesCasted(ctx, *aggregateVotesCreatorCount)
	k.SetAggregateVotesReceived(ctx, aggregateVotesReceiverCount)

	beforeCred, _ := strconv.ParseFloat(credibility.Score, 64)
	ctx.Logger().Info(fmt.Sprintf("check beforeCred here: %f", beforeCred))
	ctx.Logger().Info(fmt.Sprintf("check credibility score before updating here: %s", credibility.Score))
	k.UpdateCredibility(ctx, aggregateVotesReceiverCount, &credibility)
	ctx.Logger().Info(fmt.Sprintf("check credibility after updating here: %s", credibility.Score))
	afterCred, _ := strconv.ParseFloat(credibility.Score, 64)
	// afterCred := fmt.Sprintf("%d", credibility.Score)
	ctx.Logger().Info(fmt.Sprintf("check afterCred here: %f", afterCred))

	err := k.lockMandAndDelegateStake(ctx, msg, int64(beforeCred), int64(afterCred))
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeVoteCasted,
			sdk.NewAttribute(types.AttributeKeyCaster, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
			sdk.NewAttribute(sdk.AttributeKeyAmount, strconv.Itoa(int(msg.Count))),
		),
	)

	return nil
}
