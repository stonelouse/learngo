package main

import (
	"github.com/stonelouse/learngo/tour19excercise-slice/slice"
	"golang.org/x/tour/pic"
)

func main() {
	var picDataAlgorithm func(x, y int) uint8 = func(x, y int) uint8 { return uint8((x + y) / 2) }
	var picDataGenerator func(dx, dy int) [][]uint8 = slice.CreatePic(picDataAlgorithm)
	pic.Show(picDataGenerator)
}

/*
	copy an run in interactive playground 'A Tour of Go'

	package main

	import "golang.org/x/tour/pic"

	func CreatePic(fn func(x, y int) uint8) func(dx, dy int) [][]uint8 {
		return func(dx, dy int) [][]uint8 {
			// TODO validate dx, dy > 0
			result := make([][]uint8, dx)
			for x := 0; x < dx; x++ {
				result[x] = make([]uint8, dy)
				for y := 0; y < dy; y++ {
					result[x][y] = fn(x, y)
				}
			}
			return result
		}
	}
	func main() {
		var picDataAlgorithm func(x, y int) uint8 = func(x, y int) uint8 { return uint8((x + y) / 2) }
		var picDataGenerator func(dx, dy int) [][]uint8 = CreatePic(picDataAlgorithm)
		pic.Show(picDataGenerator)
	}
*/
