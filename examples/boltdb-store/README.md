# BoltDB Store Example

This example demonstrates how to use the BoltDB store with the semantic router for persistent storage of embeddings.

## Requirements

- Go 1.24 or higher
- Ollama installed and running locally (for the embedding model)

## Running the Example

1. Ensure Ollama is running and that you have the `mxbai-embed-large` model available:

```bash
ollama pull mxbai-embed-large
```

2. Run the example:

```bash
go run main.go
```

## What the Example Does

1. Creates a BoltDB store for persistent storage of embeddings
2. Initializes a router with two routes: "weather" and "travel"
3. Matches several test utterances against these routes
4. Prints the best matching route and score for each utterance

## Expected Output

```
Utterance: 'what's it like outside today?'
Best route: weather
Score: 0.8732

Utterance: 'how can I get to the train station?'
Best route: travel
Score: 0.9126

Utterance: 'will I need an umbrella later?'
Best route: weather
Score: 0.8245
```

## How It Works

The BoltDB store saves embeddings to a file called `embeddings.db`. On the first run, it will encode all the example utterances and store them in the database. On subsequent runs, it will reuse the stored embeddings instead of re-encoding them.

This persistence makes the semantic router much faster after the initial setup, as it doesn't need to re-encode known utterances.