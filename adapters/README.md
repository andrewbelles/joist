# adapters

Out-of-process framework adapters. Each one synthesizes edges that exist in no
source text: DI container wiring, event bus publisher to subscriber, ORM models,
generated clients.

These are separate processes, so they may be written in any language. That is the
point of running them out of process: a Spring adapter can be Java without Java
becoming a dependency of the core binary.

## Contract

An adapter is a subprocess that reads one JSON request on stdin and writes one
JSON response on stdout. It exits zero on success. Diagnostics go to stderr.

The Go side of this contract is `internal/index.Adapter`. Keep the two in sync;
a field added here has to be added there.

Request:

```json
{
  "root": "/abs/path/to/worktree",
  "files": ["src/a.java", "src/b.java"],
  "flags": ["-parameters"],
  "toolchain": { "javac": "21.0.4" }
}
```

Response:

```json
{
  "adapter": "spring",
  "version": "0.1.0",
  "confidence": "precise",
  "edges": [
    { "from": "scip-java ...#a().", "to": "scip-java ...#b().", "kind": "framework", "sites": 1 }
  ],
  "warnings": []
}
```

## Requirements

An adapter must be hermetic. Same request, same response, on any machine at any
time. Concretely it must not read the network, must not read repo state outside
the declared file set, and must not embed a timestamp, an absolute path, or a
hostname in its output. Violating any of these makes the cache key stop
determining the result, and every cache hit silently becomes a guess.

Paths in the response are repo relative. `root` is provided so the adapter can
open files, never so it can echo an absolute path back.
