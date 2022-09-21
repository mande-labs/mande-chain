package keeper

import "mande-chain/x/voting/types"

func intAbs(n int32) uint64 {
	y := n >> 31
	return uint64((n ^ y) - y)
}

func Abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
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
