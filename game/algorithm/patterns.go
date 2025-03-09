package algorithm

import (
	"slices"
)

type Pattern struct {
	pattern          []int
	placementOptions []int
	weight           int
	counterOptions   []int
}

var winningPatternsBlack = []Pattern{
	// Fives
	{pattern: []int{0, 1, 1, 1, 1}, placementOptions: []int{0}, weight: 10, counterOptions: []int{0}},
	{pattern: []int{1, 0, 1, 1, 1}, placementOptions: []int{1}, weight: 10, counterOptions: []int{1}},
	{pattern: []int{1, 1, 0, 1, 1}, placementOptions: []int{2}, weight: 10, counterOptions: []int{2}},
	{pattern: []int{1, 1, 1, 0, 1}, placementOptions: []int{3}, weight: 10, counterOptions: []int{3}},
	{pattern: []int{1, 1, 1, 1, 0}, placementOptions: []int{4}, weight: 10, counterOptions: []int{4}},
	// Fours
	{pattern: []int{0, 1, 1, 0, 1, 0}, placementOptions: []int{3}, weight: 10, counterOptions: []int{0, 3, 5}},
	{pattern: []int{0, 1, 0, 1, 1, 0}, placementOptions: []int{2}, weight: 10, counterOptions: []int{0, 2, 5}},
	{pattern: []int{0, 1, 1, 1, 0, 0}, placementOptions: []int{4}, weight: 10, counterOptions: []int{0, 4, 5}},
	{pattern: []int{0, 0, 1, 1, 1, 0}, placementOptions: []int{1}, weight: 10, counterOptions: []int{0, 1, 5}},
	{pattern: []int{2, 1, 1, 0, 1, 0, 0, 1, 1, 2}, placementOptions: []int{5}, weight: 10, counterOptions: []int{3, 5, 6}},
	{pattern: []int{2, 1, 1, 0, 0, 1, 0, 1, 1, 2}, placementOptions: []int{4}, weight: 10, counterOptions: []int{3, 4, 6}},
	{pattern: []int{2, 1, 1, 1, 0, 0, 0, 1, 1, 1, 2}, placementOptions: []int{5}, weight: 10, counterOptions: []int{4, 5, 6}},
}

var winningPatternsWhite = []Pattern{
	// Fives
	{pattern: []int{0, 2, 2, 2, 2}, placementOptions: []int{0}, weight: 10, counterOptions: []int{0}},
	{pattern: []int{2, 0, 2, 2, 2}, placementOptions: []int{1}, weight: 10, counterOptions: []int{1}},
	{pattern: []int{2, 2, 0, 2, 2}, placementOptions: []int{2}, weight: 10, counterOptions: []int{2}},
	{pattern: []int{2, 2, 2, 0, 2}, placementOptions: []int{3}, weight: 10, counterOptions: []int{3}},
	{pattern: []int{2, 2, 2, 2, 0}, placementOptions: []int{4}, weight: 10, counterOptions: []int{4}},
	// Fours
	{pattern: []int{0, 2, 2, 0, 2, 0}, placementOptions: []int{3}, weight: 10, counterOptions: []int{0, 3, 5}},
	{pattern: []int{0, 2, 0, 2, 2, 0}, placementOptions: []int{2}, weight: 10, counterOptions: []int{0, 2, 5}},
	{pattern: []int{0, 2, 2, 2, 0, 0}, placementOptions: []int{4}, weight: 10, counterOptions: []int{0, 4, 5}},
	{pattern: []int{0, 0, 2, 2, 2, 0}, placementOptions: []int{1}, weight: 10, counterOptions: []int{0, 1, 5}},
	{pattern: []int{1, 2, 2, 0, 2, 0, 0, 2, 2, 1}, placementOptions: []int{5}, weight: 10, counterOptions: []int{3, 5, 6}},
	{pattern: []int{1, 2, 2, 0, 0, 2, 0, 2, 2, 1}, placementOptions: []int{4}, weight: 10, counterOptions: []int{3, 4, 6}},
	{pattern: []int{1, 2, 2, 2, 0, 0, 0, 2, 2, 2, 1}, placementOptions: []int{5}, weight: 10, counterOptions: []int{4, 5, 6}},
}

