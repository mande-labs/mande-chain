package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		VoteBookList:           []VoteBook{},
		AggregateVoteCountList: []AggregateVoteCount{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in voteBook
	voteBookIndexMap := make(map[string]struct{})

	for _, elem := range gs.VoteBookList {
		index := string(VoteBookKey(elem.Index))
		if _, ok := voteBookIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for voteBook")
		}
		voteBookIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in aggregateVoteCount
	aggregateVoteCountIndexMap := make(map[string]struct{})

	for _, elem := range gs.AggregateVoteCountList {
		index := string(AggregateVoteCountKey(elem.Index))
		if _, ok := aggregateVoteCountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for aggregateVoteCount")
		}
		aggregateVoteCountIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
