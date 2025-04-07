# VoyageAI Encoder

The VoyageAI encoder uses VoyageAI's API to generate text embeddings.

*Coming soon: Full documentation*

## Basic Usage

```go
import "github.com/conneroisu/semanticrouter-go/encoders/voyageai"

encoder := &voyageai.Encoder{
    APIKey: "your-api-key",
    Model: "voyage-large-2",
}
```