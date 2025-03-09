package algorithm

import (
	"math"
	"strconv"
)

type candidateMove struct {
	row, col    int
	weightScore float32
}

var visitedCandidates map[candidateMove]bool = make(map[candidateMove]bool)

func CleanVisitedCandidates() {
	visitedCandidates = make(map[candidateMove]bool)
}

func dbSearch(board [][]int, color int) (string, string) {
	blackList := [][]Pattern{threatPatternsBlack, openThreeThreatsBlack, nonThreatPatternsBlack}
	whiteList := [][]Pattern{threatPatternsWhite, openThreeThreatsWhite, nonThreatPatternsWhite}

	// Get opponent color
	opponent := 3 - color // If color is 1, opponent is 2 and vice versa

	// Set search depth (adjust as needed for performance)
	depth := 3

	// Get pattern lists based on player color
	var playerPatterns, opponentPatterns [][]Pattern
	if color == 1 { // Black
		playerPatterns = blackList
		opponentPatterns = whiteList
	} else { // White
		playerPatterns = whiteList
		opponentPatterns = blackList
	}

	// Generate candidate moves based on patterns
	candidates := generateCandidateMoves(board, color, opponent, playerPatterns, opponentPatterns)

	// If no candidates, return center position or first empty position
	if len(candidates) == 0 {

		// Otherwise find first empty position
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[0]); j++ {
				if board[i][j] == 0 {
					return strconv.Itoa(i), strconv.Itoa(j)
				}
			}
		}
	}

	// // Evaluate each candidate move with minimax
	var bestScore float32 = math.MinInt32
	bestRow, bestCol := 8, 8 // Default

	for _, move := range candidates {
		row, col := move.row, move.col

		// Make move
		board[row][col] = color

		// Evaluate with minimax
		score := minimax(board, depth-1, math.MinInt32, math.MaxInt32, false, color, opponent, playerPatterns, opponentPatterns)

		// Undo move
		board[row][col] = 0

		// Update best move
		if score > bestScore {
			bestScore = score
			bestRow, bestCol = row, col
		}
	}

	return strconv.Itoa(bestCol), strconv.Itoa(bestRow)
}

// Generate candidate moves based on patterns
func generateCandidateMoves(board [][]int, player, opponent int, playerPatterns, opponentPatterns [][]Pattern) map[string]candidateMove {
	moveMap := make(map[string]candidateMove)

	// Check for threatening patterns for player (offensive moves)
	for _, patternList := range playerPatterns {
		for _, pattern := range patternList {
			moves := findPatternOffensive(board, pattern)
			for _, move := range moves {
				key := strconv.Itoa(move[0]) + "," + strconv.Itoa(move[1])
				_, ok := moveMap[key]
				if !ok && board[move[0]][move[1]] == 0 {
					moveMap[key] = candidateMove{row: move[0], col: move[1], weightScore: float32(pattern.weight)}
				} else if ok && board[move[0]][move[1]] == 0 {
					tempMove := moveMap[key]
					tempMove.weightScore += float32(pattern.weight)
					moveMap[key] = tempMove
				}
			}
		}
	}

	// Check for opponent's threatening patterns (defensive moves)
	for _, patternList := range opponentPatterns {
		for _, pattern := range patternList {
			moves := findPatternDefensive(board, pattern)

			for _, move := range moves {
				key := strconv.Itoa(move[0]) + "," + strconv.Itoa(move[1])
				_, ok := moveMap[key]
				if !ok && board[move[0]][move[1]] == 0 {
					moveMap[key] = candidateMove{row: move[0], col: move[1], weightScore: float32(pattern.weight)}
				} else if ok && board[move[0]][move[1]] == 0 {
					tempMove := moveMap[key]
					tempMove.weightScore += float32(pattern.weight)
					moveMap[key] = tempMove
				}
			}
		}
	}

	// If no pattern-based moves found, add moves near existing stones
	if len(moveMap) == 0 {
		moveMap = generateNearbyMoves(board, player)
	}

	return moveMap
}

