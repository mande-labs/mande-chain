package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/voting module sentinel errors
var (
	ErrInvalidVotingMode        = sdkerrors.Register(ModuleName, 1, "invalid voting mode")
	ErrNoVotesCasted            = sdkerrors.Register(ModuleName, 3, "no votes casted")
	ErrNoVoteRecord             = sdkerrors.Register(ModuleName, 4, "voting record not found")
	ErrPositiveVotesCastedLimit = sdkerrors.Register(ModuleName, 5, "you have not casted this many positive votes")
	ErrNegativeVotesCastedLimit = sdkerrors.Register(ModuleName, 6, "you have not casted this many negative votes")
	ErrNotEnoughMand            = sdkerrors.Register(ModuleName, 7, "not enough MAND to use for voting")
	ErrZeroVoteCount            = sdkerrors.Register(ModuleName, 8, "invalid vote count, cannot be zero")
	ErrReceiverIsNotAValidator  = sdkerrors.Register(ModuleName, 9, "validator not found, receiver is not a validator")
)
