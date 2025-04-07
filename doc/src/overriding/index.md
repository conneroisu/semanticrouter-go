# Extending Semantic Router

This section explains how to extend Semantic Router with custom implementations of core interfaces.

## Creating a Custom Encoder

To create a custom encoder, implement the `Encoder` interface:

```go
type Encoder interface {
    Encode(ctx context.Context, utterance string) ([]float64, error)
}
```

Example implementation using a hypothetical embedding service:

```go
package myencoder

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "strings"
)

type MyEncoder struct {
    APIKey string
    URL    string
}

func (e *MyEncoder) Encode(ctx context.Context, utterance string) ([]float64, error) {
    req, err := http.NewRequestWithContext(
        ctx, 
        "POST", 
        e.URL,
        strings.NewReader(fmt.Sprintf(`{"text":"%s"}`, utterance)),
    )
    if err != nil {
        return nil, err
    }
    
    req.Header.Add("Authorization", "Bearer "+e.APIKey)
    req.Header.Add("Content-Type", "application/json")
    
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    var result struct {
        Embedding []float64 `json:"embedding"`
    }
    
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }
    
    return result.Embedding, nil
}
```

## Creating a Custom Storage Backend

To create a custom storage backend, implement the `Store` interface:

```go
type Store interface {
    Setter
    Getter
    io.Closer
}

type Setter interface {
    Set(ctx context.Context, keyValPair Utterance) error
}

type Getter interface {
    Get(ctx context.Context, key string) ([]float64, error)
}
```

Example implementation using SQLite:

```go
package sqlitestore

import (
    "context"
    "database/sql"
    "encoding/json"
    
    "github.com/conneroisu/semanticrouter-go"
    _ "github.com/mattn/go-sqlite3"
)

type Store struct {
    db *sql.DB
}

func NewStore(dbPath string) (*Store, error) {
    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        return nil, err
    }
    
    // Create table if it doesn't exist
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS embeddings 
        (utterance TEXT PRIMARY KEY, embedding BLOB)`)
    if err != nil {
        db.Close()
        return nil, err
    }
    
    return &Store{db: db}, nil
}

func (s *Store) Set(ctx context.Context, utterance semanticrouter.Utterance) error {
    embJSON, err := json.Marshal(utterance.Embed)
    if err != nil {
        return err
    }
    
    _, err = s.db.ExecContext(
        ctx,
        "INSERT OR REPLACE INTO embeddings (utterance, embedding) VALUES (?, ?)",
        utterance.Utterance,
        embJSON,
    )
    return err
}

func (s *Store) Get(ctx context.Context, key string) ([]float64, error) {
    var embJSON []byte
    err := s.db.QueryRowContext(
        ctx,
        "SELECT embedding FROM embeddings WHERE utterance = ?",
        key,
    ).Scan(&embJSON)
    if err != nil {
        return nil, err
    }
    
    var embedding []float64
    if err := json.Unmarshal(embJSON, &embedding); err != nil {
        return nil, err
    }
    
    return embedding, nil
}

func (s *Store) Close() error {
    return s.db.Close()
}
```

## Creating a Custom Similarity Method

You can define a custom similarity function and use it with the router:

```go
package main

import (
    "github.com/conneroisu/semanticrouter-go"
    "gonum.org/v1/gonum/mat"
)

// Define a custom similarity function
func customSimilarity(queryVec *mat.VecDense, indexVec *mat.VecDense) (float64, error) {
    // Implementation of your custom similarity algorithm
    // ...
    return score, nil
}

// Create an option to use this function
func WithCustomSimilarity(coefficient float64) semanticrouter.Option {
    return func(r *semanticrouter.Router) {
        r.AddComparisonFunc(customSimilarity, coefficient)
    }
}

// Use it with your router
router, err := semanticrouter.NewRouter(
    routes,
    encoder,
    store,
    WithCustomSimilarity(1.0),
)
```

## Best Practices

When extending the Semantic Router:

1. Handle context cancellation appropriately
2. Implement proper error handling and propagation
3. Make implementations thread-safe when appropriate
4. Consider performance implications, especially for high-traffic applications
5. Add unit tests for your custom implementations