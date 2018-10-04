package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GameState struct {
	GameID   int    `json:"game-id"`
	PlayerID string `json:"player-id"`
	Status   string `json:"status"`
}

func CreateGame() (*GameState, error) {
	response, err := http.Post(
		"http://localhost:3000/games",
		"application/json",
		bytes.NewBuffer([]byte("")))

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var state GameState
	err = json.Unmarshal(body, &state)
	if err != nil {
		return nil, err
	}

	return &state, nil
}

func JoinGame(gameID int) (*GameState, error) {
	response, err := http.Post(
		fmt.Sprintf("http://localhost:3000/games/%d", gameID),
		"application/json",
		bytes.NewBuffer([]byte("")))

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var state GameState
	err = json.Unmarshal(body, &state)
	if err != nil {
		return nil, err
	}

	return &state, nil
}

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
