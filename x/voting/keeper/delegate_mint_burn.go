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
func (k Keeper) lockMandAndDelegateStake(ctx sdk.Context, msg *types.MsgCreateVote) error {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	voteCountAbs := int64(Abs(msg.Count))

	// send tokens from the vote creator to the voting module account
	voteCountEquivalentMand := sdk.Coins{sdk.NewInt64Coin("mand", voteCountAbs)}
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, voteCountEquivalentMand)
	if err != nil {
		return err
	}

	// mint stake token to voting module account
	voteCountEquivalentStake := sdk.Coins{sdk.NewInt64Coin("stake", voteCountAbs)}
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, voteCountEquivalentStake)
	if err != nil {
		return err
	}

	// delegate minted stake tokens to receiver
	votingModuleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))

	receiverAccAddress, _ := sdk.AccAddressFromBech32(msg.Receiver)
	receiverValAddr, err := sdk.ValAddressFromHex(hex.EncodeToString(receiverAccAddress.Bytes()))
	if err != nil {
		return sdkerrors.Wrapf(types.ErrReceiverIsNotAValidator, msg.Receiver)
	}

	recipientValidator, found := k.stakingKeeper.GetValidator(ctx, receiverValAddr)
	if !found {
		return sdkerrors.Wrapf(types.ErrNoValidatorFound, receiverValAddr.String())
	}

	_, err = k.stakingKeeper.Delegate(ctx, votingModuleAcct, sdk.NewInt(voteCountAbs), stakingtypes.Unbonded, recipientValidator, true)
	if err != nil {
		return err
	}

	return nil
}

/*
- Transfer `mand` tokens from module account back to user account
- Invoke `stakingmodule.Undelegate` and burn uncasting vote amount of `CRED` tokens from module account using `BeforeDelegationSharesModified` hook
*/
func (k Keeper) undelegateStakeAndUnlockMand(ctx sdk.Context, msg *types.MsgCreateVote) error {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	voteCountAbs := int64(Abs(msg.Count))

	// undelegate the stake tokens
	votingModuleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))

	receiverAccAddress, _ := sdk.AccAddressFromBech32(msg.Receiver)
	receiverValAddr, err := sdk.ValAddressFromHex(hex.EncodeToString(receiverAccAddress.Bytes()))
	if err != nil {
		return sdkerrors.Wrapf(types.ErrReceiverIsNotAValidator, msg.Receiver)
	}

	recipientValidator, found := k.stakingKeeper.GetValidator(ctx, receiverValAddr)
	if !found {
		return sdkerrors.Wrapf(types.ErrNoValidatorFound, receiverValAddr.String())
	}

	sharesToUnDelegate, err := recipientValidator.SharesFromTokens(sdk.NewInt(voteCountAbs))
	if err != nil {
		return err
	}

	_, err = k.stakingKeeper.Undelegate(ctx, votingModuleAcct, receiverValAddr, sharesToUnDelegate)
	if err != nil {
		return err
	}

	// burn the stake tokens from voting module
	//voteCountEquivalentStake := sdk.Coins{sdk.NewInt64Coin("stake", voteCountAbs)}
	//err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, voteCountEquivalentStake)
	//if err != nil {
	//	return err
	//}

	// send mand tokens back to creator
	voteCountEquivalentMand := sdk.Coins{sdk.NewInt64Coin("mand", voteCountAbs)}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, voteCountEquivalentMand)
	if err != nil {
		return err
	}

	return nil
}
