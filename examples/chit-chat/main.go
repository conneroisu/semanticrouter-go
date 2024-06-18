// Package main shows how to use the semantic router to find the best route for a given utterance
// in the context of a chat bot or other conversational application.
package main

import (
	"fmt"
	"log"
	"os"

	semantic_router "github.com/conneroisu/go-semantic-router"
	"github.com/conneroisu/go-semantic-router/domain"

	"github.com/conneroisu/go-semantic-router/encoders"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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
	endpoint := "play.min.io"
	accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	router, err := semantic_router.NewRouter(
		[]semantic_router.Route{PoliticsRoutes, ChitchatRoutes},
		encoders.OpenAIEncoder{
			APIKey: os.Getenv("OPENAI_API_KEY"),
			Model:  "text-embedding-3-small",
		},
	)
	if err != nil {
		return fmt.Errorf("error creating router: %w", err)
	}
	finding, p, err := router.Match("how's the weather today?")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("p:", p)
	fmt.Println("Found:", finding)
	finding, p, err = router.Match("ain't politics the best thing ever")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("p:", p)
	fmt.Println("Found:", finding)
	return nil
}
