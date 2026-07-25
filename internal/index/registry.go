// Resolution from a language or framework name to the thing that handles it.
// Indexers and adapters are kept in separate namespaces because they answer
// different questions and a name collision between them is not a conflict.

package index

// Registry holds the registered indexers and adapters. It is populated at wiring
// time in cmd/joist and is read only afterwards.
type Registry struct {
	indexers map[string]Indexer
	adapters map[string]Adapter
}

// NewRegistry returns an empty registry.
func NewRegistry() *Registry {
	return &Registry{
		indexers: make(map[string]Indexer),
		adapters: make(map[string]Adapter),
	}
}

// AddIndexer registers i by name. A duplicate name is a wiring bug and returns an
// error rather than silently replacing.
func (r *Registry) AddIndexer(i Indexer) error {
	return ErrNotImplemented
}

// AddAdapter registers a by name. A duplicate name is a wiring bug and returns an
// error rather than silently replacing.
func (r *Registry) AddAdapter(a Adapter) error {
	return ErrNotImplemented
}

// Indexer returns the indexer registered under name.
func (r *Registry) Indexer(name string) (Indexer, bool) {
	i, ok := r.indexers[name]
	return i, ok
}

// Adapter returns the adapter registered under name.
func (r *Registry) Adapter(name string) (Adapter, bool) {
	a, ok := r.adapters[name]
	return a, ok
}
