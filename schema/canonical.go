// Canonical encoding of an artifact. Identical analyses must produce identical
// bytes or the content addressed cache silently stops hitting, so this is the
// only permitted way to serialize an Artifact.
// Rules: sorted map keys, stable slice ordering by ID, repo relative paths, no
// timestamps, no absolute paths, no host or user identifiers.

package schema

import "errors"

// ErrNotImplemented marks the scaffold. Remove it with the first real encoder.
var ErrNotImplemented = errors.New("schema: not implemented")

// Encode writes the canonical byte representation of a. Callers hash the result
// to get the artifact address.
func Encode(a *Artifact) ([]byte, error) {
	return nil, ErrNotImplemented
}

// Decode parses canonical bytes. It rejects an unknown Version rather than
// tolerating missing fields.
func Decode(b []byte) (*Artifact, error) {
	return nil, ErrNotImplemented
}

// Normalize sorts every slice in a into canonical order. Encode calls it, so
// callers building an Artifact by hand do not have to.
func Normalize(a *Artifact) {
}
