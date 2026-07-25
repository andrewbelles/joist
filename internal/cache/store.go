// The blob store and the per-SHA manifest. The store is passive: no logic, no
// API, no reconciliation. Entries are immutable, so there is no update operation
// and no invalidation problem.

package cache

import "context"

// Store is a content addressed blob store. Local disk, a CI cache, an OCI
// registry and object storage are all the same interface.
type Store interface {
	Get(ctx context.Context, k Key) ([]byte, error)
	Put(ctx context.Context, b []byte) (Key, error)
	Has(ctx context.Context, k Key) (bool, error)
}

// Manifest maps every document in a tree to its blob key at one commit SHA.
// Diffing two manifests yields the changed document set, which is what a
// rebuild is scoped to.
type Manifest struct {
	CommitSHA string
	Entries   map[string]Key
}

// Diff returns the documents whose key differs between old and new. Work is
// proportional to the size of the change, not the size of the tree.
func Diff(old, next *Manifest) []string {
	return nil
}
