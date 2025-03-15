package api

import (
	"encoding/json"
	"errors"
	config "gomoku_solver/game/Config"
	"io"
	"log"
	"net/http"
	"time"
)

type GameStartResponse struct {
	GameID        int     `json:"game_id"`
	Status        string  `json:"game_status"`
	Color         string  `json:"color"`
	GameBoard     [][]int `json:"gameboard"`
	RequestStatus string  `json:"request_status"`
	TimeRemaining float32 `json:"time_remaining"`
	Turn          string  `json:"turn"`

	// Add other fields as per your API response
}

type MoveResponse struct {
	GameID        int     `json:"game_id"`
	Status        string  `json:"game_status"`
	Color         string  `json:"color"`
	GameBoard     [][]int `json:"gameboard"`
	RequestStatus string  `json:"request_status"`
	TimeRemaining float32 `json:"time_remaining"`
	Turn          string  `json:"turn"`

	// Add other fields as per your API response
}

func StartGame(cfg *config.GameConfig) (*GameStartResponse, error) {
	req, err := http.NewRequest("GET", cfg.ApiURL+"/"+cfg.StudentID+"/start", nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var gameResponse GameStartResponse
	err = json.Unmarshal(body, &gameResponse)
	if err != nil {
		return nil, err
	}

	if gameResponse.RequestStatus != "GOOD" {
		return nil, errors.New("Bad request status " + gameResponse.RequestStatus)
	}

	return &gameResponse, nil
}

func MoveRequest(cfg *config.GameConfig, gameID string, x string, y string) (*MoveResponse, error) {
	time.Sleep(100 * time.Millisecond)

	req, err := http.NewRequest("GET", cfg.ApiURL+"/"+cfg.StudentID+"/"+gameID+"/"+x+"/"+y, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var moveResponse MoveResponse
	err = json.Unmarshal(body, &moveResponse)
	if err != nil {
		return nil, err
	}

	if moveResponse.RequestStatus != "GOOD" {
		return nil, errors.New("Bad request status " + moveResponse.RequestStatus)
	}

	log.Print(string(body))

	return &moveResponse, nil

}
