package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func CreateGame() (*GameState, error) {
	response, err := http.Post(
		"http://localhost:3000/games",
		"application/json",
		bytes.NewBuffer([]byte("")))

	return responseToGameState(response, err)
}

func JoinGame(gameID int) (*GameState, error) {
	response, err := http.Post(
		fmt.Sprintf("http://localhost:3000/games/%d", gameID),
		"application/json",
		bytes.NewBuffer([]byte("")))
	return responseToGameState(response, err)
}
