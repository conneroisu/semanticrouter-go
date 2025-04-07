# Basic Usage

## Quick Start

Here's a simple example of using the semantic router:

```go
package main

import (
    "context"
    "fmt"

    "github.com/conneroisu/semanticrouter-go"
    "github.com/conneroisu/semanticrouter-go/encoders/ollama"
    "github.com/conneroisu/semanticrouter-go/stores/memory"
    "github.com/ollama/ollama/api"
)

func main() {
    // Create routes
    greetingRoute := semanticrouter.Route{
        Name: "greeting",
        Utterances: []semanticrouter.Utterance{
            {Utterance: "hello there"},
            {Utterance: "hi, how are you?"},
            {Utterance: "good morning"},
        },
    }

    weatherRoute := semanticrouter.Route{
        Name: "weather",
        Utterances: []semanticrouter.Utterance{
            {Utterance: "what's the weather like?"},
            {Utterance: "is it going to rain today?"},
            {Utterance: "temperature forecast"},
        },
    }

    // Create encoder (you need Ollama running locally)
    client, _ := api.ClientFromEnvironment()
    encoder := &ollama.Encoder{
        Client: client,
        Model:  "mxbai-embed-large",
    }

    // Create store
    store := memory.NewStore()

    // Create router
    router, err := semanticrouter.NewRouter(
        []semanticrouter.Route{greetingRoute, weatherRoute},
        encoder,
        store,
    )
    if err != nil {
        panic(err)
    }

    // Match an utterance
    ctx := context.Background()
    bestRoute, score, err := router.Match(ctx, "hi there, how are you doing?")
    if err != nil {
        panic(err)
    }

    fmt.Printf("Best route: %s with score: %.4f\n", bestRoute.Name, score)
}
```

## Understanding the Results

The `Match` function returns three values:

1. `bestRoute *Route` - A pointer to the best matching route
2. `score float64` - The similarity score (higher is better)
3. `err error` - An error, if any occurred

The score is calculated based on the similarity methods configured. By default, the router uses cosine similarity (via dot product similarity).

## Handling No Match

If no suitable route is found, the router will return an `ErrNoRouteFound` error:

```go
bestRoute, score, err := router.Match(ctx, "completely unrelated text")
if err != nil {
    if _, ok := err.(semanticrouter.ErrNoRouteFound); ok {
        fmt.Println("No matching route found!")
    } else {
        fmt.Printf("Error: %v\n", err)
    }
}
```

## Multiple Similarity Methods

You can configure the router to use multiple similarity methods:

```go
router, err := semanticrouter.NewRouter(
    routes,
    encoder,
    store,
    semanticrouter.WithSimilarityDotMatrix(0.6),
    semanticrouter.WithEuclideanDistance(0.4),
)
```

In this example, the final score will be a weighted combination of both similarity methods.