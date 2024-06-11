package semanticrouter

// Route represents a route in the semantic router.
type Route struct {
	Name       string   `json:"name" yaml:"name" toml:"name"`
	Utterances []string `json:"utterances" yaml:"utterances" toml:"utterances"`
}
