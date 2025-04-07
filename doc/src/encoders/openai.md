# OpenAI Encoder

The OpenAI encoder uses OpenAI's API to generate text embeddings.

*Coming soon: Full documentation*

## Basic Usage

```go
import "github.com/conneroisu/semanticrouter-go/encoders/closedai"

encoder := &closedai.Encoder{
    APIKey: "your-api-key",
    Model: "text-embedding-ada-002",
}
```