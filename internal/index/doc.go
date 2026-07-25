// Package index orchestrates the indexer pool and defines the adapter contract
// that both precise indexers and out of process framework adapters implement.
// It runs subprocesses and collects their output. It does not interpret symbols;
// that is the scip package.
package index
