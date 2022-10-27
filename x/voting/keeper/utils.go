package keeper

import "mande-chain/x/voting/types"

func intAbs(n int64) uint64 {
	y := n >> 63
	return uint64((n ^ y) - y)
}

func Max(x int64, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func calculateBallot(aggregateVoteCreatorCount *types.AggregateVoteCount) int64 {
	return int64(aggregateVoteCreatorCount.PositiveReceived - aggregateVoteCreatorCount.NegativeReceived)
}
