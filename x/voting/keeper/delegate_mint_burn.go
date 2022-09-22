package keeper

import (
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/tendermint/tendermint/crypto"
	"mande-chain/x/voting/types"
)

/*
- Transfer `mand` tokens from caster account and lock it in module account
- Mint calculated vote amount of `CRED` tokens to module account and invoke `stakingmodule.Delegate` to delegate to receiver (delegator should be voting module account)
*/
func (k Keeper) lockMandAndDelegateStake(ctx sdk.Context, msg *types.MsgCreateVote, ballotBefore int64, ballotAfter int64) error {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	voteCountAbs := int64(Abs(msg.Count))

	// send tokens from the vote creator to the voting module account
	voteCountEquivalentMand := sdk.Coins{sdk.NewInt64Coin("mand", voteCountAbs)}
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, voteCountEquivalentMand)
	if err != nil {
		return err
	}

	err = k.delegationHandler(ctx, msg, ballotBefore, ballotAfter)
	if err != nil {
		return err
	}

	return nil
}

/*
- Transfer `mand` tokens from module account back to user account
- Invoke `stakingmodule.Undelegate` and burn uncasting vote amount of `CRED` tokens from module account using `BeforeDelegationSharesModified` hook
*/
func (k Keeper) undelegateStakeAndUnlockMand(ctx sdk.Context, msg *types.MsgCreateVote, ballotBefore int64, ballotAfter int64) error {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	voteCountAbs := int64(Abs(msg.Count))

	// send mand tokens back to creator
	voteCountEquivalentMand := sdk.Coins{sdk.NewInt64Coin("mand", voteCountAbs)}
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, voteCountEquivalentMand)
	if err != nil {
		return err
	}

	err = k.delegationHandler(ctx, msg, ballotBefore, ballotAfter)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) delegationHandler(ctx sdk.Context, msg *types.MsgCreateVote, ballotBefore int64, ballotAfter int64) error {
	ballotDiff := ballotAfter - ballotBefore

	if ballotDiff > 0 { // delegate
		ballotDiffAdjusted := ballotAfter - Max(ballotBefore, 0)
		if ballotDiffAdjusted > 0 {
			err := k.delegate(ctx, msg, ballotDiffAdjusted)
			if err != nil {
				return err
			}
		}
		return nil
	}

	// un-delegate
	if ballotBefore > 0 {
		ballotDiffAdjusted := ballotBefore - Max(ballotAfter, 0)
		err := k.undelegate(ctx, msg, ballotDiffAdjusted)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k Keeper) delegate(ctx sdk.Context, msg *types.MsgCreateVote, amount int64) error {
	// mint staking token to voting module account
	coinsToMint := sdk.Coins{sdk.NewInt64Coin("cred", amount)}
	err := k.bankKeeper.MintCoins(ctx, types.ModuleName, coinsToMint)
	if err != nil {
		return err
	}

	// delegate minted staking tokens to receiver
	votingModuleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))

	receiverAccAddress, _ := sdk.AccAddressFromBech32(msg.Receiver)
	receiverValAddr, err := sdk.ValAddressFromHex(hex.EncodeToString(receiverAccAddress.Bytes()))
	if err != nil {
		return err
	}

	recipientValidator, found := k.stakingKeeper.GetValidator(ctx, receiverValAddr)
	if !found {
		return sdkerrors.Wrapf(types.ErrReceiverIsNotAValidator, receiverValAddr.String())
	}

	_, err = k.stakingKeeper.Delegate(ctx, votingModuleAcct, sdk.NewInt(amount), stakingtypes.Unbonded, recipientValidator, true)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) undelegate(ctx sdk.Context, msg *types.MsgCreateVote, amount int64) error {
	// undelegate the staking tokens
	votingModuleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	receiverAccAddress, _ := sdk.AccAddressFromBech32(msg.Receiver)
	receiverValAddr, err := sdk.ValAddressFromHex(hex.EncodeToString(receiverAccAddress.Bytes()))
	if err != nil {
		return err
	}

	recipientValidator, found := k.stakingKeeper.GetValidator(ctx, receiverValAddr)
	if !found {
		return sdkerrors.Wrapf(types.ErrReceiverIsNotAValidator, receiverValAddr.String())
	}

	sharesToUnDelegate, err := recipientValidator.SharesFromTokens(sdk.NewInt(amount))
	if err != nil {
		return err
	}

	_, err = k.stakingKeeper.Undelegate(ctx, votingModuleAcct, receiverValAddr, sharesToUnDelegate)
	if err != nil {
		return err
	}

	return nil
}
