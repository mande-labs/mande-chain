package keeper

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
