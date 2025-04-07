# Google Encoder

The Google encoder uses Google's API to generate text embeddings.

*Coming soon: Full documentation*

## Basic Usage

```go
import "github.com/conneroisu/semanticrouter-go/encoders/google"

encoder := &google.Encoder{
    APIKey: "your-api-key",
    Model: "textembedding-gecko@latest",
}
```