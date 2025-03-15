package visualization

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func UpdateGameBoard(board [][]int, x int, y int, color int) [][]int {
	if board[y][x] != 0 {
		return nil
	}

	// Create a deep copy of the board
	newBoard := make([][]int, len(board))
	for i := range board {
		newBoard[i] = make([]int, len(board[i]))
		copy(newBoard[i], board[i])
	}

	// Update the copy
	newBoard[y][x] = color
	return newBoard
}

// ReadInput reads user input from the console and returns it as a string
func ReadInput() (int, int) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return -1, -1
	}
	// Trim newline and carriage return characters
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	// split input by space and convert 1st and second elemnt to int
	xVal, yVal := strings.Split(input, " ")[0], strings.Split(input, " ")[1]
	xInt, _ := strconv.Atoi(xVal)
	yInt, _ := strconv.Atoi(yVal)
	return xInt, yInt
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

func CleanBoard(board [][]int) [][]int {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j] = 0
		}
	}
	return board
}
