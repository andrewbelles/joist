# adapters

Out-of-process framework adapters. Each one synthesizes edges that exist in no
source text: DI container wiring, event bus publisher to subscriber, ORM models,
generated clients.

These are separate processes, so they may be written in any language. That is the
point of running them out of process: a Spring adapter can be Java without Java
becoming a dependency of the core binary.

## What does not belong here

Precise indexers are not adapters. `scip-java`, `scip-typescript`, `scip-python`,
`scip-clang` and rust-analyzer already emit SCIP, so there is no contract to
define for them and nothing to write in this directory. They are invoked directly
by `internal/index` against the `Indexer` interface.

This directory exists only because SCIP has no vocabulary for framework
synthesized coupling. Anything expressible in SCIP should be expressed in SCIP.

## Contract

An adapter is a subprocess that reads one JSON request on stdin and writes one
JSON response on stdout. It exits zero on success. Diagnostics go to stderr.

The Go side is `internal/index.Adapter`. Keep the two in sync; a field added here
has to be added there.

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
  "edges": [
    {
      "from": "scip-java maven com.example/service 1.0.0 com/example/Publisher#emit().",
      "to": "scip-java maven com.example/service 1.0.0 com/example/Listener#onEvent().",
      "kind": "framework",
      "sites": 1
    }
  ],
  "warnings": []
}
```

## Endpoints must be SCIP symbol strings

This is the requirement the whole contract rests on. `from` and `to` are SCIP
symbols in the standard grammar:

```
<scheme> <manager> <package-name> <version> <descriptor>
```

An adapter that invents its own identifiers produces edges that join to nothing.
The synthesized edge is only useful because it lands on the same node the precise
index already produced, so emitting a class name, a file path, or a bean id
instead of a SCIP symbol makes the output worthless rather than merely wrong.

Because the endpoints are already canonical, synthesized edges do not pass
through `internal/scip`. There is nothing to normalize. That is the difference
between this contract and a precise index, and it is why the two are separate
interfaces rather than one interface returning opaque bytes.

## Requirements

An adapter must be hermetic. Same request, same response, on any machine at any
time. Concretely it must not read the network, must not read repo state outside
the declared file set, and must not embed a timestamp, an absolute path, or a
hostname in its output. Violating any of these makes the cache key stop
determining the result, and every cache hit silently becomes a guess.

Paths in the response are repo relative. `root` is provided so the adapter can
open files, never so it can echo an absolute path back.

There is no confidence field. Confidence describes whether a real build backed an
index, which is a property of the precise path only. A framework adapter either
resolves an edge or does not emit it.
