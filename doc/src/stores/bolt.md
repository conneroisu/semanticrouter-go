# BoltDB Store

The BoltDB store provides persistent storage for embeddings using [BoltDB](https://github.com/etcd-io/bbolt), a fast key-value store written in Go.

## Features

- Persistent, file-based storage of embeddings
- Fast read and write operations
- Thread-safe access
- Configurable bucket name
- Context-aware operations with cancellation support

## Installation

```bash
go get github.com/conneroisu/semanticrouter-go/stores/bolt
```

## Usage

```go
import (
    "github.com/conneroisu/semanticrouter-go"
    "github.com/conneroisu/semanticrouter-go/encoders/ollama"
    "github.com/conneroisu/semanticrouter-go/stores/bolt"
    "github.com/ollama/ollama/api"
)

func main() {
    // Create a new BoltDB store
    store, err := bolt.NewStore("embeddings.db")
    if err != nil {
        log.Fatalf("Failed to create BoltDB store: %v", err)
    }
    defer store.Close()

    // Create a client for Ollama
    client, err := api.ClientFromEnvironment()
    if err != nil {
        log.Fatalf("Failed to create Ollama client: %v", err)
    }

    // Create a router with the BoltDB store
    router, err := semanticrouter.NewRouter(
        []semanticrouter.Route{
            {
                Name: "greeting",
                Utterances: []semanticrouter.Utterance{
                    {Utterance: "hello there"},
                    {Utterance: "good morning"},
                },
            },
            {
                Name: "farewell",
                Utterances: []semanticrouter.Utterance{
                    {Utterance: "goodbye"},
                    {Utterance: "see you later"},
                },
            },
        },
        &ollama.Encoder{
            Client: client,
            Model:  "mxbai-embed-large",
        },
        store,
    )
    if err != nil {
        log.Fatalf("Failed to create router: %v", err)
    }

    // Match an utterance
    match, score, err := router.Match(context.Background(), "hi there")
    if err != nil {
        log.Fatalf("Failed to match utterance: %v", err)
    }

    fmt.Printf("Match: %s, Score: %f\n", match.Name, score)
}
```

## Configuration Options

The BoltDB store can be configured with the following options:

### Custom Bucket Name

By default, the BoltDB store uses a bucket named "embeddings". You can specify a custom bucket name:

```go
store, err := bolt.NewStore(
    "embeddings.db",
    bolt.WithBucketName("custom-bucket"),
)
```

## Considerations

- BoltDB only allows one write transaction at a time, but allows multiple concurrent read transactions.
- For high-concurrency applications, consider using the Memory store with periodic snapshots to BoltDB.