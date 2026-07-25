// Package history mines git for co-change, normalized against an independent
// change null model, with exponential time decay and rename tracking.
// It reads first-parent commits on trunk and takes paths only, never diffs.
// It reads git and writes nothing.
package history
