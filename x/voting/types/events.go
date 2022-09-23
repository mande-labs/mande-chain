package types

// bank module event types
const (
	EventTypeVoteCasted = "vote_casted"
	EventTypeVoteUncasted = "vote_uncasted"

	AttributeKeyCaster    = "sender"
	AttributeKeyReceiver = "recipient"
)
