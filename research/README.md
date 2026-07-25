# research

Python harness for evaluating graph metrics, primarily NetworKit for community
detection and centrality, while deciding which metrics are worth keeping.

This is not part of the product. It is never built by the Makefile default
target, never imported by Go, and never installed alongside the binary. The
tool ships as one Go binary; the moment Python becomes a runtime dependency that
story is gone.

Metrics that survive evaluation here get reimplemented in `internal/explore`,
which is itself advisory only and never on the CI path.

## Input

The harness reads exported artifacts. It does not clone or build anything; the Go
side does that and exports the result. Building the corpus needs a per-language
toolchain for every target repository, and duplicating that here would make
Python a real dependency of the analysis path.

Which repositories the corpus covers is recorded in `testdata/corpus.json`, shared
with the Go integration tests so both sides evaluate the same inputs. See
[../testdata/README.md](../testdata/README.md).

## Setup

```sh
cd research
python -m venv .venv && . .venv/bin/activate
pip install -e .
```
