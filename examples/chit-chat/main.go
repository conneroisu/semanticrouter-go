package main

import (
	"fmt"
	"os"

	semantic_router "github.com/conneroisu/go-semantic-router"
	"github.com/conneroisu/go-semantic-router/providers"
)

var PoliticsRoutes = semantic_router.Routes{
	Name: "politics",
	Utterances: []string{
		"isn't politics the best thing ever",
		"why don't you tell me about your political opinions",
		"don't you just love the president",
		"they're going to destroy this country!",
		"they will save the country!",
	},
}

var ChitchatRoutes = semantic_router.Routes{
	Name: "chitchat",
	Utterances: []string{
		"how's the weather today?",
		"how are things going?",
		"lovely weather today",
		"the weather is horrendous",
		"let's go to the chippy",
	},
}

func main() {
	router := semantic_router.NewRouter(
		[]semantic_router.Routes{PoliticsRoutes, ChitchatRoutes},
		providers.OpenAIEncoder{os.Getenv("OPENAI_API_KEY")},
	)
	finding, err := router.Match("how's the weather today?")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Found:", finding)
}
