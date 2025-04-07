# Frequently Asked Questions

## General Questions

### What is Semantic Router?

Semantic Router is a Go library that routes requests based on semantic meaning rather than explicit rules or keywords. It uses vector embeddings to understand the meaning of text and route it to the appropriate handler.

### How does it compare to traditional routing approaches?

Traditional routing approaches rely on explicit rules, regular expressions, or keyword matching. Semantic Router uses the meaning of the text itself, allowing for more flexible and natural language understanding.

## Technical Questions

### Which embedding models work best?

For optimal performance, we recommend using models specifically designed for embeddings.
For example:

- OpenAI: `text-embedding-3-large` or `text-embedding-3-small`
- Ollama: `mxbai-embed-large` or `nomic-embed-text`
- Google: `textembedding-gecko` 
- VoyageAI: `voyage-large-2`

### How many example utterances should I provide per route?

We recommend at least 3-5 utterances per route for good performance. More examples generally lead to better results, especially for routes with diverse semantic content.

### How can I optimize performance?

- Use the `WithWorkers` option to increase concurrency
- Use the in-memory store for fastest performance
- Consider precomputing and caching embeddings for known routes
- Use a smaller embedding model if latency is critical

### Which similarity method should I use?

The default dot product similarity (cosine similarity) works well for most cases. For specific use cases:

- Dot product similarity: Good default for most text similarity tasks
- Euclidean distance: Better for when the magnitude of vectors matters
- Manhattan distance: Less sensitive to outliers than Euclidean
- Jaccard similarity: Good for comparing sets of items
- Pearson correlation: Good for finding linear relationships

### Can I create my own encoder or storage backend?

Yes! You can implement the `Encoder` interface to create a custom encoder and the `Store` interface to create a custom storage backend.

## Troubleshooting

### My router isn't matching correctly

Check the following:

1. Ensure your example utterances are diverse and representative
2. Try different similarity methods or adjust their coefficients
3. Make sure your embedding model is appropriate for the task
4. Add more example utterances for challenging routes

### I'm getting "context canceled" errors

This usually means the context passed to the `Match` function was canceled before the operation completed. Make sure your context has an appropriate timeout and isn't being canceled prematurely.
