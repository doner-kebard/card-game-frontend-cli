package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Lobby struct {
	GameID      string    `json:"game-id"`
	PlayerID    string `json:"player-id"`
	LobbyStatus string `json:"lobby-status"`
}

func responseToLobby(response *http.Response, err error) (*Lobby, error) {
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var state Lobby
	err = json.Unmarshal(body, &state)
	if err != nil {
		return nil, err
	}

	return &state, nil
}

func CreateLobby(baseURL string) (*Lobby, error) {
	response, err := http.Post(
		fmt.Sprintf("%s/lobby", baseURL),
		"application/json",
		strings.NewReader(""))

	return responseToLobby(response, err)
}

func JoinLobby(baseURL string, gameID string) (*Lobby, error) {
	response, err := http.Post(
		fmt.Sprintf("%s/lobby/%s", baseURL, gameID),
		"application/json",
		strings.NewReader(""))
	return responseToLobby(response, err)
}
