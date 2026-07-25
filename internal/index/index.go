// The adapter contract. One Adapter turns a set of source files plus build
// metadata into index blobs. Adapters are interchangeable subprocesses, so this
// contract is the only thing the orchestrator knows about them.

package index

import (
	"context"
	"errors"
)

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

// Request is one unit of indexing work. Files are repo relative and sorted, so
// two identical requests are byte identical and hash the same.
type Request struct {
	Root      string
	Files     []string
	Flags     []string
	Toolchain map[string]string
}

// Response is what an adapter returns. Blobs are opaque here and are handed to
// the scip package for normalization.
type Response struct {
	Adapter    string
	Confidence Confidence
	Blobs      [][]byte
	Warnings   []string
}

// Adapter indexes a Request. Implementations must be hermetic: same Request,
// same Response, no network, no ambient repo state, no wall clock.
type Adapter interface {
	Name() string
	Version() string
	Index(ctx context.Context, req Request) (Response, error)
}

// Registry resolves a language or framework to its adapter. It is populated at
// wiring time in cmd/arch and is read only afterwards.
type Registry struct {
	adapters map[string]Adapter
}

// Register adds a by name. Registering a duplicate name is a wiring bug and
// returns an error rather than silently replacing.
func (r *Registry) Register(a Adapter) error {
	return ErrNotImplemented
}

// Lookup returns the adapter registered under name.
func (r *Registry) Lookup(name string) (Adapter, bool) {
	return nil, false
}
