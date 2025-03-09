package config

import (
	"os"
	"strconv"
)

type GameConfig struct {
	ApiURL    string
	StudentID string
	BoardSize int
}

func SetUpGameConfig() (*GameConfig, error) {

	game := GameConfig{}

	game.ApiURL = os.Getenv("API_URL")
	game.StudentID = os.Getenv("STUDENT_ID")

	var err error
	game.BoardSize, err = strconv.Atoi(os.Getenv("BOARD_SIZE"))
	if err != nil {
		return nil, err
	}

	return &game, nil

}
