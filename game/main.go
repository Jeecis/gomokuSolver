package main

import config "gomoku_solver/game/Config"

func main() {
	_, err := config.SetUpGameConfig()
	if err != nil {
		panic(err)
	}

}
