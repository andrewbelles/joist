// Package scip collapses every indexer's output into one canonical symbol graph.
// It owns SCIP symbol identity and cross language symbol resolution.
// No indexer specific behavior may leak past this package. Anything downstream
// that branches on which indexer produced an edge is a bug here.
package scip
