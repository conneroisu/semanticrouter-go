module github.com/conneroisu/semanticrouter-go/examples/boltdb-store

go 1.24.0

require (
	github.com/conneroisu/semanticrouter-go v0.9.4
	github.com/conneroisu/semanticrouter-go/encoders/ollama v0.0.0-20250519171221-384fa0cd4b21
	github.com/conneroisu/semanticrouter-go/stores/bolt v0.0.0
	github.com/ollama/ollama v0.3.10
)

replace github.com/conneroisu/semanticrouter-go/stores/bolt => ../../stores/bolt

require (
	go.etcd.io/bbolt v1.3.9 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	gonum.org/v1/gonum v0.15.0 // indirect
)
