// Package main shows how to use the semantic router to find the best route for a given utterance
// in the context of a chat bot or other conversational application.
package main

import (
	"context"
	"fmt"
	"os"

	semantic_router "github.com/conneroisu/go-semantic-router"
	"github.com/conneroisu/go-semantic-router/domain"
	encoders "github.com/conneroisu/go-semantic-router/encoders/openai"
	"github.com/conneroisu/go-semantic-router/stores/memory"
)

// PoliticsRoutes represents a set of routes that are noteworthy.
var PoliticsRoutes = semantic_router.Route{
	Name: "politics",
	Utterances: []domain.Utterance{
		{Utterance: "isn't politics the best thing ever"},
		{Utterance: "why don't you tell me about your political opinions"},
		{Utterance: "don't you just love the president"},
		{Utterance: "they're going to destroy this country!"},
		{Utterance: "they will save the country!"},
	},
}

// ChitchatRoutes represents a set of routes that are noteworthy.
var ChitchatRoutes = semantic_router.Route{
	Name: "chitchat",
	Utterances: []domain.Utterance{
		{Utterance: "how's the weather today?"},
		{Utterance: "how are things going?"},
		{Utterance: "lovely weather today"},
		{Utterance: "the weather is horrendous"},
		{Utterance: "let's go to the chippy"},
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
	store := memory.NewStore()

	router, err := semantic_router.NewRouter(
		[]semantic_router.Route{PoliticsRoutes, ChitchatRoutes},
		encoders.OpenAIEncoder{
			APIKey: os.Getenv("OPENAI_API_KEY"),
			Model:  "text-embedding-3-small",
		},
		store,
	)
	if err != nil {
		return fmt.Errorf("error creating router: %w", err)
	}
	finding, p, err := router.Match(ctx, "how's the weather today?")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("p:", p)
	fmt.Println("Found:", finding)
	finding, p, err = router.Match(ctx, "ain't politics the best thing ever")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("p:", p)
	fmt.Println("Found:", finding)
	return nil
}
