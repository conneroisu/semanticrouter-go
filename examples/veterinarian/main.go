// Package main shows how to use the semantic router to find the best route for a given utterance
// in the context of a veterinarian appointment.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/conneroisu/semanticrouter-go"
	"github.com/conneroisu/semanticrouter-go/encoders/ollama"
	"github.com/conneroisu/semanticrouter-go/stores/memory"
	"github.com/ollama/ollama/api"
)

// NoteworthyRoutes represents a set of routes that are noteworthy.
// noteworthy here means that the routes are likely to be relevant to a noteworthy conversation in a veterinarian appointment.
var NoteworthyRoutes = semanticrouter.Route{
	Name: "noteworthy",
	Utterances: []semanticrouter.Utterance{
		{Utterance: "what is the best way to treat a dog with a cold?"},
		{Utterance: "my cat has been limping, what should I do?"},
	},
}

// ChitchatRoutes represents a set of routes that are chitchat.
// chitchat here means that the routes are likely to be relevant to a chitchat conversation in a veterinarian appointment.
var ChitchatRoutes = semanticrouter.Route{
	Name: "chitchat",
	Utterances: []semanticrouter.Utterance{
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
	ctx := context.Background()
	cli, err := api.ClientFromEnvironment()
	if err != nil {
		return fmt.Errorf("error creating client: %w", err)
	}
	router, err := semanticrouter.NewRouter(
		[]semanticrouter.Route{NoteworthyRoutes, ChitchatRoutes},
		&ollama.Encoder{
			Client: cli,
			Model:  "mxbai-embed-large",
		},
		memory.NewStore(),
	)
	if err != nil {
		return fmt.Errorf("error creating router: %w", err)
	}
	finding, p, err := router.Match(ctx, "how's the weather today?")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Found:", finding.Name)
	fmt.Println("p:", p)
	return nil
}
