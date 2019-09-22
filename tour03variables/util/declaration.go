package util

import "math/cmplx"

var ( // variables factored into a block
	ToBe   bool       = false
	maxInt uint64     = 1<<64 - 1
	Z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func GetMaxInt() uint64 {
	return maxInt
}