var threatPatternsBlack = []Pattern{
	{pattern: []int{2, 0, 1, 1, 1, 0}, placementOptions: []int{1, 5}, weight: 7, counterOptions: []int{5, 1}},
	{pattern: []int{2, 1, 0, 1, 1, 0}, placementOptions: []int{2, 5}, weight: 7, counterOptions: []int{5, 2}},
	{pattern: []int{2, 1, 1, 0, 1, 0}, placementOptions: []int{3, 5}, weight: 7, counterOptions: []int{5, 3}},
	{pattern: []int{2, 1, 1, 1, 0, 0}, placementOptions: []int{4, 5}, weight: 7, counterOptions: []int{5, 4}},

	{pattern: []int{0, 1, 1, 1, 0, 2}, placementOptions: []int{4, 5}, weight: 7, counterOptions: []int{0, 4}},
	{pattern: []int{0, 1, 1, 0, 1, 2}, placementOptions: []int{3, 5}, weight: 7, counterOptions: []int{0, 3}},
	{pattern: []int{0, 1, 0, 1, 1, 2}, placementOptions: []int{2, 5}, weight: 7, counterOptions: []int{0, 2}},
	{pattern: []int{0, 0, 1, 1, 1, 2}, placementOptions: []int{1, 5}, weight: 7, counterOptions: []int{0, 1}},
}

var threatPatternsWhite = []Pattern{
	{pattern: []int{1, 0, 2, 2, 2, 0}, placementOptions: []int{1, 5}, weight: 7, counterOptions: []int{5, 1}},
	{pattern: []int{1, 2, 0, 2, 2, 0}, placementOptions: []int{2, 5}, weight: 7, counterOptions: []int{5, 2}},
	{pattern: []int{1, 2, 2, 0, 2, 0}, placementOptions: []int{3, 5}, weight: 7, counterOptions: []int{5, 3}},
	{pattern: []int{1, 2, 2, 2, 0, 0}, placementOptions: []int{4, 5}, weight: 7, counterOptions: []int{5, 4}},

	{pattern: []int{0, 2, 2, 2, 0, 1}, placementOptions: []int{4, 5}, weight: 7, counterOptions: []int{0, 4}},
	{pattern: []int{0, 2, 2, 0, 2, 1}, placementOptions: []int{3, 5}, weight: 7, counterOptions: []int{0, 3}},
	{pattern: []int{0, 2, 0, 2, 2, 1}, placementOptions: []int{2, 5}, weight: 7, counterOptions: []int{0, 2}},
	{pattern: []int{0, 0, 2, 2, 2, 1}, placementOptions: []int{1, 5}, weight: 7, counterOptions: []int{0, 1}},
}

var openThreeThreatsBlack = []Pattern{
	{pattern: []int{0, 1, 1, 0, 0}, placementOptions: []int{3}, weight: 5, counterOptions: []int{0, 3}},
	{pattern: []int{0, 1, 0, 1, 0}, placementOptions: []int{2}, weight: 5, counterOptions: []int{2}},
	{pattern: []int{0, 0, 1, 1, 0}, placementOptions: []int{1}, weight: 5, counterOptions: []int{1, 4}},
}

var openThreeThreatsWhite = []Pattern{
	{pattern: []int{0, 2, 2, 0, 0}, placementOptions: []int{3}, weight: 5, counterOptions: []int{0, 3}},
	{pattern: []int{0, 2, 0, 2, 0}, placementOptions: []int{2}, weight: 5, counterOptions: []int{2}},
	{pattern: []int{0, 0, 2, 2, 0}, placementOptions: []int{1}, weight: 5, counterOptions: []int{1, 4}},
}

var nonThreatPatternsBlack = []Pattern{
	{pattern: []int{0, 1, 0, 0, 0}, placementOptions: []int{2, 3}, weight: 1, counterOptions: []int{0, 2, 3}},
	{pattern: []int{0, 0, 1, 0, 0}, placementOptions: []int{1, 3}, weight: 1, counterOptions: []int{0, 1, 3}},
	{pattern: []int{0, 0, 0, 1, 0}, placementOptions: []int{3, 2}, weight: 1, counterOptions: []int{0, 2, 3}},
	{pattern: []int{2, 1, 1, 0, 0, 0}, placementOptions: []int{3, 4, 5}, weight: 3, counterOptions: []int{3, 4, 5}},
	{pattern: []int{0, 0, 0, 1, 1, 2}, placementOptions: []int{2, 1, 0}, weight: 3, counterOptions: []int{0, 1, 2}},
}

var nonThreatPatternsWhite = []Pattern{
	{pattern: []int{0, 2, 0, 0, 0}, placementOptions: []int{2, 3}, weight: 1, counterOptions: []int{0, 2, 3}},
	{pattern: []int{0, 0, 2, 0, 0}, placementOptions: []int{1, 3}, weight: 1, counterOptions: []int{0, 1, 3}},
	{pattern: []int{0, 0, 0, 2, 0}, placementOptions: []int{3, 2}, weight: 1, counterOptions: []int{0, 2, 3}},
	{pattern: []int{1, 2, 2, 0, 0, 0}, placementOptions: []int{3, 4, 5}, weight: 3, counterOptions: []int{3, 4, 5}},
	{pattern: []int{0, 0, 0, 2, 2, 1}, placementOptions: []int{2, 1, 0}, weight: 3, counterOptions: []int{0, 1, 2}},
}

