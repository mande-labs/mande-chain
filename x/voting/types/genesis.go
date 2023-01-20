package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		VoteBookList:               []VoteBook{},
		AggregateVotesCastedList:   []AggregateVotesCasted{},
		AggregateVotesReceivedList: []AggregateVotesReceived{},
		CredibilityList:            []Credibility{},
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
	// Check for duplicated index in aggregateVotesCasted
	aggregateVotesCastedIndexMap := make(map[string]struct{})

	for _, elem := range gs.AggregateVotesCastedList {
		index := string(AggregateVotesCastedKey(elem.Index))
		if _, ok := aggregateVotesCastedIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for aggregateVotesCasted")
		}
		aggregateVotesCastedIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in aggregateVotesReceived
	aggregateVotesReceivedIndexMap := make(map[string]struct{})

	for _, elem := range gs.AggregateVotesReceivedList {
		index := string(AggregateVotesReceivedKey(elem.Index))
		if _, ok := aggregateVotesReceivedIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for aggregateVotesReceived")
		}
		aggregateVotesReceivedIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in credibility
	credibilityIndexMap := make(map[string]struct{})

	for _, elem := range gs.CredibilityList {
		index := string(CredibilityKey(elem.Index))
		if _, ok := credibilityIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for credibility")
		}
		credibilityIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
