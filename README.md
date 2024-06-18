# go-semantic-router

<p align="center">
    <a href="https://pkg.go.dev/github.com/conneroisu/go-semantic-router?tab=doc"><img src="https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white" alt="go.dev"></a>
    <a href="https://github.com/conneroisu/go-semantic-router/actions/workflows/test.yaml"><img src="https://github.com/conneroisu/go-semantic-router/actions/workflows/test.yaml/badge.svg" alt="Build Status"></a>
    <a href="https://codecov.io/gh/conneroisu/go-semantic-router" > <img src="https://codecov.io/gh/conneroisu/go-semantic-router/graph/badge.svg?token=JAGYI2V82D"/> </a>
    <a href="https://goreportcard.com/report/github.com/conneroisu/go-semantic-router"><img src="https://goreportcard.com/badge/github.com/conneroisu/go-semantic-router" alt="Go Report Card"></a>
    <a href="https://www.phorm.ai/query?projectId=fd665f24-5c41-42ed-907b-f322457a562d"><img src="https://img.shields.io/badge/Phorm-Ask_AI-%23F2777A.svg?&logo=data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNSIgaGVpZ2h0PSI0IiBmaWxsPSJub25lIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPgogIDxwYXRoIGQ9Ik00LjQzIDEuODgyYTEuNDQgMS40NCAwIDAgMS0uMDk4LjQyNmMtLjA1LjEyMy0uMTE1LjIzLS4xOTI"
</p>

Go Semantic Router is a superfast decision-making layer for your LLMs and agents.

Rather than waiting for slow LLM generations to make tool-use decisions, use the magic of semantic vector space to make those decisions — routing requests using semantic meaning.

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
	"github.com/conneroisu/go-semantic-router/encoders"
)

// NoteworthyRoutes represents a set of routes that are noteworthy.
// noteworthy here means that the routes are likely to be relevant to a noteworthy conversation in a veterinarian appointment.
var NoteworthyRoutes = semantic_router.Route{
	Name: "noteworthy",
	Utterances: []semantic_router.Utterance{
		{Utterance: "what is the best way to treat a dog with a cold?"},
		{Utterance: "my cat has been limping, what should I do?"},
	},
}

var ChitchatRoutes = semantic_router.Route{
	Name: "chitchat",
	Utterances: []semantic_router.Utterance{
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
	router, err := semantic_router.NewRouter(
		[]semantic_router.Route{NoteworthyRoutes, ChitchatRoutes},
		encoders.OpenAIEncoder{
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

```
Found: chitchat
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
	},
}

var ChitchatRoutes = semantic_router.Route{
	Name: "chitchat",
	Utterances: []semantic_router.Utterance{
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
	router, err := semantic_router.NewRouter(
		[]semantic_router.Route{NoteworthyRoutes, ChitchatRoutes},
		encoders.OpenAIEncoder{
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
Found: chitchat
p: 0.7631368810166642
```