func findPatternOffensive(grid [][]int, gamePattern Pattern) [][]int {
	n := len(gamePattern.pattern)
	if n >= 16 || n == 0 {
		return nil
	}

	// Check horizontal direction
	for row := 0; row < 16; row++ {
		for col := 0; col < 17-n; col++ {
			window := make([]int, n)
			for i := 0; i < n; i++ {
				window[i] = grid[row][col+i]
			}
			if slices.Equal(window, gamePattern.pattern) {
				possibleMoves := make([][]int, 0)
				for _, option := range gamePattern.placementOptions {
					possibleMoves = append(possibleMoves, []int{row, col + option})
				}
				return possibleMoves
			}
		}
	}

	// Check vertical direction
	for col := 0; col < 16; col++ {
		for row := 0; row < 17-n; row++ {
			window := make([]int, n)
			for i := 0; i < n; i++ {
				window[i] = grid[row+i][col]
			}
			if slices.Equal(window, gamePattern.pattern) {
				possibleMoves := make([][]int, 0)
				for _, option := range gamePattern.placementOptions {
					possibleMoves = append(possibleMoves, []int{row + option, col})
				}
				return possibleMoves
			}
		}
	}

	// Check diagonal (top-left to bottom-right)
	for row := 0; row < 17-n; row++ {
		for col := 0; col < 17-n; col++ {
			window := make([]int, n)
			for i := 0; i < n; i++ {
				window[i] = grid[row+i][col+i]
			}
			if slices.Equal(window, gamePattern.pattern) {
				possibleMoves := make([][]int, 0)
				for _, option := range gamePattern.placementOptions {
					possibleMoves = append(possibleMoves, []int{row + option, col + option})
				}
				return possibleMoves
			}
		}
	}

	// Check diagonal (top-right to bottom-left)
	for row := 0; row < 17-n; row++ {
		for col := n - 1; col < 16; col++ {
			window := make([]int, n)
			for i := 0; i < n; i++ {
				window[i] = grid[row+i][col-i]
			}
			if slices.Equal(window, gamePattern.pattern) {
				possibleMoves := make([][]int, 0)
				for _, option := range gamePattern.placementOptions {
					possibleMoves = append(possibleMoves, []int{row + option, col - option})
				}
				return possibleMoves
			}
		}
	}

	return nil
}

func findPatternDefensive(grid [][]int, gamePattern Pattern) [][]int {
	n := len(gamePattern.pattern)
	if n == 0 || n > 16 {
		return nil
	}

	// Check horizontal direction
	for row := 0; row < 16; row++ {
		for col := 0; col < 17-n; col++ {
			window := make([]int, n)
			for i := 0; i < n; i++ {
				window[i] = grid[row][col+i]
			}
			if slices.Equal(window, gamePattern.pattern) {
				possibleMoves := make([][]int, len(gamePattern.counterOptions))
				for i, option := range gamePattern.counterOptions {
					possibleMoves[i] = []int{row, col + option}
				}
				return possibleMoves
			}
		}
	}

	// Check vertical direction
	for col := 0; col < 16; col++ {
		for row := 0; row < 17-n; row++ {
			window := make([]int, n)
			for i := 0; i < n; i++ {
				window[i] = grid[row+i][col]
			}
			if slices.Equal(window, gamePattern.pattern) {
				possibleMoves := make([][]int, len(gamePattern.counterOptions))
				for i, option := range gamePattern.counterOptions {
					possibleMoves[i] = []int{row + option, col}
				}
				return possibleMoves
			}
		}
	}

	// Check diagonal (top-left to bottom-right)
	for row := 0; row < 17-n; row++ {
		for col := 0; col < 17-n; col++ {
			window := make([]int, n)
			for i := 0; i < n; i++ {
				window[i] = grid[row+i][col+i]
			}
			if slices.Equal(window, gamePattern.pattern) {
				possibleMoves := make([][]int, len(gamePattern.counterOptions))
				for i, option := range gamePattern.counterOptions {
					possibleMoves[i] = []int{row + option, col + option}
				}
				return possibleMoves
			}
		}
	}

	// Check diagonal (top-right to bottom-left)
	for row := 0; row < 17-n; row++ {
		for col := n - 1; col < 16; col++ {
			window := make([]int, n)
			for i := 0; i < n; i++ {
				window[i] = grid[row+i][col-i]
			}
			if slices.Equal(window, gamePattern.pattern) {
				possibleMoves := make([][]int, len(gamePattern.counterOptions))
				for i, option := range gamePattern.counterOptions {
					possibleMoves[i] = []int{row + option, col - option}
				}
				return possibleMoves
			}
		}
	}

	return nil
}
