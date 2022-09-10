package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AggregateVoteCountKeyPrefix is the prefix to retrieve all AggregateVoteCount
	AggregateVoteCountKeyPrefix = "AggregateVoteCount/value/"
)

// AggregateVoteCountKey returns the store key to retrieve a AggregateVoteCount from the index fields
func AggregateVoteCountKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
