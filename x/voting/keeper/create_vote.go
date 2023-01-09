package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"mande-chain/x/voting/types"
	"strconv"
)

func (k Keeper) CreateVote(ctx sdk.Context, msg *types.MsgCreateVote) error {
	aggregateVoteCreatorCount, found := k.GetAggregateVoteCount(ctx, msg.Creator)
	if !found {
		aggregateVoteCreatorCount.Index = msg.Creator
		aggregateVoteCreatorCount.Casted = 0
		aggregateVoteCreatorCount.PositiveReceived = 0
		aggregateVoteCreatorCount.NegativeReceived = 0
	}

	switch msg.Mode {
	case 0:
		if err := k.uncastVote(ctx, msg, &aggregateVoteCreatorCount); err != nil {
			return err
		}
	case 1:
		if err := k.castVote(ctx, msg, &aggregateVoteCreatorCount); err != nil {
			return err
		}
	default:
		return sdkerrors.Wrapf(types.ErrInvalidVotingMode, "mode - (%d)", msg.Mode)
	}

	return nil
}

func (k Keeper) uncastVote(ctx sdk.Context, msg *types.MsgCreateVote, aggregateVoteCreatorCount *types.AggregateVoteCount) error {
	voteBookIndex := types.VoteBookIndex(msg.Creator, msg.Receiver)
	voteBookEntry, found := k.GetVoteBook(ctx, voteBookIndex)
	if !found {
		return sdkerrors.Wrap(types.ErrNoVoteRecord, msg.Receiver)
	}

	aggregateVoteReceiverCount, found := k.GetAggregateVoteCount(ctx, msg.Receiver)
	if !found {
		return sdkerrors.Wrap(types.ErrNoVotesCasted, msg.Receiver)
	}

	ballotBefore := calculateBallot(&aggregateVoteReceiverCount)

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

	if msg.Receiver == msg.Creator {
		k.ReconcileCreatorAggregatedVotes(msg, &aggregateVoteReceiverCount)
		k.ReconcileReceiverAggregatedVotes(msg, &aggregateVoteReceiverCount)
		k.SetAggregateVoteCount(ctx, aggregateVoteReceiverCount)
	} else {
		k.ReconcileCreatorAggregatedVotes(msg, aggregateVoteCreatorCount)
		k.SetAggregateVoteCount(ctx, *aggregateVoteCreatorCount)
		k.ReconcileReceiverAggregatedVotes(msg, &aggregateVoteReceiverCount)
		k.SetAggregateVoteCount(ctx, aggregateVoteReceiverCount)
	}

	ballotAfter := calculateBallot(&aggregateVoteReceiverCount)
	err := k.undelegateStakeAndUnlockMand(ctx, msg, ballotBefore, ballotAfter)
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

func (k Keeper) castVote(ctx sdk.Context, msg *types.MsgCreateVote, aggregateVoteCreatorCount *types.AggregateVoteCount) error {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	mandBalance := k.bankKeeper.GetBalance(ctx, creator, "mand").Amount.Uint64()

	if mandBalance < intAbs(msg.Count) {
		return sdkerrors.Wrapf(types.ErrNotEnoughMand, "count - (%d)", msg.Count)
	}

	voteBookIndex := types.VoteBookIndex(msg.Creator, msg.Receiver)
	voteBookEntry, found := k.GetVoteBook(ctx, voteBookIndex)
	if !found {
		voteBookEntry.Index = voteBookIndex
		voteBookEntry.Positive = 0
		voteBookEntry.Negative = 0
	}

	voteCount := intAbs(msg.Count)
	if msg.Count < 0 {
		voteBookEntry.Negative += voteCount
	} else {
		voteBookEntry.Positive += voteCount
	}

	aggregateVoteReceiverCount, found := k.GetAggregateVoteCount(ctx, msg.Receiver)
	if !found {
		aggregateVoteReceiverCount.Index = msg.Receiver
		aggregateVoteReceiverCount.Casted = 0
		aggregateVoteReceiverCount.PositiveReceived = 0
		aggregateVoteReceiverCount.NegativeReceived = 0
	}

	ballotBefore := calculateBallot(&aggregateVoteReceiverCount)

	k.SetVoteBook(ctx, voteBookEntry)

	if msg.Receiver == msg.Creator {
		k.ReconcileCreatorAggregatedVotes(msg, &aggregateVoteReceiverCount)
		k.ReconcileReceiverAggregatedVotes(msg, &aggregateVoteReceiverCount)
		k.SetAggregateVoteCount(ctx, aggregateVoteReceiverCount)
	} else {
		k.ReconcileCreatorAggregatedVotes(msg, aggregateVoteCreatorCount)
		k.SetAggregateVoteCount(ctx, *aggregateVoteCreatorCount)
		k.ReconcileReceiverAggregatedVotes(msg, &aggregateVoteReceiverCount)
		k.SetAggregateVoteCount(ctx, aggregateVoteReceiverCount)
	}

	ballotAfter := calculateBallot(&aggregateVoteReceiverCount)
	err := k.lockMandAndDelegateStake(ctx, msg, ballotBefore, ballotAfter)
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
