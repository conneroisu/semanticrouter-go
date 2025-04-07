# Ollama Encoder

The Ollama encoder uses [Ollama](https://ollama.ai/) to generate text embeddings.

## Installation

First, ensure you have Ollama installed and running:

```bash
# Install Ollama
curl -fsSL https://ollama.ai/install.sh | sh

# Start Ollama
ollama serve
```

Then add the Ollama encoder to your Go project:

```bash
go get github.com/conneroisu/semanticrouter-go/encoders/ollama
```

## Configuration

```go
import (
    "github.com/conneroisu/semanticrouter-go/encoders/ollama"
    "github.com/ollama/ollama/api"
)

// Create client using environment variables (OLLAMA_HOST and OLLAMA_SCHEMA)
client, err := api.ClientFromEnvironment()
if err != nil {
    // Handle error
}

// Alternatively, create client with specific address
client, err := api.NewClient("http://localhost:11434")
if err != nil {
    // Handle error
}

// Create encoder
encoder := &ollama.Encoder{
    Client: client,
    Model: "mxbai-embed-large", // or another embedding model
}
```

## Supported Models

The Ollama encoder supports any embedding model available in Ollama. Here are some recommended models:

| Model | Size | Description |
|-------|------|-------------|
| `mxbai-embed-large` | ~1.5GB | High quality embeddings, excellent performance |
| `nomic-embed-text` | ~270MB | Good quality with smaller resource footprint |
| `all-minilm` | ~134MB | Small model with good quality for basic tasks |

Pull your desired model before using it:

```bash
ollama pull mxbai-embed-large
```

## Example Usage

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/conneroisu/semanticrouter-go/encoders/ollama"
    "github.com/ollama/ollama/api"
)

func main() {
    client, err := api.ClientFromEnvironment()
    if err != nil {
        log.Fatalf("Failed to create Ollama client: %v", err)
    }

    encoder := &ollama.Encoder{
        Client: client,
        Model:  "mxbai-embed-large",
    }

    ctx := context.Background()
    embedding, err := encoder.Encode(ctx, "Hello, world!")
    if err != nil {
        log.Fatalf("Failed to encode text: %v", err)
    }

    fmt.Printf("Embedding has %d dimensions\n", len(embedding))
}
```

## Error Handling

The Ollama encoder may return errors in the following situations:
- The Ollama service is not running
- The specified model is not available
- The context is canceled
- Network issues occur

Always handle these errors appropriately in production code.