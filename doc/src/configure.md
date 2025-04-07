# Configuration

## Basic Configuration

The Semantic Router can be configured using various options when creating a new router instance:

```go
router, err := semanticrouter.NewRouter(
    []semanticrouter.Route{route1, route2},
    encoder,
    store,
    // Optional configuration options
    semanticrouter.WithSimilarityDotMatrix(1.0),
    semanticrouter.WithWorkers(4),
)
```

## Routes

Routes represent the destinations to which utterances can be matched:

```go
route := semanticrouter.Route{
    Name: "greeting",
    Utterances: []semanticrouter.Utterance{
        {Utterance: "hello there"},
        {Utterance: "hi, how are you?"},
        {Utterance: "good morning"},
    },
}
```

## Encoders

Encoders convert text into vector embeddings. The library supports multiple encoder implementations:

### Ollama

```go
import (
    "github.com/conneroisu/semanticrouter-go/encoders/ollama"
    "github.com/ollama/ollama/api"
)

client, _ := api.ClientFromEnvironment()
encoder := &ollama.Encoder{
    Client: client,
    Model: "mxbai-embed-large",
}
```

### OpenAI

```go
import "github.com/conneroisu/semanticrouter-go/encoders/closedai"

encoder := &closedai.Encoder{
    APIKey: "your-api-key",
    Model: "text-embedding-ada-002",
}
```

### Google

```go
import "github.com/conneroisu/semanticrouter-go/encoders/google"

encoder := &google.Encoder{
    APIKey: "your-api-key",
    Model: "textembedding-gecko@latest",
}
```

### VoyageAI

```go
import "github.com/conneroisu/semanticrouter-go/encoders/voyageai"

encoder := &voyageai.Encoder{
    APIKey: "your-api-key",
    Model: "voyage-large-2",
}
```

## Storage

Storage interfaces allow you to persist vector embeddings. The library provides multiple storage options:

### In-Memory Storage

```go
import "github.com/conneroisu/semanticrouter-go/stores/memory"

store := memory.NewStore()
```

### MongoDB Storage

```go
import "github.com/conneroisu/semanticrouter-go/stores/mongo"

store, err := mongo.NewStore(ctx, "mongodb://localhost:27017", "database", "collection")
```

### Valkey Storage

```go
import "github.com/conneroisu/semanticrouter-go/stores/valkey"

store, err := valkey.NewStore("localhost:6379", "", 0)
```

## Similarity Methods

The router supports multiple similarity calculation methods that can be used together:

```go
router, err := semanticrouter.NewRouter(
    routes,
    encoder,
    store,
    semanticrouter.WithSimilarityDotMatrix(1.0),     // Cosine similarity (default)
    semanticrouter.WithEuclideanDistance(0.5),       // Euclidean distance
    semanticrouter.WithManhattanDistance(0.3),       // Manhattan/city block distance
    semanticrouter.WithJaccardSimilarity(0.2),       // Jaccard similarity
    semanticrouter.WithPearsonCorrelation(0.1),      // Pearson correlation coefficient
)
```

## Concurrency

You can configure the number of worker goroutines used for computing similarity scores:

```go
semanticrouter.WithWorkers(8) // Use 8 worker goroutines
```