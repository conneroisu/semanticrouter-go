package semanticrouter

// ErrNoRouteFound is an error that is returned when no route is found.
type ErrNoRouteFound struct {
	Message   string
	Utterance string
}

// Error returns the error message.
func (e ErrNoRouteFound) Error() string {
	return e.Message + " : utterance : " + e.Utterance
}

// ErrEncoding is an error that is returned when an error occurs during encoding.
type ErrEncoding struct {
	Message string
}

// Error returns the error message.
func (e ErrEncoding) Error() string {
	return e.Message
}

// ErrGetEmbedding is an error that is returned when an error occurs during getting an embedding.
type ErrGetEmbedding struct {
	Message string
	Storage Store
}

// Error returns the error message.
func (e ErrGetEmbedding) Error() string {
	return e.Message
}
