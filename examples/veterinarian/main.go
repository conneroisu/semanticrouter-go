package main

import (
	"fmt"
	"os"

	semantic_router "github.com/conneroisu/go-semantic-router"
	"github.com/conneroisu/go-semantic-router/providers"
)

// NoteworthyRoutes represents a set of routes that are noteworthy.
// noteworthy here means that the routes are likely to be relevant to a noteworthy conversation in a veterinarian appointment.
var NoteworthyRoutes = semantic_router.Route{
	Name: "noteworthy",
	Utterances: []semantic_router.Utterance{
		{Utterance: "what is the best way to treat a dog with a cold?"},
		{Utterance: "my cat has been limping, what should I do?"},
		{Utterance: "how often should I vaccinate my pet?"},
		{Utterance: "what are the symptoms of parvovirus in dogs?"},
		{Utterance: "my rabbit isn't eating, is that normal?"},
		{Utterance: "can you recommend a good diet for an overweight cat?"},
		{Utterance: "what should I do if my dog is having a seizure?"},
		{Utterance: "how can I tell if my pet is in pain?"},
		{Utterance: "what are the side effects of this medication?"},
		{Utterance: "is it safe to give human medication to my pet?"},
		{Utterance: "how do I prevent fleas and ticks?"},
		{Utterance: "my dog has been vomiting, what could be wrong?"},
		{Utterance: "what's the best way to clean my cat's ears?"},
		{Utterance: "how do I know if my pet has allergies?"},
		{Utterance: "what should I feed my puppy?"},
		{Utterance: "how can I tell if my pet is overweight?"},
		{Utterance: "what should I do if my pet ingests something toxic?"},
		{Utterance: "can you explain the spaying/neutering process?"},
		{Utterance: "how can I make my pet more comfortable during travel?"},
		{Utterance: "what vaccinations does my pet need?"},
		{Utterance: "how do I deal with a pet that has separation anxiety?"},
		{Utterance: "what are the signs of dental problems in pets?"},
		{Utterance: "how often should I bathe my dog?"},
		{Utterance: "my pet has a lump, should I be concerned?"},
		{Utterance: "what should I do if my pet has diarrhea?"},
		{Utterance: "how can I improve my pet's coat and skin health?"},
		{Utterance: "what are the signs of heartworm in dogs?"},
		{Utterance: "how do I train my pet to use a litter box?"},
		{Utterance: "my pet is drinking a lot of water, is that normal?"},
		{Utterance: "what are the signs of ear infections in pets?"},
		{Utterance: "how can I help my pet lose weight?"},
		{Utterance: "what should I do if my pet is not eating?"},
	},
}

var ChitchatRoutes = semantic_router.Route{
	Name: "chitchat",
	Utterances: []semantic_router.Utterance{
		{Utterance: "what is your favorite color?"},
		{Utterance: "what is your favorite animal?"},
		{Utterance: "what is your favorite food?"},
		{Utterance: "what is your favorite movie?"},
		{Utterance: "what is that on your shirt?"},
		{Utterance: "what is that on your pants?"},
		{Utterance: "what is that on your shoes?"},
		{Utterance: "what is that on your hat?"},
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
	router, err := semantic_router.NewRouter(
		[]semantic_router.Route{NoteworthyRoutes, ChitchatRoutes},
		providers.OpenAIEncoder{
			APIKey: os.Getenv("OPENAI_API_KEY"),
			Model:  "text-embedding-3-small",
		},
	)
	finding, p, err := router.Match("how's the weather today?")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Found:", finding)
	fmt.Println("p:", p)
	return nil
}
