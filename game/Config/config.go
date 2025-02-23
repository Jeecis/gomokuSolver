package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type GameConfig struct {
	ApiURL    string
	StudentID string
	BoardSize int
}

func SetUpGameConfig() (*GameConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	game := GameConfig{}

	game.ApiURL = os.Getenv("API_URL")
	game.StudentID = os.Getenv("STUDENT_ID")

	game.BoardSize, err = strconv.Atoi(os.Getenv("BOARD_SIZE"))
	if err != nil {
		return nil, err
	}

	return &game, nil

}
