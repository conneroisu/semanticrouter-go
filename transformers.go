package semanticrouter

import "context"

// Chatter is an interface that defines a method, Chat, which takes a string
type Chatter interface {
	Chat(ctx context.Context, utterance string) (string, error)
}
