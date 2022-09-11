package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"mande-chain/x/voting/types"
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
	k.ReconcileAggregatedVotes(msg, aggregateVoteCreatorCount, &aggregateVoteReceiverCount)
	k.SetAggregateVoteCount(ctx, aggregateVoteReceiverCount)
	k.SetAggregateVoteCount(ctx, *aggregateVoteCreatorCount)

	return nil
}

func (k Keeper) castVote(ctx sdk.Context, msg *types.MsgCreateVote, aggregateVoteCreatorCount *types.AggregateVoteCount) error {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	mandBalance := k.bankKeeper.GetBalance(ctx, creator, "mand").Amount.Uint64()
	voteCastCount := aggregateVoteCreatorCount.Casted + intAbs(msg.Count)

	if mandBalance < voteCastCount {
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

	k.SetVoteBook(ctx, voteBookEntry)

	aggregateVoteReceiverCount, found := k.GetAggregateVoteCount(ctx, msg.Receiver)
	if !found {
		aggregateVoteReceiverCount.Index = msg.Receiver
		aggregateVoteReceiverCount.Casted = 0
		aggregateVoteReceiverCount.PositiveReceived = 0
		aggregateVoteReceiverCount.NegativeReceived = 0
	}

	k.ReconcileAggregatedVotes(msg, aggregateVoteCreatorCount, &aggregateVoteReceiverCount)
	k.SetAggregateVoteCount(ctx, aggregateVoteReceiverCount)
	k.SetAggregateVoteCount(ctx, *aggregateVoteCreatorCount)

	return nil
}
