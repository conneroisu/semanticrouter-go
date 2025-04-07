# Valkey Store

The Valkey store uses Valkey (Redis-compatible) to persist embeddings.

*Coming soon: Full documentation*

## Basic Usage

```go
import "github.com/conneroisu/semanticrouter-go/stores/valkey"

store, err := valkey.NewStore("localhost:6379", "", 0)
```