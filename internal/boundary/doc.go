// Package boundary parses the hand authored arch.yaml and resolves its globs to
// module membership over the symbol graph.
// It answers which module a symbol belongs to. It does not evaluate rules
// against edges; that is the conform package.
package boundary
