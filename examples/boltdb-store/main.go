// Package main demonstrates using the BoltDB store with the semantic router.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/conneroisu/semanticrouter-go"
	"github.com/conneroisu/semanticrouter-go/encoders/ollama"
	"github.com/conneroisu/semanticrouter-go/stores/bolt"
	"github.com/ollama/ollama/api"
)

// WeatherRoutes represents routes related to weather.
var WeatherRoutes = semanticrouter.Route{
	Name: "weather",
	Utterances: []semanticrouter.Utterance{
		{Utterance: "how's the weather today?"},
		{Utterance: "will it rain tomorrow?"},
		{Utterance: "is it going to be sunny this weekend?"},
		{Utterance: "what's the temperature outside?"},
	},
}

// TravelRoutes represents routes related to travel.
var TravelRoutes = semanticrouter.Route{
	Name: "travel",
	Utterances: []semanticrouter.Utterance{
		{Utterance: "how do I get to the airport?"},
		{Utterance: "what time does the train leave?"},
		{Utterance: "I need directions to downtown"},
		{Utterance: "is there a bus to the museum?"},
	},
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run() error {
	// Create a BoltDB store
	dbPath := "embeddings.db"
	store, err := bolt.NewStore(dbPath)
	if err != nil {
		return fmt.Errorf("failed to create BoltDB store: %w", err)
	}
	defer func() {
		store.Close()
		// Clean up the database file (in a real app, you would keep this file)
		os.Remove(dbPath)
	}()

	// Create a client for Ollama
	client, err := api.ClientFromEnvironment()
	if err != nil {
		return fmt.Errorf("error creating Ollama client: %w", err)
	}

	// Create a router with the BoltDB store
	router, err := semanticrouter.NewRouter(
		[]semanticrouter.Route{WeatherRoutes, TravelRoutes},
		&ollama.Encoder{
			Client: client,
			Model:  "mxbai-embed-large", // You may need to adjust the model name
		},
		store,
	)
	if err != nil {
		return fmt.Errorf("error creating router: %w", err)
	}

	// Test utterances to match
	testUtterances := []string{
		"what's it like outside today?",
		"how can I get to the train station?",
		"will I need an umbrella later?",
	}

	ctx := context.Background()
	for _, utterance := range testUtterances {
		bestRoute, score, err := router.Match(ctx, utterance)
		if err != nil {
			fmt.Printf("Error matching '%s': %v\n", utterance, err)
			continue
		}

		fmt.Printf("Utterance: '%s'\nBest route: %s\nScore: %.4f\n\n", 
			utterance, bestRoute.Name, score)
	}

	return nil
}