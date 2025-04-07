# Introduction to Semantic Router

Go Semantic Router is a decision-making layer for your large language models (LLMs) and agents written in pure Go. It routes requests using semantic meaning rather than waiting for slow LLM generations to make tool-use decisions.

## What is Semantic Routing?

Semantic routing uses vector embeddings to understand and categorize the meaning of text. Instead of relying on keyword matching or complex rules, it captures the semantic essence of queries and routes them to the appropriate handlers based on similarity in vector space.

## Key Features

- **Fast Decision Making**: Make routing decisions in milliseconds rather than waiting for LLM generations
- **Multiple Similarity Methods**: Support for various similarity metrics including:
  - Dot Product Similarity
  - Euclidean Distance
  - Manhattan Distance
  - Jaccard Similarity
  - Pearson Correlation
- **Pluggable Architecture**: 
  - Multiple encoder options (Ollama, OpenAI, Google, VoyageAI)
  - Various vector storage backends (in-memory, MongoDB, Valkey)
- **Concurrent Processing**: Process multiple routes in parallel for maximum performance

## Use Cases

- **Conversational Agents**: Route user queries to appropriate response handlers
- **Tool Selection**: Determine which tools an agent should use based on query semantics
- **Intent Classification**: Classify user intents without relying on slow LLM reasoning
- **Multi-agent Systems**: Route requests to specialized agents based on query content