# CLAUDE.md

## What this is

A build-level call graph analyzer that checks a codebase against a declared
architecture, rebuilds incrementally per commit SHA, and publishes one
content-addressed artifact that the CLI, the CI check, the MCP server, and the
viewer all read. See [README.md](README.md) for the repo shape.

## Comment style

These rules apply to every comment and docstring in the repository, in every
language. They are not stylistic preferences; inconsistent commentary is what
makes a codebase unreadable at scale.

- Terse and coherent. Complete sentences, no filler, no restating the code.
- Say why, not what. If the code is unclear, fix the code.
- No git history and no versioning in comments. The exception is a warning that
  explains why something has to be the way it is, when the obvious alternative is
  wrong and someone will otherwise try it.
- No markdown, no RST, no LaTeX, no math symbols, no em dashes, no arrows, no
  decorative punctuation. Plain prose.
- No multi-line comment blocks anywhere except the file banner. In Go that means
  no `/* */`, ever.
- No mid-file banners or separators. No `// ---- helpers ----`, no section
  headers, no ASCII rules. If a file needs sections, it needs splitting.
- No commented-out code. Delete it; git has it.
- No TODO without a concrete condition that resolves it.

Bad:

```go
/* ============ Cache Key Helpers ============ */

// Compute computes the key.
// TODO: fix this later
// Added 2026-07-14 by AB, see PR #12 -- refactored from cache.go v2
func Compute(in Inputs) (Key, error) {
```

Good:

```go
// Compute returns the key for in. It is a pure function; identical Inputs must
// produce an identical Key on any machine at any time.
func Compute(in Inputs) (Key, error) {
```

## File banners

Every Go file opens with a banner: two to four `//` lines giving the
responsibility, then the prohibition or the constraint that is easy to violate.
The prohibition is the useful half.

In `doc.go` the banner is the package doc comment and sits directly above the
`package` clause. In every other file the banner is followed by a blank line
before the `package` clause, otherwise godoc concatenates them into one package
comment.

```go
// Package cache is the content addressed blob store and the per-SHA manifest.
// A key covers file content, compile flags, toolchain version and indexer
// version, so an entry is valid on any machine that computes the same key.
// A key must fully determine its value. Reading anything undeclared here turns
// every cache hit into a guess.
package cache
```

## Git workflow

Solo development, pragmatic and safe.

`main` is a first class object. It must always build and always pass CI. Changes
reach it through a `dev/<topic>` or `feat/<topic>` branch and a pull request,
never through a direct commit.

### Claude and git

Never, because these write history or touch the remote: `git commit`,
`git push`, `git merge`, `git rebase`, `git cherry-pick`, `git tag`,
`gh pr create`.

This applies to `dev/` and feature branches exactly as it applies to `main`. Being
asked in conversation does not grant permission, and neither does a stated reason
why a particular change is safe. Relaxing it means editing this file.

Never, because these discard uncommitted work: `git reset --hard`,
`git checkout -- <path>`, `git restore`, `git clean`, `git stash drop`,
`git stash clear`.

Fine, because these move refs or park work without losing any: `git branch`,
`git checkout <branch>`, `git switch`, `git stash push`, `git stash pop`, and
every read-only command such as `status`, `log`, `diff`, `show`, `blame`.

Note that `checkout` falls on both sides. Switching branches is safe and git
refuses it when it would overwrite uncommitted changes. Checking out a path
discards those changes with no warning.

Finish the work, leave it uncommitted in the working tree, and say plainly what
changed and what was verified. Andrew stages, commits, and pushes.

Commit messages describe the change, not the process. Write one on request as
text for Andrew to use; do not run it.

## Invariants

These are expensive to recover once violated, so check them before writing code
that crosses a package boundary.

- `internal/conform` must never import `internal/explore`. The deterministic path
  gates CI and has to be reproducible bit for bit; advisory signal entering it
  makes the gate unstable. Everything in `explore` is seed dependent, advisory,
  and never keyed into the artifact.
- The `schema` module takes no third party dependency, permanently. It is a
  separate Go module so CI tooling can read artifacts without the analyzer.
- Analysis reads nothing undeclared. No network, no wall clock, no ambient repo
  state, no absolute paths. An input that is not in the cache key will be served
  stale and nothing downstream will catch it.
- Artifact serialization is canonical: sorted keys, stable ordering, repo
  relative paths, no timestamps. Identical analyses must produce identical bytes
  or the cache silently stops hitting.
- Python is never a runtime dependency. `research/` is offline evaluation only.
  Metrics that survive get reimplemented in Go.
- The viewer and the MCP server are read only and bind loopback only. Nothing
  leaves the machine.

## Commands

| Command | Effect |
| --- | --- |
| `make dev` | Go only build, no Node needed, binary at `bin/joist` |
| `make build` | Web build, then Go with `-tags spa`, full binary |
| `make web` | Vite build into `internal/viewer/dist` |
| `make test` | `go test` in both modules |
| `make lint` | `go vet` in both modules |
| `make fmt` / `make fmt-check` | Format, or fail when formatting is needed |
| `make tidy` | `go mod tidy` in both modules |
| `make clean` | Remove `bin`, `dist`, `node_modules` |

There are two Go modules. Root commands do not reach into `schema/`; run them
there separately, which is what the Makefile targets already do.
