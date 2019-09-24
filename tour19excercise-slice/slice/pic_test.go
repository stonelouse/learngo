package slice

import (
	"reflect"
	"testing"
)

func TestCratePic(t *testing.T) {
	pic := CreatePic(func(x, y int) uint8 { return uint8(x + y) })
	cases := []struct {
		dx       int
		dy       int
		expected [][]uint8
	}{
		{1, 1, [][]uint8{[]uint8{0}}},
		{2, 3, [][]uint8{
			[]uint8{0, 1, 2},
			[]uint8{1, 2, 3}}},
		{3, 2, [][]uint8{
			[]uint8{0, 1},
			[]uint8{1, 2},
			[]uint8{2, 3}}},
	}
	for index, cse := range cases {
		actual := pic(cse.dx, cse.dy)
		if !reflect.DeepEqual(actual, cse.expected) { // slow! only for test purposes!
			t.Errorf("Actual=%v Expected=%v at index=%d", actual, cse.expected, index)
		}
	}
}
