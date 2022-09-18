package keeper

import (
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

	// send tokens from the vote creator to the voting module account
	voteCountEquivalentMand := sdk.Coins{sdk.NewInt64Coin("mand", int64(msg.Count))}
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, voteCountEquivalentMand)
	if err != nil {
		return err
	}

	// mint stake token to voting module account
	voteCountEquivalentStake := sdk.Coins{sdk.NewInt64Coin("stake", int64(msg.Count))}
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, voteCountEquivalentStake)
	if err != nil {
		return err
	}

	// delegate minted stake tokens to receiver
	votingModuleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))

	valAddr, err := sdk.ValAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrReceiverIsNotAValidator, msg.Receiver)
	}

	recipientValidator, found := k.stakingKeeper.GetValidator(ctx, valAddr)
	if !found {
		return sdkerrors.Wrapf(types.ErrNoValidatorFound, valAddr.String())
	}

	_, err = k.stakingKeeper.Delegate(ctx, votingModuleAcct, sdk.NewInt(int64(msg.Count)), stakingtypes.Unbonded, recipientValidator, true)
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

	// undelegate the stake tokens
	votingModuleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))

	validatorAddr, err := sdk.ValAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrReceiverIsNotAValidator, msg.Receiver)
	}

	_, err = k.stakingKeeper.Undelegate(ctx, votingModuleAcct, validatorAddr, sdk.NewDec(int64(msg.Count)))
	if err != nil {
		return err
	}

	// burn the stake tokens from voting module
	voteCountEquivalentStake := sdk.Coins{sdk.NewInt64Coin("stake", int64(msg.Count))}
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, voteCountEquivalentStake)
	if err != nil {
		return err
	}

	// send mand tokens back to creator
	voteCountEquivalentMand := sdk.Coins{sdk.NewInt64Coin("mand", int64(msg.Count))}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, voteCountEquivalentMand)
	if err != nil {
		return err
	}

	return nil
}
