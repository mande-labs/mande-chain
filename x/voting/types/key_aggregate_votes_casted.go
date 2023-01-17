package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AggregateVotesCastedKeyPrefix is the prefix to retrieve all AggregateVotesCasted
	AggregateVotesCastedKeyPrefix = "AggregateVotesCasted/value/"
)

// AggregateVotesCastedKey returns the store key to retrieve a AggregateVotesCasted from the index fields
func AggregateVotesCastedKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
