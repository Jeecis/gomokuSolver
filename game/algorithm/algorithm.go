package algorithm

import (
	"strconv"
)

func CalculateMove(board [][]int, color int, move int, timeRemaining float32) (string, string) {

	if move < 5 {
		return openingMove(board, move)
	}

	var depth int
	if timeRemaining == -1 || timeRemaining > 200 {
		depth = 4
	} else if timeRemaining > 100 {
		depth = 3
	} else if timeRemaining > 50 {
		depth = 2
	} else {
		depth = 1
	}

	var winningPattern []Pattern
	if color == 1 {
		winningPattern = winningPatternsBlack
	} else {
		winningPattern = winningPatternsWhite
	}

	for _, pattern := range winningPattern {
		winningOptions := findPatternOffensive(board, pattern)

		if winningOptions != nil {
			return strconv.Itoa(winningOptions[0][1]), strconv.Itoa(winningOptions[0][0])
		}
	}

	//search for opponents winning patterns and block them
	var opponentWinningPattern []Pattern
	if color == 2 {
		opponentWinningPattern = winningPatternsBlack
	} else {
		opponentWinningPattern = winningPatternsWhite
	}

	for _, pattern := range opponentWinningPattern {

		winningOptions := findPatternDefensive(board, pattern)

		if winningOptions != nil {
			return strconv.Itoa(winningOptions[0][1]), strconv.Itoa(winningOptions[0][0])
		}
	}

	// No winning patterns found, generate candidate moves
	return dbSearch(board, color, depth)
}

// Make move towards center of the board but next to the opponents piece
func openingMove(board [][]int, move int) (string, string) {
	if move == 1 {
		return "7", "7"
	} else if move == 4 {
		return moveTowardsCenter(board, 2) // we want to put next to white piece
	}
	return moveTowardsCenter(board, 1) // we want to put next to black piece because its a stronger start that way
}

func moveTowardsCenter(board [][]int, color int) (string, string) {
	var pinX, pinY int
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			if board[i][j] == color {
				pinY = i
				pinX = j
				break
			}
		}
	}

	directions := []struct{ dy, dx int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	// Calculate center of the board
	centerY, centerX := 7, 7

	// Find best adjacent position (closest to center)
	bestDist := 100.0
	bestY, bestX := -1, -1

	for _, dir := range directions {
		y := pinY + dir.dy
		x := pinX + dir.dx

		// Check if position is valid and empty
		if y >= 0 && y < 16 && x >= 0 && x < 16 && board[y][x] == 0 {
			// Calculate distance to center
			distToCenter := ((centerY - y) * (centerY - y)) + ((centerX - x) * (centerX - x))

			if float64(distToCenter) < bestDist {
				bestDist = float64(distToCenter)
				bestY = y
				bestX = x
			}
		}
	}

	return strconv.Itoa(bestX), strconv.Itoa(bestY)
}
