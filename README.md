# semanticrouter-go

<p align="center">
    <a href="https://pkg.go.dev/github.com/conneroisu/semanticrouter-go?tab=doc"><img src="https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white" alt="go.dev"></a>
    <a href="https://github.com/conneroisu/semanticrouter-go/actions/workflows/test.yaml"><img src="https://github.com/conneroisu/semanticrouter-go/actions/workflows/test.yaml/badge.svg" alt="Build Status"></a>
    <a href="https://codecov.io/gh/conneroisu/semanticrouter-go" > <img src="https://codecov.io/gh/conneroisu/semanticrouter-go/graph/badge.svg?token=JAGYI2V82D"/> </a>
    <a href="https://goreportcard.com/report/github.com/conneroisu/semanticrouter-go"><img src="https://goreportcard.com/badge/github.com/conneroisu/semanticrouter-go" alt="Go Report Card"></a>
    <a href="https://www.phorm.ai/query?projectId=fd665f24-5c41-42ed-907b-f322457a562d"><img src="https://img.shields.io/badge/Phorm-Ask_AI-%23F2777A.svg?&logo=data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNSIgaGVpZ2h0PSI0IiBmaWxsPSJub25lIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPgogIDxwYXRoIGQ9Ik00LjQzIDEuODgyYTEuNDQgMS40NCAwIDAgMS0uMDk4LjQyNmMtLjA1LjEyMy0uMTE1LjIzLS4xOTI"
</p>

Go Semantic Router is a superfast decision-making layer for your LLMs and agents written in pure [ Go ](https://go.dev/).

Rather than waiting for slow LLM generations to make tool-use decisions, use the magic of semantic vector space to make those decisions — routing requests using configurable semantic meaning.

A pure-go package for abstractly computing similarity scores between a query vector embedding and a set of vector embeddings.

## Installation

```bash
go get github.com/conneroisu/semanticrouter-go
```

### Conversational Agents Example

```go
// Package main shows how to use the semantic router to find the best route for a given utterance
// in the context of a veterinarian appointment.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/conneroisu/semanticrouter-go"
	"github.com/conneroisu/semanticrouter-go/encoders/ollama"
	"github.com/conneroisu/semanticrouter-go/stores/memory"
	"github.com/ollama/ollama/api"
)

// NoteworthyRoutes represents a set of routes that are noteworthy.
// noteworthy here means that the routes are likely to be relevant to a noteworthy conversation in a veterinarian appointment.
var NoteworthyRoutes = semanticrouter.Route{
	Name: "noteworthy",
	Utterances: []semanticrouter.Utterance{
		{Utterance: "what is the best way to treat a dog with a cold?"},
		{Utterance: "my cat has been limping, what should I do?"},
	},
}

// ChitchatRoutes represents a set of routes that are chitchat.
// chitchat here means that the routes are likely to be relevant to a chitchat conversation in a veterinarian appointment.
var ChitchatRoutes = semanticrouter.Route{
	Name: "chitchat",
	Utterances: []semanticrouter.Utterance{
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
	ctx := context.Background()
	cli, err := api.ClientFromEnvironment()
	if err != nil {
		return fmt.Errorf("error creating client: %w", err)
	}
	router, err := semanticrouter.NewRouter(
		[]semanticrouter.Route{NoteworthyRoutes, ChitchatRoutes},
		&ollama.Encoder{
			Client: cli,
			Model:  "mxbai-embed-large",
		},
		memory.NewStore(),
	)
	if err != nil {
		return fmt.Errorf("error creating router: %w", err)
	}
	finding, p, err := router.Match(ctx, "how's the weather today?")
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

The following example shows how to use the semantic router to find the best route for a given utterance in the context of a veterinarian appointment.

The goal of the example is to decide whether spoken utterances are relevant to a noteworthy conversation or a chitchat conversation.


#### The Code Example:
```go
// Package main shows how to use the semantic router to find the best route for a given utterance
// in the context of a veterinarian appointment.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/conneroisu/semanticrouter-go"
	"github.com/conneroisu/semanticrouter-go/encoders/ollama"
	"github.com/conneroisu/semanticrouter-go/stores/memory"
	"github.com/ollama/ollama/api"
)

