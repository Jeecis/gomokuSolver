package main

import (
	"errors"
	"fmt"
	api "gomoku_solver/game/API"
	config "gomoku_solver/game/Config"
	"gomoku_solver/game/algorithm"
	"gomoku_solver/game/visualization"
	"log"
	"strconv"
	"time"
)

func main() {
	go apiLogic()
	visualization.StartGameLoop()
}

func apiLogic() {
	// Small delay to ensure the visualization window is ready
	time.Sleep(500 * time.Millisecond)

	cfg, err := config.SetUpGameConfig()
	if err != nil {
		panic(err)
	}

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

		// Important: Visualize the initial board state
		visualization.VisualizeGame(board)

		// calculate move
		for {
			// Make a move
			x, y := algorithm.CalculateMove(board, color, move)
			moveResponse, err := api.MoveRequest(cfg, gameID, x, y)
			if err != nil {
				visualization.PrintGameBoard(board)
				panic(err)
			}
			board = moveResponse.GameBoard
			visualization.VisualizeGame(board)
			move += 1

			currentStatus = moveResponse.Status
			if currentStatus == "BLACKWON" || currentStatus == "WHITEWON" || currentStatus == "DRAW" || currentStatus == "LEAVE" {
				algorithm.CleanVisitedCandidates()
				log.Print("Game Over! Color: ", color, " Status: ", currentStatus)
				break
			}
			move += 1

			x, y = algorithm.CalculateMove(board, 1, move)
		}
		if currentStatus == "LEAVE" {
			break
		}
	}
}

func playerVsAILogicBotStart() {
	time.Sleep(500 * time.Millisecond)

	var board [][]int
	for {
		move := 1

		board = make([][]int, 16)
		for i := range board {
			board[i] = make([]int, 16)
		}

		x, y := algorithm.CalculateMove(board, 1, move)
		xInt, _ := strconv.Atoi(x)
		yInt, _ := strconv.Atoi(y)
		board = visualization.UpdateGameBoard(board, xInt, yInt, 1)
		visualization.VisualizeGame(board)
		move += 1
		for {

			xInt, yInt = visualization.ReadInput()
			board = visualization.UpdateGameBoard(board, xInt, yInt, 2)
			visualization.VisualizeGame(board)
			move += 1

			res := algorithm.VerifyVictory(board, 2, move)
			if res {
				fmt.Println("White wins")
				break
			}

			x, y = algorithm.CalculateMove(board, 1, move)
			xInt, _ = strconv.Atoi(x)
			yInt, _ = strconv.Atoi(y)
			board = visualization.UpdateGameBoard(board, xInt, yInt, 1)
			visualization.VisualizeGame(board)
			move += 1

			res = algorithm.VerifyVictory(board, 1, move)
			if res {
				fmt.Println("Black wins")
				break
			}
		}

	}
}

func playerVsAILogicHumanStart() {
	time.Sleep(500 * time.Millisecond)

	var board [][]int
	for {
		move := 1

		board = make([][]int, 16)
		for i := range board {
			board[i] = make([]int, 16)
		}

		xInt, yInt := visualization.ReadInput()
		board = visualization.UpdateGameBoard(board, xInt, yInt, 1)
		visualization.PrintGameBoard(board)
		move += 1
		log.Print("=====================================")

		for {

			x, y := algorithm.CalculateMove(board, 2, move)
			xInt, _ = strconv.Atoi(x)
			yInt, _ = strconv.Atoi(y)
			board = visualization.UpdateGameBoard(board, xInt, yInt, 2)
			// visualization.VisualizeGame(board)
			visualization.PrintGameBoard(board)
			move += 1
			log.Print("=====================================")

			res := algorithm.VerifyVictory(board, 2, move)
			if res {
				fmt.Println("White wins")
				break
			}

			xInt, yInt = visualization.ReadInput()
			board = visualization.UpdateGameBoard(board, xInt, yInt, 1)
			visualization.PrintGameBoard(board)
			move += 1
			log.Print("=====================================")

			res = algorithm.VerifyVictory(board, 1, move)
			if res {
				fmt.Println("Black wins")
				break
			}
		}

	}
}
