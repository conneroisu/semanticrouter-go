// Package main shows how to use the semantic router to find the best route for a given utterance
// in the context of a veterinarian appointment.
package main

import (
	"fmt"
	"os"

	semantic_router "github.com/conneroisu/go-semantic-router"
	"github.com/conneroisu/go-semantic-router/domain"
	"github.com/conneroisu/go-semantic-router/encoders/ollama"
	"github.com/conneroisu/go-semantic-router/stores/memory"
	"github.com/ollama/ollama/api"
)

// NoteworthyRoutes represents a set of routes that are noteworthy.
// noteworthy here means that the routes are likely to be relevant to a noteworthy conversation in a veterinarian appointment.
var NoteworthyRoutes = semantic_router.Route{
	Name: "noteworthy",
	Utterances: []domain.Utterance{
		{Utterance: "what is the best way to treat a dog with a cold?"},
		{Utterance: "my cat has been limping, what should I do?"},
	},
}

// ChitchatRoutes represents a set of routes that are chitchat.
// chitchat here means that the routes are likely to be relevant to a chitchat conversation in a veterinarian appointment.
var ChitchatRoutes = semantic_router.Route{
	Name: "chitchat",
	Utterances: []domain.Utterance{
		{Utterance: "what is your favorite color?"},
		{Utterance: "what is your favorite animal?"},
	},
}

// main runs the example.
func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// run runs the example.
func run() error {
	store := memory.NewStore()
	cli, err := api.ClientFromEnvironment()
	router, err := semantic_router.NewRouter(
		[]semantic_router.Route{NoteworthyRoutes, ChitchatRoutes},
		&ollama.Encoder{
			Client: cli,
		},
		store,
	)
	finding, p, err := router.Match("how's the weather today?")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Found:", finding)
	fmt.Println("p:", p)
	return nil
}