// Generate moves near existing stones
func generateNearbyMoves(board [][]int, player int) map[string]candidateMove {
	moveMap := make(map[string]candidateMove)
	size := len(board)
	visited := make([][]bool, size)
	for i := range visited {
		visited[i] = make([]bool, size)
	}

	// choose 1 at random thats free
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == player {
				// Check adjacent cells
				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						newX, newY := i+x, j+y
						if newX >= 0 && newX < size && newY >= 0 && newY < size && board[newX][newY] == 0 && !visited[newX][newY] {
							key := strconv.Itoa(newX) + "," + strconv.Itoa(newY)
							moveMap[key] = candidateMove{newX, newY, 0}
							visited[newX][newY] = true
						}
					}
				}
			}
		}
	}

	return moveMap
}

// Minimax algorithm with alpha-beta pruning
func minimax(board [][]int, depth int, alpha float32, beta float32, isMaximizing bool,
	player, opponent int, playerPatterns, opponentPatterns [][]Pattern) float32 {
	// Base case: if terminal state or max depth reached
	if depth == 0 || isGameOver(board) {
		return evaluateBoard(board, player, opponent, playerPatterns, opponentPatterns)
	}

	if isMaximizing {
		// Generate moves for player
		moves := generateCandidateMoves(board, player, opponent, playerPatterns, opponentPatterns)
		if len(moves) == 0 {
			return evaluateBoard(board, player, opponent, playerPatterns, opponentPatterns)
		}

		var maxEval float32 = math.MinInt32
		for _, move := range moves {
			row, col := move.row, move.col

			// Make move
			board[row][col] = player

			// Recursive evaluation
			eval := minimax(board, depth-1, alpha, beta, false, player, opponent, playerPatterns, opponentPatterns)

			// Undo move
			board[row][col] = 0

			// Update max evaluation
			maxEval = max(maxEval, eval)

			// Alpha-beta pruning
			alpha = max(alpha, eval)
			if beta <= alpha {
				break
			}
		}
		return maxEval
	} else {
		// Generate moves for opponent
		moves := generateCandidateMoves(board, opponent, player, opponentPatterns, playerPatterns)
		if len(moves) == 0 {
			return evaluateBoard(board, player, opponent, playerPatterns, opponentPatterns)
		}

		var minEval float32 = math.MaxFloat32
		for _, move := range moves {
			row, col := move.row, move.col

			// Make move
			board[row][col] = opponent

			// Recursive evaluation
			eval := minimax(board, depth-1, alpha, beta, true, player, opponent, playerPatterns, opponentPatterns)

			// Undo move
			board[row][col] = 0

			// Update min evaluation
			minEval = min(minEval, eval)

			// Alpha-beta pruning
			beta = min(beta, eval)
			if beta <= alpha {
				break
			}
		}
		return minEval
	}
}

// Evaluate the board state
func evaluateBoard(board [][]int, player, opponent int, playerPatterns, opponentPatterns [][]Pattern) float32 {
	var score float32 = 0

	// Pattern weights based on their type (adjust as needed)
	threatWeight := 100   // Threat patterns
	openThreeWeight := 50 // Open three patterns

	// Check player's patterns
	// Threats
	for _, pattern := range playerPatterns[0] {
		if moves := findPatternOffensive(board, pattern); moves != nil {
			score += float32(pattern.weight * threatWeight)
		}
	}

	// Open threes
	for _, pattern := range playerPatterns[1] {
		if moves := findPatternOffensive(board, pattern); moves != nil {
			score += float32(pattern.weight * openThreeWeight)
		}
	}

	// Check opponent's patterns (defensive evaluation)
	// Threats - higher weight for defense to prioritize blocking
	for _, pattern := range opponentPatterns[0] {
		if moves := findPatternOffensive(board, pattern); moves != nil {
			score -= float32(pattern.weight*threatWeight) * 1.2 // Slightly higher weight for defense
		}
	}

	// Open threes
	for _, pattern := range opponentPatterns[1] {
		if moves := findPatternOffensive(board, pattern); moves != nil {
			score -= float32(pattern.weight * openThreeWeight)
		}
	}

	return score
}

// Check if game is over (win or draw)
func isGameOver(board [][]int) bool {

	if VerifyVictory(board, 1, 10) || VerifyVictory(board, 2, 10) {
		return true
	}

	// Check for draw (full board)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 0 {
				return false
			}
		}
	}

	return true
}

// Helper functions
func max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}
