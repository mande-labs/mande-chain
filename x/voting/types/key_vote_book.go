package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// VoteBookKeyPrefix is the prefix to retrieve all VoteBook
	VoteBookKeyPrefix = "VoteBook/value/"
)

// VoteBookKey returns the store key to retrieve a VoteBook from the index fields
func VoteBookKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
