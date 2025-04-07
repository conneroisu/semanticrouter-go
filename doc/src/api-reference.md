# API Reference

This page provides a reference for the most important types and functions in the Semantic Router library.

## Core Types

### Router

```go
type Router struct {
    Routes  []Route
    Encoder Encoder
    Storage Store
    // contains filtered or unexported fields
}
```

The `Router` is the main struct for the semantic router. It contains the routes, encoder, and storage.

#### Creating a Router

```go
func NewRouter(routes []Route, encoder Encoder, store Store, opts ...Option) (router *Router, err error)
```

Creates a new semantic router with the given routes, encoder, and storage. Additional options can be provided to configure the router.

#### Matching an Utterance

```go
func (r *Router) Match(ctx context.Context, utterance string) (bestRoute *Route, bestScore float64, err error)
```

Matches the given utterance against the routes in the router and returns the best matching route and its score.

### Route

```go
type Route struct {
    Name       string
    Utterances []Utterance
}
```

A `Route` represents a possible destination for an utterance. It contains a name and a list of example utterances.

### Utterance

```go
type Utterance struct {
    ID        int
    Utterance string
    Embed     []float64
}
```

An `Utterance` represents a text string and its embedding vector.

## Interfaces

### Encoder

```go
type Encoder interface {
    Encode(ctx context.Context, utterance string) ([]float64, error)
}
```

The `Encoder` interface is implemented by encoder drivers to convert text into vector embeddings.

### Store

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

The `Store` interface is implemented by storage backends to store and retrieve embeddings.

## Configuration Options

### WithSimilarityDotMatrix

```go
func WithSimilarityDotMatrix(coefficient float64) Option
```

Sets the similarity dot matrix function with the given coefficient. This implements cosine similarity.

### WithEuclideanDistance

```go
func WithEuclideanDistance(coefficient float64) Option
```

Sets the Euclidean distance function with the given coefficient.

### WithManhattanDistance

```go
func WithManhattanDistance(coefficient float64) Option
```

Sets the Manhattan distance function with the given coefficient.

### WithJaccardSimilarity

```go
func WithJaccardSimilarity(coefficient float64) Option
```

Sets the Jaccard similarity function with the given coefficient.

### WithPearsonCorrelation

```go
func WithPearsonCorrelation(coefficient float64) Option
```

Sets the Pearson correlation function with the given coefficient.

### WithWorkers

```go
func WithWorkers(workers int) Option
```

Sets the number of worker goroutines to use for computing similarity scores.

## Error Types

### ErrNoRouteFound

```go
type ErrNoRouteFound struct {
    Message   string
    Utterance string
}
```

Error returned when no route is found for an utterance.

### ErrEncoding

```go
type ErrEncoding struct {
    Message string
}
```

Error returned when an error occurs during encoding.

### ErrGetEmbedding

```go
type ErrGetEmbedding struct {
    Message string
    Storage Store
}
```

Error returned when an error occurs while retrieving an embedding from storage.