// NoteworthyRoutes represents a set of routes that are noteworthy.
// noteworthy here means that the routes are likely to be relevant to a noteworthy conversation in a veterinarian appointment.
var NoteworthyRoutes = semanticrouter.Route{
	Name: "noteworthy",
	Utterances: []semanticrouter.Utterance{
		{Utterance: "what is the best way to treat a dog with a cold?"},
		{Utterance: "my cat has been limping, what should I do?"},
	},
}

// ChitchatRoutes represents a set of routes that are chitchat.
// chitchat here means that the routes are likely to be relevant to a chitchat conversation in a veterinarian appointment.
var ChitchatRoutes = semanticrouter.Route{
	Name: "chitchat",
	Utterances: []semanticrouter.Utterance{
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
	ctx := context.Background()
	cli, err := api.ClientFromEnvironment()
	if err != nil {
		return fmt.Errorf("error creating client: %w", err)
	}
	router, err := semanticrouter.NewRouter(
		[]semanticrouter.Route{NoteworthyRoutes, ChitchatRoutes},
		&ollama.Encoder{
			Client: cli,
			Model:  "mxbai-embed-large",
		},
		memory.NewStore(),
	)
	if err != nil {
		return fmt.Errorf("error creating router: %w", err)
	}
	finding, p, err := router.Match(ctx, "how's the weather today?")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Found:", finding)
	fmt.Println("p:", p)
	return nil
}
```

#### The Output

The output of the veterinarian example is:
```bash
Found: chitchat
```

## Development

### Testing

To run the tests, run the following command:

```bash
make test
```

### Making a new Store Implementation

Implement the Store interface in the `stores` package.

The store interface is defined as follows:

```go
type Store interface { // size=16 (0x10)
	Storer
	Getter
	io.Closer
}
```

Store is an interface that defines a method, Store, which takes a \[]float64 and stores it in a some sort of data store, and a method, Get, which takes a string and returns a \[]float64 from the data store.

```go
func (io.Closer) Close() error
func (Getter) Get(ctx context.Context, key string) ([]float64, error)
func (Storer) Store(ctx context.Context, keyValPair Utterance) error
```

[`semanticrouter.Store` on pkg.go.dev](https://pkg.go.dev/github.com/conneroisu/semanticrouter-go#Store)


## Code Generated Documentation


<!-- gomarkdoc:embed:start -->

<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# semanticrouter

```go
import "github.com/conneroisu/semanticrouter-go"
```

Package semanticrouter is a package for abstractly computing similarity scores between a query vector embedding and a set of vector embeddings.

It provides a simple key\-value store for embeddings and a router that can be used to find the best route for a given utterance.

The router uses a similarity score to determine the best route for a given utterance.

The semantic router is designed to be used in conjunction with LLMs and agents to provide a superfast decision\-making layer.

## Index

- [type Encoder](<#Encoder>)
- [type ErrEncoding](<#ErrEncoding>)
  - [func \(e ErrEncoding\) Error\(\) string](<#ErrEncoding.Error>)
- [type ErrGetEmbedding](<#ErrGetEmbedding>)
  - [func \(e ErrGetEmbedding\) Error\(\) string](<#ErrGetEmbedding.Error>)
- [type ErrNoRouteFound](<#ErrNoRouteFound>)
  - [func \(e ErrNoRouteFound\) Error\(\) string](<#ErrNoRouteFound.Error>)
- [type Getter](<#Getter>)
- [type Option](<#Option>)
  - [func WithEuclideanDistance\(coefficient float64\) Option](<#WithEuclideanDistance>)
  - [func WithJaccardSimilarity\(coefficient float64\) Option](<#WithJaccardSimilarity>)
  - [func WithManhattanDistance\(coefficient float64\) Option](<#WithManhattanDistance>)
  - [func WithPearsonCorrelation\(coefficient float64\) Option](<#WithPearsonCorrelation>)
  - [func WithSimilarityDotMatrix\(coefficient float64\) Option](<#WithSimilarityDotMatrix>)
  - [func WithWorkers\(workers int\) Option](<#WithWorkers>)
- [type Route](<#Route>)
- [type Router](<#Router>)
  - [func NewRouter\(routes \[\]Route, encoder Encoder, store Store, opts ...Option\) \(router \*Router, err error\)](<#NewRouter>)
  - [func \(r \*Router\) Match\(ctx context.Context, utterance string\) \(bestRoute \*Route, bestScore float64, err error\)](<#Router.Match>)
- [type Setter](<#Setter>)
- [type Store](<#Store>)
- [type Utterance](<#Utterance>)


<a name="Encoder"></a>
## type [Encoder](<https://github.com/conneroisu/go-semantic-router/blob/main/encoder.go#L14-L16>)

Encoder represents a encoding driver in the semantic router.

It is an interface that defines a single method, Encode, which takes a string and returns a \[\]float64 representing the embedding of the string.

```go
type Encoder interface {
    Encode(ctx context.Context, utterance string) ([]float64, error)
}
```

<a name="ErrEncoding"></a>
## type [ErrEncoding](<https://github.com/conneroisu/go-semantic-router/blob/main/errors.go#L15-L17>)

ErrEncoding is an error that is returned when an error occurs during encoding.

```go
type ErrEncoding struct {
    Message string
}
```

<a name="ErrEncoding.Error"></a>
### func \(ErrEncoding\) [Error](<https://github.com/conneroisu/go-semantic-router/blob/main/errors.go#L20>)

```go
func (e ErrEncoding) Error() string
```

Error returns the error message.

<a name="ErrGetEmbedding"></a>
## type [ErrGetEmbedding](<https://github.com/conneroisu/go-semantic-router/blob/main/errors.go#L25-L28>)

ErrGetEmbedding is an error that is returned when an error occurs during getting an embedding.

```go
type ErrGetEmbedding struct {
    Message string
    Storage Store
}
```

<a name="ErrGetEmbedding.Error"></a>
### func \(ErrGetEmbedding\) [Error](<https://github.com/conneroisu/go-semantic-router/blob/main/errors.go#L31>)

```go
func (e ErrGetEmbedding) Error() string
```

Error returns the error message.

<a name="ErrNoRouteFound"></a>
## type [ErrNoRouteFound](<https://github.com/conneroisu/go-semantic-router/blob/main/errors.go#L4-L7>)

ErrNoRouteFound is an error that is returned when no route is found.

```go
type ErrNoRouteFound struct {
    Message   string
    Utterance string
}
```

<a name="ErrNoRouteFound.Error"></a>
### func \(ErrNoRouteFound\) [Error](<https://github.com/conneroisu/go-semantic-router/blob/main/errors.go#L10>)

```go
func (e ErrNoRouteFound) Error() string
```

Error returns the error message.

<a name="Getter"></a>
## type [Getter](<https://github.com/conneroisu/go-semantic-router/blob/main/encoder.go#L37-L39>)

Getter is an interface that defines a method, Get, which takes a string and returns a \[\]float64 from the data store.

If the key does not exist, it returns an error.

```go
type Getter interface {
    Get(ctx context.Context, key string) ([]float64, error)
}
```

<a name="Option"></a>
## type [Option](<https://github.com/conneroisu/go-semantic-router/blob/main/encoder.go#L42>)

Option is a function that configures a Router.

```go
type Option func(*Router)
```

<a name="WithEuclideanDistance"></a>
### func [WithEuclideanDistance](<https://github.com/conneroisu/go-semantic-router/blob/main/similarity.go#L65>)

```go
func WithEuclideanDistance(coefficient float64) Option
```

WithEuclideanDistance sets the EuclideanDistance function with a coefficient.

$$d\(x, y\) = \\sqrt\{\\sum\_\{i=1\}^\{n\}\(x\_i \- y\_i\)^2\}$$

<a name="WithJaccardSimilarity"></a>
### func [WithJaccardSimilarity](<https://github.com/conneroisu/go-semantic-router/blob/main/similarity.go#L95>)

```go
func WithJaccardSimilarity(coefficient float64) Option
```

WithJaccardSimilarity sets the JaccardSimilarity function with a coefficient.

$$J\(A, B\)=\\frac\{|A \\cap B|\}\{|A \\cup B|\}$$

It adds the jaccard similarity to the comparision functions with the given coefficient.

<a name="WithManhattanDistance"></a>
### func [WithManhattanDistance](<https://github.com/conneroisu/go-semantic-router/blob/main/similarity.go#L80>)

```go
func WithManhattanDistance(coefficient float64) Option
```

WithManhattanDistance sets the ManhattanDistance function with a coefficient.

$$d\(x, y\) = |x\_1 \- y\_1| \+ |x\_2 \- y\_2| \+ ... \+ |x\_n \- y\_n|$$

It adds the manhatten distance to the comparision functions with the given coefficient.

<a name="WithPearsonCorrelation"></a>
### func [WithPearsonCorrelation](<https://github.com/conneroisu/go-semantic-router/blob/main/similarity.go#L111>)

```go
func WithPearsonCorrelation(coefficient float64) Option
```

WithPearsonCorrelation sets the PearsonCorrelation function with a coefficient.

$$r=\\frac\{\\sum\\left\(x\_\{i\}\-\\bar\{x\}\\right\)\\left\(y\_\{i\}\-\\bar\{y\}\\right\)\}\{\\sqrt\{\\sum\\left\(x\_\{i\}\-\\bar\{x\}\\right\)^\{2\} \\sum\\left\(y\_\{i\}\-\\bar\{y\}\\right\)^\{2\}\}\}$$

It adds the pearson correlation to the comparision functions with the given coefficient.

<a name="WithSimilarityDotMatrix"></a>
### func [WithSimilarityDotMatrix](<https://github.com/conneroisu/go-semantic-router/blob/main/similarity.go#L53>)

```go
func WithSimilarityDotMatrix(coefficient float64) Option
```

WithSimilarityDotMatrix sets the similarity function to use with a coefficient.

$$a \\cdot b=\\sum\_\{i=1\}^\{n\} a\_\{i\} b\_\{i\}$$

It adds the similarity dot matrix to the comparision functions with the given coefficient.

<a name="WithWorkers"></a>
### func [WithWorkers](<https://github.com/conneroisu/go-semantic-router/blob/main/route.go#L26>)

```go
func WithWorkers(workers int) Option
```

WithWorkers sets the number of workers to use for computing similarity scores.

<a name="Route"></a>
## type [Route](<https://github.com/conneroisu/go-semantic-router/blob/main/route.go#L35-L38>)

Route represents a route in the semantic router.

It is a struct that contains a name and a slice of Utterances.

```go
type Route struct {
    Name       string      // Name is the name of the route.
    Utterances []Utterance // Utterances is a slice of Utterances.
}
```

<a name="Router"></a>
## type [Router](<https://github.com/conneroisu/go-semantic-router/blob/main/route.go#L16-L23>)

Router represents a semantic router.

Router is a struct that contains a slice of Routes and an Encoder.

Match can be called on a Router to find the best route for a given utterance.

```go
type Router struct {
    Routes  []Route // Routes is a slice of Routes.
    Encoder Encoder // Encoder is an Encoder that encodes utterances into vectors.
    Storage Store   // Storage is a Store that stores the utterances.
    // contains filtered or unexported fields
}
```

<a name="NewRouter"></a>
### func [NewRouter](<https://github.com/conneroisu/go-semantic-router/blob/main/route.go#L47-L52>)

```go
func NewRouter(routes []Route, encoder Encoder, store Store, opts ...Option) (router *Router, err error)
```

NewRouter creates a new semantic router.

<a name="Router.Match"></a>
### func \(\*Router\) [Match](<https://github.com/conneroisu/go-semantic-router/blob/main/route.go#L103-L106>)

```go
func (r *Router) Match(ctx context.Context, utterance string) (bestRoute *Route, bestScore float64, err error)
```

Match returns the route that matches the given utterance.

The score is the similarity score between the query vector and the index vector.

If the given context is canceled, the context's error is returned if it is non\-nil.

<a name="Setter"></a>
## type [Setter](<https://github.com/conneroisu/go-semantic-router/blob/main/encoder.go#L29-L31>)

Setter is an interface that defines a method, Store, which takes a \[\]float64 and stores it in a some sort of data store.

```go
type Setter interface {
    Set(ctx context.Context, keyValPair Utterance) error
}
```

<a name="Store"></a>
## type [Store](<https://github.com/conneroisu/go-semantic-router/blob/main/encoder.go#L21-L25>)

Store is an interface that defines a method, Store, which takes a \[\]float64 and stores it in a some sort of data store, and a method, Get, which takes a string and returns a \[\]float64 from the data store.

```go
type Store interface {
    Setter
    Getter
    io.Closer
}
```

<a name="Utterance"></a>
## type [Utterance](<https://github.com/conneroisu/go-semantic-router/blob/main/similarity.go#L15-L22>)

Utterance represents a utterance in the semantic router.

```go
type Utterance struct {
    // ID is the ID of the utterance.
    ID  int
    // Utterance is the text of the utterance.
    Utterance string
    // Embed is the embedding of the utterance. It is a vector of floats.
    Embed embedding
}
```

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- gomarkdoc:embed:end -->
