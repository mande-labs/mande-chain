package types

import (
	"encoding/binary"
	"fmt"
)

var _ binary.ByteOrder

const (
	// VoteBookKeyPrefix is the prefix to retrieve all VoteBook
	VoteBookKeyPrefix = "VoteBook/value/"
)

func VoteBookIndex(voterAddress string, receiverAddress string) string {
	return fmt.Sprintf("%s-%s", voterAddress, receiverAddress)
}

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
