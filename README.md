# go-semantic-router

<p align="center">
    <a href="https://pkg.go.dev/github.com/conneroisu/go-semantic-router?tab=doc"><img src="https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white" alt="go.dev"></a>
    <a href="https://github.com/conneroisu/go-semantic-router/actions/workflows/test.yaml"><img src="https://github.com/conneroisu/go-semantic-router/actions/workflows/test.yaml/badge.svg" alt="Build Status"></a>
    <a href="https://codecov.io/gh/conneroisu/go-semantic-router" > <img src="https://codecov.io/gh/conneroisu/go-semantic-router/graph/badge.svg?token=JAGYI2V82D"/> </a>
    <a href="https://goreportcard.com/report/github.com/conneroisu/go-semantic-router"><img src="https://goreportcard.com/badge/github.com/conneroisu/go-semantic-router" alt="Go Report Card"></a>
    <a href="https://www.phorm.ai/query?projectId=3e6f9d42-0098-4178-ab54-4a0b9c89353b"><img src="https://img.shields.io/badge/Phorm-Ask_AI-%23F2777A.svg?&logo=data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNSIgaGVpZ2h0PSI0IiBmaWxsPSJub25lIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPgogIDxwYXRoIGQ9Ik00LjQzIDEuODgyYTEuNDQgMS40NCAwIDAgMS0uMDk4LjQyNmMtLjA1LjEyMy0uMTE1LjIzLS4xOTIuMzIyLS4wNzUuMDktLjE2LjE2NS0uMjU1LjIyNmExLjM1MyAxLjM1MyAwIDAgMS0uNTk1LjIxMmMtLjA5OS4wMTItLjE5Mi4wMTQtLjI3OS4wMDZsLTEuNTkzLS4xNHYtLjQwNmgxLjY1OGMuMDkuMDAxLjE3LS4xNjkuMjQ2LS4xOTFhLjYwMy42MDMgMCAwIDAgLjItLjEwNi41MjkuNTI5IDAgMCAwIC4xMzgtLjE3LjY1NC42NTQgMCAwIDAgLjA2NS0uMjRsLjAyOC0uMzJhLjkzLjkzIDAgMCAwLS4wMzYtLjI0OS41NjcuNTY3IDAgMCAwLS4xMDMtLjIuNTAyLjUwMiAwIDAgMC0uMTY4LS4xMzguNjA4LjYwOCAwIDAgMC0uMjQtLjA2N0wyLjQzNy43MjkgMS42MjUuNjcxYS4zMjIuMzIyIDAgMCAwLS4yMzIuMDU4LjM3NS4zNzUgMCAwIDAtLjExNi4yMzJsLS4xMTYgMS40NS0uMDU4LjY5Ny0uMDU4Ljc1NEwuNzA1IDRsLS4zNTctLjA3OUwuNjAyLjkwNkMuNjE3LjcyNi42NjMuNTc0LjczOS40NTRhLjk1OC45NTggMCAwIDEgLjI3NC0uMjg1Ljk3MS45NzEgMCAwIDEgLjMzNy0uMTRjLjExOS0uMDI2LjIyNy0uMDM0LjMyNS0uMDI2TDMuMjMyLjE2Yy4xNTkuMDE0LjMzNi4wMy40NTkuMDgyYTEuMTczIDEuMTczIDAgMCAxIC41NDUuNDQ3Yy4wNi4wOTQuMTA5LjE5Mi4xNDQuMjkzYTEuMzkyIDEuMzkyIDAgMCAxIC4wNzguNThsLS4wMjkuMzJaIiBmaWxsPSIjRjI3NzdBIi8+CiAgPHBhdGggZD0iTTQuMDgyIDIuMDA3YTEuNDU1IDEuNDU1IDAgMCAxLS4wOTguNDI3Yy0uMDUuMTI0LS4xMTQuMjMyLS4xOTIuMzI0YTEuMTMgMS4xMyAwIDAgMS0uMjU0LjIyNyAxLjM1MyAxLjM1MyAwIDAgMS0uNTk1LjIxNGMtLjEuMDEyLS4xOTMuMDE0LS4yOC4wMDZsLTEuNTYtLjEwOC4wMzQtLjQwNi4wMy0uMzQ4IDEuNTU5LjE1NGMuMDkgMCAuMTczLS4wMS4yNDgtLjAzM2EuNjAzLjYwMyAwIDAgMCAuMi0uMTA2LjUzMi41MzIgMCAwIDAgLjEzOS0uMTcyLjY2LjY2IDAgMCAwIC4wNjQtLjI0MWwuMDI5LS4zMjFhLjk0Ljk0IDAgMCAwLS4wMzYtLjI1LjU3LjU3IDAgMCAwLS4xMDMtLjIwMi41MDIuNTAyIDAgMCAwLS4xNjgtLjEzOC42MDUuNjA1IDAgMCAwLS4yNC0uMDY3TDEuMjczLjgyN2MtLjA5NC0uMDA4LS4xNjguMDEtLjIyMS4wNTUtLjA1My4wNDUtLjA4NC4xMTQtLjA5Mi4yMDZMLjcwNSA0IDAgMy45MzhsLjI1NS0yLjkxMUExLjAxIDEuMDEgMCAwIDEgLjM5My41NzIuOTYyLjk2MiAwIDAgMSAuNjY2LjI4NmEuOTcuOTcgMCAwIDEgLjMzOC0uMTRDMS4xMjIuMTIgMS4yMy4xMSAxLjMyOC4xMTlsMS41OTMuMTRjLjE2LjAxNC4zLjA0Ny40MjMuMWExLjE3IDEuMTcgMCAwIDEgLjU0NS40NDhjLjA2MS4wOTUuMTA5LjE5My4xNDQuMjk1YTEuNDA2IDEuNDA2IDAgMCAxIC4wNzcuNTgzbC0uMDI4LjMyMloiIGZpbGw9IndoaXRlIi8+CiAgPHBhdGggZD0iTTQuMDgyIDIuMDA3YTEuNDU1IDEuNDU1IDAgMCAxLS4wOTguNDI3Yy0uMDUuMTI0LS4xMTQuMjMyLS4xOTIuMzI0YTEuMTMgMS4xMyAwIDAgMS0uMjU0LjIyNyAxLjM1MyAxLjM1MyAwIDAgMS0uNTk1LjIxNGMtLjEuMDEyLS4xOTMuMDE0LS4yOC4wMDZsLTEuNTYtLjEwOC4wMzQtLjQwNi4wMy0uMzQ4IDEuNTU5LjE1NGMuMDkgMCAuMTczLS4wMS4yNDgtLjAzM2EuNjAzLjYwMyAwIDAgMCAuMi0uMTA2LjUzMi41MzIgMCAwIDAgLjEzOS0uMTcyLjY2LjY2IDAgMCAwIC4wNjQtLjI0MWwuMDI5LS4zMjFhLjk0Ljk0IDAgMCAwLS4wMzYtLjI1LjU3LjU3IDAgMCAwLS4xMDMtLjIwMi41MDIuNTAyIDAgMCAwLS4xNjgtLjEzOC42MDUuNjA1IDAgMCAwLS4yNC0uMDY3TDEuMjczLjgyN2MtLjA5NC0uMDA4LS4xNjguMDEtLjIyMS4wNTUtLjA1My4wNDUtLjA4NC4xMTQtLjA5Mi4yMDZMLjcwNSA0IDAgMy45MzhsLjI1NS0yLjkxMUExLjAxIDEuMDEgMCAwIDEgLjM5My41NzIuOTYyLjk2MiAwIDAgMSAuNjY2LjI4NmEuOTcuOTcgMCAwIDEgLjMzOC0uMTRDMS4xMjIuMTIgMS4yMy4xMSAxLjMyOC4xMTlsMS41OTMuMTRjLjE2LjAxNC4zLjA0Ny40MjMuMWExLjE3IDEuMTcgMCAwIDEgLjU0NS40NDhjLjA2MS4wOTUuMTA5LjE5My4xNDQuMjk1YTEuNDA2IDEuNDA2IDAgMCAxIC4wNzcuNTgzbC0uMDI4LjMyMloiIGZpbGw9IndoaXRlIi8+Cjwvc3ZnPgo="></img>
