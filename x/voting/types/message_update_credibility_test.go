package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"mande-chain/testutil/sample"
)

func TestMsgUpdateCredibility_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateCredibility
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateCredibility{
				Address: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateCredibility{
				Address: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
