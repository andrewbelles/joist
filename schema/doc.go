// Package schema defines the conformance artifact, the one file every consumer
// reads. It is a separate Go module so that CI tooling can depend on the format
// without depending on the analyzer.
// This module must never take a third party dependency.
package schema
