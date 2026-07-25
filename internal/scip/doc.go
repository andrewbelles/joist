// Package scip decodes raw SCIP output and collapses every indexer into one
// canonical symbol graph. It owns SCIP symbol identity and cross language symbol
// resolution. Framework adapter edges do not pass through here; their endpoints
// are already SCIP symbols, so there is nothing to normalize.
// No indexer specific behavior may leak past this package. Anything downstream
// that branches on which indexer produced an edge is a bug here.
package scip
