package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AggregateVotesReceivedKeyPrefix is the prefix to retrieve all AggregateVotesReceived
	AggregateVotesReceivedKeyPrefix = "AggregateVotesReceived/value/"
)

// AggregateVotesReceivedKey returns the store key to retrieve a AggregateVotesReceived from the index fields
func AggregateVotesReceivedKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
