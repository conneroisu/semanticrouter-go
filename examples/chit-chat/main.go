// Package main shows how to use the semantic router to find the best route for a given utterance
// in the context of a chat bot or other conversational application.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/conneroisu/semanticrouter-go"
	opeenai "github.com/conneroisu/semanticrouter-go/encoders/openai"
	"github.com/conneroisu/semanticrouter-go/stores/memory"
	"github.com/sashabaranov/go-openai"
)

// PoliticsRoutes represents a set of routes that are noteworthy.
var PoliticsRoutes = semanticrouter.Route{
	Name: "politics",
	Utterances: []semanticrouter.Utterance{
		{Utterance: "isn't politics the best thing ever"},
		{Utterance: "why don't you tell me about your political opinions"},
		{Utterance: "don't you just love the president"},
		{Utterance: "they're going to destroy this country!"},
		{Utterance: "they will save the country!"},
	},
}

// ChitchatRoutes represents a set of routes that are noteworthy.
var ChitchatRoutes = semanticrouter.Route{
	Name: "chitchat",
	Utterances: []semanticrouter.Utterance{
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
	router, err := semanticrouter.NewRouter(
		[]semanticrouter.Route{
			PoliticsRoutes,
			ChitchatRoutes,
		},
		opeenai.Encoder{
			Client: openai.NewClient(os.Getenv("OPENAI_API_KEY")),
			Model:  openai.AdaCodeSearchCode,
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
	return nil
}
