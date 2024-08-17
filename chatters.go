package semanticrouter

import "context"

// ChatMessage is a struct that represents a chat message.
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Chatter is an interface that defines a method, Chat, which takes a string
// as input and returns a string as output.
//
// It is used to generate text based on a given prompt.
//
// If the context is canceled, the context's error is returned if it is non-nil.
type Chatter interface {
	Chat(ctx context.Context, prompt []ChatMessage) (string, error)
}
