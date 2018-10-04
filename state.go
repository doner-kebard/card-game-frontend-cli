package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type GameState struct {
	GameID   int    `json:"game-id"`
	PlayerID string `json:"player-id"`
	Status   string `json:"status"`
}

func responseToGameState(response *http.Response, err error) (*GameState, error) {
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
