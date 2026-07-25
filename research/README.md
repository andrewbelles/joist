# research

Python harness for evaluating graph metrics, primarily NetworKit for community
detection and centrality, while deciding which metrics are worth keeping.

This is not part of the product. It is never built by the Makefile default
target, never imported by Go, and never installed alongside the binary. The
tool ships as one Go binary; the moment Python becomes a runtime dependency that
story is gone.

Metrics that survive evaluation here get reimplemented in `internal/explore`,
which is itself advisory only and never on the CI path.

## Setup

```sh
cd research
python -m venv .venv && . .venv/bin/activate
pip install -e .
```
