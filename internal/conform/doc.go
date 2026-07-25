// Package conform is the deterministic path: rule checking, violation ranking,
// and the disagreement quadrant that crosses structural edges with co-change.
// Its output is reproducible bit for bit and is what CI gates on.
// It must never import internal/explore. Advisory signal entering here would
// make the artifact unreproducible and the CI gate unstable.
package conform
