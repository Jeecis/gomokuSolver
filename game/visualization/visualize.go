package visualization

import (
	"image/color"
	"log"
	"strconv"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 640
	cellSize     = 40
	boardSize    = 16
)

// Game represents the game state
type Game struct {
	board      [][]int
	boardMutex sync.RWMutex
}

// UserMoveChannel is used to communicate user clicks to the game logic
var UserMoveChannel = make(chan [2]int)

// waitingForMove indicates if the game is expecting a move from the user
var waitingForMove bool

func (g *Game) Update() error {
	if waitingForMove && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		// Get the mouse position
		x, y := ebiten.CursorPosition()

		// Convert to board coordinates
		boardX := x / cellSize
		boardY := y / cellSize

		// Check if the click is within board bounds
		if boardX >= 0 && boardX < boardSize && boardY >= 0 && boardY < boardSize {
			// Check if the cell is empty
			g.boardMutex.RLock()
			isEmpty := g.board[boardY][boardX] == 0
			g.boardMutex.RUnlock()

			if isEmpty {
				// Send the move through the channel
				select {
				case UserMoveChannel <- [2]int{boardY, boardX}:
					waitingForMove = false
				default:
					// Channel is not ready, will try on next update
				}
			}
		}
	}
	return nil
}

func WaitForUserMove() (row, col int) {
	waitingForMove = true
	move := <-UserMoveChannel
	return move[0], move[1]
}

// Draw draws the game screen
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw background
	screen.Fill(color.RGBA{238, 220, 180, 255}) // Light wooden color

	// Draw grid lines
	for i := 0; i <= boardSize; i++ {
		// Vertical lines
		vector.StrokeLine(screen,
			float32(i*cellSize), 0,
			float32(i*cellSize), float32(boardSize*cellSize),
			1, color.Black, true)
		// Horizontal lines
		vector.StrokeLine(screen,
			0, float32(i*cellSize),
			float32(boardSize*cellSize), float32(i*cellSize),
			1, color.Black, true)
	}

	// Lock before reading the board
	g.boardMutex.RLock()
	defer g.boardMutex.RUnlock()

	// Check if board is initialized
	if g.board == nil || len(g.board) == 0 {
		ebitenutil.DebugPrintAt(screen, "Initializing board...", 10, 10)
		return
	}

	// Draw pieces
	for i := 0; i < boardSize; i++ {
		if i >= len(g.board) {
			continue
		}
		for j := 0; j < boardSize; j++ {
			if j >= len(g.board[i]) {
				continue
			}
			x, y := float32(j*cellSize+cellSize/2), float32(i*cellSize+cellSize/2)

			if g.board[i][j] == 1 { // Black piece
				vector.DrawFilledCircle(screen, x, y, float32(cellSize/2-4), color.Black, true)
			} else if g.board[i][j] == 2 { // White piece
				vector.DrawFilledCircle(screen, x, y, float32(cellSize/2-4), color.White, true)
				vector.StrokeCircle(screen, x, y, float32(cellSize/2-4), 1, color.Black, true)
			}
		}
	}

	// Draw coordinates for reference
	for i := 0; i < boardSize; i++ {
		// convert i to string
		iStr := strconv.Itoa(i)
		// Draw column numbers on top
		ebitenutil.DebugPrintAt(screen, iStr, i*cellSize+cellSize/3, 5)
		// Draw row numbers on the left
		ebitenutil.DebugPrintAt(screen, iStr, 5, i*cellSize+cellSize/3)
	}
}

// Layout returns the game's layout
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// UpdateBoard updates the game board with a new state
func (g *Game) UpdateBoard(newBoard [][]int) {
	g.boardMutex.Lock()
	defer g.boardMutex.Unlock()
	g.board = newBoard
}

// Global game instance
var gameInstance *Game
var gameStarted bool = false
var initCh = make(chan [][]int)

// StartGameLoop starts the game visualization loop (call this from main)
func StartGameLoop() {
	// Initialize the game instance with an empty board first
	emptyBoard := make([][]int, boardSize)
	for i := range emptyBoard {
		emptyBoard[i] = make([]int, boardSize)
	}

	gameInstance = &Game{board: emptyBoard}

	// Configure the window
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Gomoku Board")

	// Start a goroutine to receive board updates
	go func() {
		for board := range initCh {
			gameInstance.UpdateBoard(board)
		}
	}()

	// This must run on the main thread
	if err := ebiten.RunGame(gameInstance); err != nil {
		log.Fatal(err)
	}
}

// VisualizeGame updates the game board
func VisualizeGame(board [][]int) {
	// Mark that we've started the visualization
	if !gameStarted {
		// Send the initial board state
		select {
		case initCh <- board:
			gameStarted = true
		default:
			// The window hasn't been initialized yet, this is fine
		}
		return
	}

	// Update the board state
	gameInstance.UpdateBoard(board)
}
