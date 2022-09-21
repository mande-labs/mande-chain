package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"mande-chain/testutil/sample"
)

func TestMsgCreateVote_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateVote
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateVote{
				Creator: "invalid_address",
				Mode:    1,
				Count:   1,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateVote{
				Creator: sample.AccAddress(),
				Mode:    1,
				Count:   1,
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
