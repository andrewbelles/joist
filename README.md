# call-graph-tooling

A build-level call graph analyzer that checks a codebase against the architecture
it claims to have.

Existing tools are mostly research grade, language scoped, and disconnected from
git. This one is wired at the build level, so coupling is resolved from compiled
output rather than guessed from source text, and it treats git as a first class
input: the graph is rebuilt incrementally per commit SHA in the way Bazel rebuilds
targets, and history is mined for co-change so evolutionary coupling can be
compared against structural coupling. It also reads what source-only tools drop,
including configs and framework wiring that carry real coupling between services.

Everything a run produces lands in one versioned, content-addressed artifact keyed
by commit SHA. The CLI, the CI check, the MCP server, and the viewer all read that
one file and nothing upstream of it. The MCP surface is the reason for the
precision: an agent that can ask who calls this and what is the blast radius in one
cheap, deterministic call does not have to guess, and does not have to read half
the repo to find out.

The whole thing runs on one developer machine or one CI runner. No server, no code
egress.

## Status

Early scaffold. The repository shape, the package boundaries, and the artifact
format are defined. No analysis is implemented; every subcommand returns a
not-implemented error.

## Build

Requires Go 1.26. The full build additionally requires Node 22.

```sh
make dev      # Go only, viewer serves a placeholder page
make build    # web build, then Go with -tags spa, full binary at bin/arch
make test lint fmt-check
```

The SPA is compiled in under the `spa` build tag. A plain `go build ./...` needs
neither Node nor a built `dist` tree, so the Go and TypeScript builds have no
coupling in either direction.

The binary is named `arch`, which shadows the coreutils `arch(1)` if installed to
a directory earlier in `PATH` than `/usr/bin`.

## Repo shape

| Path | Holds |
| --- | --- |
| `cmd/arch/` | Wiring only: flags, command tree, exit codes |
| `internal/index/` | Indexer orchestration and the adapter contract |
| `internal/scip/` | Normalization of every indexer into one canonical symbol graph |
| `internal/graph/` | Node and edge model, weighting. Pure data, no IO |
| `internal/history/` | Git mining, co-change |
| `internal/boundary/` | `arch.yaml` parsing, glob resolution to module membership |
| `internal/conform/` | Deterministic checker, violation ranker, disagreement quadrant |
| `internal/explore/` | Communities, smells, scattering. Advisory, offline only |
| `internal/cache/` | Content addressing, blob store, per-SHA manifests |
| `internal/mcpsrv/` | MCP server over stdio |
| `internal/viewer/` | Loopback HTTP server and the embedded SPA |
| `schema/` | The artifact format. Separate Go module |
| `adapters/` | Out-of-process framework adapters, any language |
| `web/` | TypeScript SPA source |
| `research/` | Python metric harness. Never shipped |
| `testdata/` | Cross-cutting golden fixtures |

Every Go package carries a `doc.go` stating its responsibility and its
prohibition. The prohibitions are load bearing and are the fastest way to learn
the shape.

Three splits matter more than the rest:

**Deterministic against exploratory.** `conform` is reproducible bit for bit and
is what CI gates on. `explore` is seed dependent, advisory, and never keyed into
the artifact. `conform` must never import `explore`.

**The artifact format is its own module.** `schema/` takes no third party
dependency, so CI tooling can read artifacts without depending on the analyzer.

**Python is research only.** Metrics are evaluated in `research/` and
reimplemented in Go if they survive. The tool ships as a single binary.

## License

Apache 2.0. See [LICENSE](LICENSE).
