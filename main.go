package main

import (
	"fmt"
)

func main() {
	gameState, err := CreateGame()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(gameState)
	}
	gameState2, err := JoinGame(gameState.GameID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(gameState2)
	}
}
