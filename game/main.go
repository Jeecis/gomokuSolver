package main

import (
	"errors"
	"fmt"
	api "gomoku_solver/game/API"
	config "gomoku_solver/game/Config"
	"gomoku_solver/game/algorithm"

	"log"
	"time"
)

func main() {
	cfg, err := config.SetUpGameConfig()
	if err != nil {
		panic(err)
	}

	gameMode := cfg.GameMode
	if gameMode == "api" {
		apiLogicNoVisuals(cfg)
		return
		// Uncomment the following code to visualize the game on windows
		// } else if gameMode == "apiVisualize" {
		// 	go apiLogic(cfg)
		// 	visualization.StartGameLoop()
		// } else if gameMode == "pvb" {
		// 	go playerVsAILogicBotStart()
		// 	visualization.StartGameLoop()
		// } else if gameMode == "pvh" {
		// 	go playerVsAILogicHumanStart()
		// 	visualization.StartGameLoop()
	} else {
		panic(errors.New("Specify the game mode with the -gm flag. Modes available: api, pvb, pvh"))
	}
}

func PrintGameBoard(board [][]int) {
	fmt.Println("Printing game board...")
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Print(board[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println("+++++++++++++++++++++++++++++++")
}

func apiLogicNoVisuals(cfg *config.GameConfig) {
	// Small delay to ensure the visualization window is ready
	time.Sleep(500 * time.Millisecond)

	var currentStatus string
	var color int
	var board [][]int
	for {
		move := 1
		// Start the game
		gameStartResponse, err := api.StartGame(cfg)
		if err != nil {
			panic(err)
		}
		if gameStartResponse.Color == "BLACK" {
			color = 1
		} else {
			color = 2
		}
		gameID := fmt.Sprintf("%v", gameStartResponse.GameID)
		board = gameStartResponse.GameBoard

		if gameStartResponse.RequestStatus != "GOOD" || currentStatus == "LEAVE" {
			panic(errors.New("Error starting the game"))
		}

		timeRemaining := gameStartResponse.TimeRemaining

		// calculate move
		for {
			// Make a move
			x, y := algorithm.CalculateMove(board, color, move, timeRemaining)
			moveResponse, err := api.MoveRequest(cfg, gameID, x, y)
			if err != nil {
				// visualization.PrintGameBoard(board)
				PrintGameBoard(board)
				panic(err)
			}
			board = moveResponse.GameBoard
			timeRemaining = moveResponse.TimeRemaining
			move += 1

			currentStatus = moveResponse.Status
			if currentStatus == "BLACKWON" || currentStatus == "WHITEWON" || currentStatus == "DRAW" || currentStatus == "LEAVE" {
				algorithm.CleanVisitedCandidates()
				log.Print("Game Over! Color: ", color, " Status: ", currentStatus)
				break
			}
			move += 1
		}
		if currentStatus == "LEAVE" {
			break
		}
	}
}

// func apiLogic(cfg *config.GameConfig) {
// 	// Small delay to ensure the visualization window is ready
// 	time.Sleep(500 * time.Millisecond)

// 	var currentStatus string
// 	var color int
// 	var board [][]int
// 	for {
// 		move := 1
// 		// Start the game
// 		gameStartResponse, err := api.StartGame(cfg)
// 		if err != nil {
// 			panic(err)
// 		}
// 		if gameStartResponse.Color == "black" {
// 			color = 1
// 		} else {
// 			color = 2
// 		}
// 		gameID := fmt.Sprintf("%v", gameStartResponse.GameID)
// 		board = gameStartResponse.GameBoard

// 		if gameStartResponse.RequestStatus != "GOOD" || currentStatus == "LEAVE" {
// 			panic(errors.New("Error starting the game"))
// 		}

// 		// Important: Visualize the initial board state
// 		visualization.VisualizeGame(board)

// 		timeRemaining := gameStartResponse.TimeRemaining

// 		// calculate move
// 		for {
// 			// Make a move
// 			x, y := algorithm.CalculateMove(board, color, move, timeRemaining)
// 			moveResponse, err := api.MoveRequest(cfg, gameID, x, y)
// 			if err != nil {
// 				visualization.PrintGameBoard(board)
// 				panic(err)
// 			}
// 			board = moveResponse.GameBoard
// 			timeRemaining = moveResponse.TimeRemaining
// 			visualization.VisualizeGame(board)
// 			move += 1

// 			currentStatus = moveResponse.Status
// 			if currentStatus == "BLACKWON" || currentStatus == "WHITEWON" || currentStatus == "DRAW" || currentStatus == "LEAVE" {
// 				algorithm.CleanVisitedCandidates()
// 				log.Print("Game Over! Color: ", color, " Status: ", currentStatus)
// 				break
// 			}
// 			move += 1
// 		}
// 		if currentStatus == "LEAVE" {
// 			break
// 		}
// 	}
// }

// func playerVsAILogicBotStart() {
// 	time.Sleep(500 * time.Millisecond)

// 	var board [][]int
// 	for {
// 		move := 1

// 		board = make([][]int, 16)
// 		for i := range board {
// 			board[i] = make([]int, 16)
// 		}

// 		x, y := algorithm.CalculateMove(board, 1, move, -1)
// 		xInt, _ := strconv.Atoi(x)
// 		yInt, _ := strconv.Atoi(y)
// 		board = visualization.UpdateGameBoard(board, xInt, yInt, 1)
// 		visualization.VisualizeGame(board)
// 		move += 1
// 		for {

// 			yInt, xInt = visualization.WaitForUserMove()
// 			board = visualization.UpdateGameBoard(board, xInt, yInt, 2)
// 			visualization.VisualizeGame(board)
// 			move += 1

// 			res := algorithm.VerifyVictory(board, 2, move)
// 			if res {
// 				fmt.Println("White wins")
// 				_, _ = visualization.WaitForUserMove()
// 				board = visualization.CleanBoard(board)

// 				visualization.VisualizeGame(board)
// 				break
// 			}

// 			x, y = algorithm.CalculateMove(board, 1, move, -1)
// 			xInt, _ = strconv.Atoi(x)
// 			yInt, _ = strconv.Atoi(y)
// 			board = visualization.UpdateGameBoard(board, xInt, yInt, 1)
// 			visualization.VisualizeGame(board)
// 			move += 1

// 			res = algorithm.VerifyVictory(board, 1, move)
// 			if res {
// 				fmt.Println("Black wins")
// 				_, _ = visualization.WaitForUserMove()
// 				board = visualization.CleanBoard(board)
// 				visualization.VisualizeGame(board)
// 				break
// 			}
// 		}

// 	}
// }

// func playerVsAILogicHumanStart() {
// 	time.Sleep(500 * time.Millisecond)

// 	var board [][]int
// 	for {
// 		move := 1

// 		board = make([][]int, 16)
// 		for i := range board {
// 			board[i] = make([]int, 16)
// 		}

// 		yInt, xInt := visualization.WaitForUserMove()
// 		board = visualization.UpdateGameBoard(board, xInt, yInt, 1)
// 		visualization.VisualizeGame(board)
// 		move += 1

// 		for {

// 			x, y := algorithm.CalculateMove(board, 2, move, -1)
// 			xInt, _ = strconv.Atoi(x)
// 			yInt, _ = strconv.Atoi(y)
// 			board = visualization.UpdateGameBoard(board, xInt, yInt, 2)
// 			visualization.VisualizeGame(board)
// 			move += 1

// 			res := algorithm.VerifyVictory(board, 2, move)
// 			if res {
// 				fmt.Println("White wins")
// 				_, _ = visualization.WaitForUserMove()
// 				board = visualization.CleanBoard(board)
// 				visualization.VisualizeGame(board)
// 				break
// 			}

// 			yInt, xInt := visualization.WaitForUserMove()
// 			board = visualization.UpdateGameBoard(board, xInt, yInt, 1)
// 			visualization.VisualizeGame(board)
// 			move += 1

// 			res = algorithm.VerifyVictory(board, 1, move)
// 			if res {
// 				fmt.Println("Black wins")
// 				_, _ = visualization.WaitForUserMove()
// 				board = visualization.CleanBoard(board)
// 				log.Print(board)
// 				visualization.VisualizeGame(board)
// 				break
// 			}
// 		}

// 	}
// }
