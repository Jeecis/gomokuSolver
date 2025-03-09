package algorithm

func VerifyVictory(grid [][]int, color int, move int) bool {
	if move < 9 {
		return false
	}

	pattern := []int{color, color, color, color, color}
	res := findPatternOffensive(grid, Pattern{pattern: pattern, placementOptions: []int{0}, weight: 10, counterOptions: []int{0}})
	return res != nil
}
