// The framework adapter contract. Adapters synthesize edges that exist in no
// source text: DI wiring, event bus publisher to subscriber, ORM models,
// generated clients. SCIP has no vocabulary for those, which is the only reason
// this contract exists separately from Indexer.
// Endpoints must be SCIP symbol strings. That single requirement is what lets a
// synthesized edge join the precise graph, and an adapter that invents its own
// identifiers produces edges that attach to nothing.

package index

import (
	"context"

	"github.com/andrewbelles/joist/internal/graph"
)

// Synthesis is what a framework adapter returns. Edges are already canonical,
// since their endpoints are SCIP symbols, so they bypass the scip package rather
// than being normalized like a precise index.
type Synthesis struct {
	Adapter  string
	Version  string
	Edges    []graph.Edge
	Warnings []string
}

// Adapter synthesizes framework edges for a Request. Adapters run out of process
// and may be written in any language, so the same hermeticity rules apply and are
// harder to enforce: no network, no repo state outside the declared file set, no
// timestamps or absolute paths in the output.
type Adapter interface {
	Name() string
	Version() string
	Synthesize(ctx context.Context, req Request) (Synthesis, error)
}
