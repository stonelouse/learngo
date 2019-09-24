package slice

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