</p>

Go Semantic Router is a superfast decision-making layer for your LLMs and agents.

Rather than waiting for slow LLM generations to make tool-use decisions, use the magic of semantic vector space to make those decisions — routing our requests using semantic meaning.

A pure-go package for abstractly computing similarity scores between a query vector embedding and a set of vector embeddings.

## Installation

```bash
go get github.com/conneroisu/go-semantic-router
```

## Usage

### Conversational Agents Example

```go
package main

import (
	"fmt"
	"os"

	semantic_router "github.com/conneroisu/go-semantic-router"
	"github.com/conneroisu/go-semantic-router/providers"
)

// PoliticsRoutes represents a set of routes that are noteworthy.
var PoliticsRoutes = semantic_router.Route{
	Name: "politics",
	Utterances: []semantic_router.Utterance{
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
	Utterances: []semantic_router.Utterance{
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
	router, err := semantic_router.NewRouter(
		[]semantic_router.Route{PoliticsRoutes, ChitchatRoutes},
		providers.OpenAIEncoder{
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
```

Output:

```
Found: politics
```

### Veterinarian Example

```go

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
		{Utterance: "how are you?"},
		{Utterance: "what can you do for me?"},
		{Utterance: "you seen the new movie?"},
		{Utterance: "what is the weather?"},
		{Utterance: "how are you feeling?"},
		{Utterance: "what is your name?"},
		{Utterance: "what is your favorite movie?"},
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
```

Output:
```bash
p: 0.999899
Found: noteworthy
p: 0.999899
Found: noteworthy
```
