# In-Memory Store

The in-memory store is the simplest storage backend for Semantic Router, keeping all embeddings in memory.

## Installation

The in-memory store is included with Semantic Router:

```bash
go get github.com/conneroisu/semanticrouter-go/stores/memory
```

## Configuration

Using the in-memory store is straightforward:

```go
import "github.com/conneroisu/semanticrouter-go/stores/memory"

// Create a new in-memory store
store := memory.NewStore()
```

The in-memory store doesn't require any additional configuration.

## Example Usage

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/conneroisu/semanticrouter-go"
    "github.com/conneroisu/semanticrouter-go/stores/memory"
)

func main() {
    // Create a new in-memory store
    store := memory.NewStore()
    
    // Store an embedding
    ctx := context.Background()
    err := store.Set(ctx, semanticrouter.Utterance{
        Utterance: "hello world",
        Embed:     []float64{0.1, 0.2, 0.3, 0.4, 0.5},
    })
    if err != nil {
        log.Fatalf("Failed to store embedding: %v", err)
    }
    
    // Retrieve the embedding
    embedding, err := store.Get(ctx, "hello world")
    if err != nil {
        log.Fatalf("Failed to get embedding: %v", err)
    }
    
    fmt.Printf("Retrieved embedding: %v\n", embedding)
    
    // Close the store when done
    store.Close()
}
```

## Characteristics

### Advantages

- **Very fast**: No network or disk I/O overhead
- **Zero dependencies**: No external services required
- **Simple**: Easiest store to set up and use
- **Thread-safe**: Implements mutex locking for concurrent access

### Limitations

- **Non-persistent**: Data is lost when the program exits
- **Limited capacity**: Bounded by available RAM
- **Single-instance**: Can't be shared across multiple processes

## Best Practices

The in-memory store is ideal for:

1. Development and testing environments
2. Simple applications with a small number of routes
3. Applications where startup time isn't critical (routes need to be re-encoded on startup)
4. Low-latency requirements where every millisecond counts

For production use with many routes or where persistence is important, consider using the MongoDB or Valkey stores instead.