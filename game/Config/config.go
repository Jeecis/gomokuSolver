package config

import (
	"errors"
	"flag"
	"os"
)

type GameConfig struct {
	ApiURL    string
	StudentID string
	GameMode  string
}

func SetUpGameConfig() (*GameConfig, error) {
	game := GameConfig{}

	apiURL := flag.String("url", "", "API URL")
	studentID := flag.String("id", "", "Student ID")
	gameMode := flag.String("gm", "", "Game Mode")

	flag.Parse()

	game.ApiURL = os.Getenv("API_URL")
	if *apiURL != "" {
		game.ApiURL = *apiURL
	}

	game.StudentID = os.Getenv("STUDENT_ID")
	if *studentID != "" {
		game.StudentID = *studentID
	}

	game.GameMode = os.Getenv("GAME_MODE")
	if *gameMode != "" {
		game.GameMode = *gameMode
	}

	if (game.GameMode == "api" || game.GameMode == "apiVisualize") && (game.ApiURL == "" || game.StudentID == "") {
		return nil, errors.New("API URL and Student ID are required for API mode")
	}

	return &game, nil
}
