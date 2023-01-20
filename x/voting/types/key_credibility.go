package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CredibilityKeyPrefix is the prefix to retrieve all Credibility
	CredibilityKeyPrefix = "Credibility/value/"
)

// CredibilityKey returns the store key to retrieve a Credibility from the index fields
func CredibilityKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
