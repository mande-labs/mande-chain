package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateCredibility = "update_credibility"

var _ sdk.Msg = &MsgUpdateCredibility{}

func NewMsgUpdateCredibility(creator string, address string) *MsgUpdateCredibility {
	return &MsgUpdateCredibility{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgUpdateCredibility) Route() string {
	return RouterKey
}

func (msg *MsgUpdateCredibility) Type() string {
	return TypeMsgUpdateCredibility
}

func (msg *MsgUpdateCredibility) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateCredibility) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateCredibility) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}
	return nil
}
