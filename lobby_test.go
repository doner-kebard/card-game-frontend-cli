package main

import (
	"fmt"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
)

func TestLobbyCreate(t *testing.T) {

	pact := &dsl.Pact{
		Consumer: "CLI",
		Provider: "Lobby",
	}
	defer pact.Teardown()

	var test = func() error {
		url := fmt.Sprintf("http://localhost:%d", pact.Server.Port)

		_, err := CreateLobby(url)

		return err
	}

	pact.
		AddInteraction().
		Given("The server is online").
		UponReceiving("A request to create a lobby").
		WithRequest(dsl.Request{
			Method:  "POST",
			Path:    dsl.String("/lobby"),
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
		}).
		WillRespondWith(dsl.Response{
			Status:  200,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			Body:    dsl.Match(&Lobby{}),
		})

	if err := pact.Verify(test); err != nil {
		t.Errorf("Error on Verify: %v", err)
	}
}
