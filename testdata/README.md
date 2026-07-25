# testdata

Golden fixtures that cross package boundaries: small repositories with a known
graph, their expected canonical artifact bytes, and the arch.yaml that produced
them.

Fixtures scoped to one package belong in that package's own `testdata/`
directory, per Go convention. This directory is only for end-to-end cases.

Golden artifact bytes are compared exactly, not field by field. That is
deliberate: the canonical encoding is what the cache addresses, so a diff in
serialization is a real failure even when the decoded values match.

Regenerate goldens explicitly, never automatically on failure.
