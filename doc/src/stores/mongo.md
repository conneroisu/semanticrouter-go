# MongoDB Store

The MongoDB store uses MongoDB to persist embeddings.

*Coming soon: Full documentation*

## Basic Usage

```go
import "github.com/conneroisu/semanticrouter-go/stores/mongo"

store, err := mongo.NewStore(ctx, "mongodb://localhost:27017", "database", "collection")
```