package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgNetworkConstantData = "network_constant_data"

var (
	_ sdk.Msg = &MsgNetworkConstantData{}

	// NetworkConstantResultStoreKeyPrefix is a prefix for storing result
	NetworkConstantResultStoreKeyPrefix = "network_constant_result"

	// LastNetworkConstantIDKey is the key for the last request id
	LastNetworkConstantIDKey = "network_constant_last_id"

	// NetworkConstantClientIDKey is query request identifier
	NetworkConstantClientIDKey = "network_constant_id"
)

// NewMsgNetworkConstantData creates a new NetworkConstant message
func NewMsgNetworkConstantData(
	creator string,
	oracleScriptID OracleScriptID,
	sourceChannel string,
	calldata *NetworkConstantCallData,
	askCount uint64,
	minCount uint64,
	feeLimit sdk.Coins,
	prepareGas uint64,
	executeGas uint64,
) *MsgNetworkConstantData {
	return &MsgNetworkConstantData{
		ClientID:       NetworkConstantClientIDKey,
		Creator:        creator,
		OracleScriptID: uint64(oracleScriptID),
		SourceChannel:  sourceChannel,
		Calldata:       calldata,
		AskCount:       askCount,
		MinCount:       minCount,
		FeeLimit:       feeLimit,
		PrepareGas:     prepareGas,
		ExecuteGas:     executeGas,
	}
}

// Route returns the message route
func (m *MsgNetworkConstantData) Route() string {
	return RouterKey
}

// Type returns the message type
func (m *MsgNetworkConstantData) Type() string {
	return TypeMsgNetworkConstantData
}

// GetSigners returns the message signers
func (m *MsgNetworkConstantData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes returns the signed bytes from the message
func (m *MsgNetworkConstantData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic check the basic message validation
func (m *MsgNetworkConstantData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if m.SourceChannel == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid source channel")
	}
	return nil
}

// NetworkConstantResultStoreKey is a function to generate key for each result in store
func NetworkConstantResultStoreKey(requestID OracleRequestID) []byte {
	return append(KeyPrefix(NetworkConstantResultStoreKeyPrefix), int64ToBytes(int64(requestID))...)
}
