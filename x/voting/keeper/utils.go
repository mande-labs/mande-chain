package keeper

func intAbs(n int32) uint64 {
	y := n >> 31
	return uint64((n ^ y) - y)
}
