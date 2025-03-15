//go:build headless
// +build headless

package visualization

// Stub implementations for headless mode
func StartGameLoop()                                         {}
func VisualizeGame(board [][]int)                            {}
func WaitForUserMove() (int, int)                            { return 0, 0 }
func UpdateGameBoard(board [][]int, x, y, color int) [][]int { return board }
func CleanBoard(board [][]int) [][]int                       { return board }
func PrintGameBoard(board [][]int)                           {}

// Add stubs for any other visualization functions
