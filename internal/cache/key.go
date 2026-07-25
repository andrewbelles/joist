// Cache keys. The key is the hash of everything an indexing action reads. If an
// input is not in the key, a stale result will be served and no later check will
// catch it, so adding an input to an action means adding it here too.

package cache

import "errors"

// ErrNotImplemented marks the scaffold.
var ErrNotImplemented = errors.New("cache: not implemented")

// Key is a hex encoded content address. Fixed size, so comparing two arbitrarily
// large objects is comparing two of these.
type Key string

// Inputs is the complete declared input set of one action. Every field
// participates in the key.
type Inputs struct {
	Content        []byte
	Flags          []string
	ToolchainID    string
	IndexerID      string
	IndexerVersion string
}

// Compute returns the key for in. It is a pure function; identical Inputs must
// produce an identical Key on any machine at any time.
func Compute(in Inputs) (Key, error) {
	return "", ErrNotImplemented
}

// Hash returns the content address of raw bytes, used to store a result under
// the address of its own serialized form.
func Hash(b []byte) Key {
	return ""
}
