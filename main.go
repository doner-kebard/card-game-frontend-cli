package main

import (
	"fmt"
)

const DEV_URL string = "http://localhost:3000"

func main() {
	lobby, err := CreateLobby(DEV_URL)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(lobby)
	}
	lobby2, err := JoinLobby(DEV_URL, lobby.GameID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(lobby2)
	}
}
