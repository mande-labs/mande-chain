package keeper

import (
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

func (k Keeper) uncastVote(ctx sdk.Context, msg *types.MsgCreateVote, aggregateVoteCreatorCount *types.AggregateVotesCasted) error {
	voteBookIndex := types.VoteBookIndex(msg.Creator, msg.Receiver)
	voteBookEntry, found := k.GetVoteBook(ctx, voteBookIndex)
	if !found {
		return sdkerrors.Wrap(types.ErrNoVoteRecord, msg.Receiver)
	}

	aggregateVoteReceiverCount, found := k.GetAggregateVotesReceived(ctx, msg.Receiver)
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
	
	k.ReconcileCreatorAggregatedVotes(msg, aggregateVoteCreatorCount)
	k.ReconcileReceiverAggregatedVotes(msg, &aggregateVoteReceiverCount)
	k.SetAggregateVotesCasted(ctx, *aggregateVoteCreatorCount)
	k.SetAggregateVotesReceived(ctx, aggregateVoteReceiverCount)

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

func (k Keeper) castVote(ctx sdk.Context, msg *types.MsgCreateVote, aggregateVotesCreatorCount *types.AggregateVotesCasted) error {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	mandBalance := k.bankKeeper.GetBalance(ctx, creator, "mand").Amount.Uint64()

	if mandBalance < intAbs(msg.Count) {
		return sdkerrors.Wrapf(types.ErrNotEnoughMand, "count - (%d)", msg.Count)
	}

	voteBookIndex := types.VoteBookIndex(msg.Creator, msg.Receiver)
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

	ballotBefore := calculateBallot(&aggregateVotesReceiverCount)

	k.SetVoteBook(ctx, voteBookEntry)

	k.ReconcileCreatorAggregatedVotes(msg, aggregateVotesCreatorCount)
	k.ReconcileReceiverAggregatedVotes(msg, &aggregateVotesReceiverCount)
	k.SetAggregateVotesCasted(ctx, *aggregateVotesCreatorCount)
	k.SetAggregateVotesReceived(ctx, aggregateVotesReceiverCount)

	ballotAfter := calculateBallot(&aggregateVotesReceiverCount)
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
