// The precise indexer contract. There is no custom wire format here because SCIP
// already defines one. An Indexer is scip-java, scip-typescript, scip-python,
// scip-clang, rust-analyzer or the tree-sitter fallback, so this package only has
// to know how to invoke one and where its output lands.

package index

import "context"

// Index is raw SCIP output. It stays bytes at this boundary; decoding and symbol
// resolution belong to the scip package, which is what keeps protobuf out of the
// orchestrator.
type Index struct {
	Indexer    string
	Version    string
	Confidence Confidence
	SCIP       [][]byte
	Warnings   []string
}

// Indexer produces a precise index for a Request. Implementations must be
// hermetic: same Request, same output, no network, no ambient repo state, no
// wall clock.
type Indexer interface {
	Name() string
	Version() string
	Index(ctx context.Context, req Request) (Index, error)
}
