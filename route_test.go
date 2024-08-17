package semanticrouter_test

import (
	"testing"

	"github.com/conneroisu/semanticrouter-go"
	"github.com/conneroisu/semanticrouter-go/encoders/ollama"
	"github.com/conneroisu/semanticrouter-go/stores/memory"
	"github.com/ollama/ollama/api"
	"github.com/stretchr/testify/assert"
)

// NoteworthyRoutes represents a set of routes that are noteworthy.
// noteworthy here means that the routes are likely to be relevant to a noteworthy conversation in a veterinarian appointment.
var NoteworthyRoutes = semanticrouter.Route{
	Name: "noteworthy",
	Utterances: []semanticrouter.Utterance{
		{Utterance: "what is the best way to treat a dog with a cold?"},
		{Utterance: "my cat has been limping, what should I do?"},
	},
}

// ChitchatRoutes represents a set of routes that are chitchat.
// chitchat here means that the routes are likely to be relevant to a chitchat conversation in a veterinarian appointment.
var ChitchatRoutes = semanticrouter.Route{
	Name: "chitchat",
	Utterances: []semanticrouter.Utterance{
		{Utterance: "what is your favorite color?"},
		{Utterance: "what is your favorite animal?"},
	},
}

// TestNewRouter tests the NewRouter function.
func TestNewRouter(t *testing.T) {
	a := assert.New(t)
	client, err := api.ClientFromEnvironment()
	if err != nil {
		t.Fatal(err)
	}
	stor := memory.NewStore()
	rout, err := semanticrouter.NewRouter(
		[]semanticrouter.Route{NoteworthyRoutes, ChitchatRoutes},
		ollama.NewEncoder(client, "all-minilm"),
		stor,
	)
	a.NoError(err)
	a.NotNil(rout)

}
