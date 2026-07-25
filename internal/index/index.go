// Vocabulary shared by both contracts in this package. A Request is the declared
// input set of one unit of work, so it is also what gets hashed into the cache
// key. Adding a field an implementation reads without adding it here makes the
// key stop determining the result.

package index

import "errors"

// ErrNotImplemented marks the scaffold.
var ErrNotImplemented = errors.New("index: not implemented")

// Confidence separates output recovered from a real build from output recovered
// by the syntactic fallback. Low confidence propagates into the artifact and
// stops findings from blocking a merge.
type Confidence int

const (
	// ConfidenceLow is the tree-sitter fallback, used when the build is broken.
	ConfidenceLow Confidence = iota
	// ConfidencePrecise is a real compile backed index.
	ConfidencePrecise
)

// Request is one unit of work. Files are repo relative and sorted, so two
// identical requests are byte identical and hash the same.
type Request struct {
	Root      string
	Files     []string
	Flags     []string
	Toolchain map[string]string
}
