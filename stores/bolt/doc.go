// Package bolt provides a BoltDB-based implementation of the semanticrouter Store interface.
//
// BoltDB is a simple, fast, reliable key-value store written in pure Go. This implementation
// provides persistent storage of embeddings using BoltDB (specifically the etcd fork bbolt).
//
// Example usage:
//
//	import (
//	    "github.com/conneroisu/semanticrouter-go/stores/bolt"
//	)
//
//	// Create a new BoltDB store
//	store, err := bolt.NewStore("embeddings.db")
//	if err != nil {
//	    log.Fatalf("Failed to create BoltDB store: %v", err)
//	}
//	defer store.Close()
//
//	// Use the store with a router
//	router, err := semanticrouter.NewRouter(
//	    []semanticrouter.Route{...},
//	    encoder,
//	    store,
//	)
package bolt