# testdata

Shared test corpus. Two consumers read this directory: Go regression and
integration tests, and the offline harness in `research/`. A single source of
truth for which repositories to analyze keeps the two from drifting apart and
keeps research conclusions traceable to a specific input.

Nothing here vendors third-party source. The corpus is a set of pointers.

## Tiers

Three tiers, separated by what they cost to run.

**Synthetic fixtures** live in `fixtures/`. Small repositories with a hand-known
graph, checked in whole, alongside their `arch.yaml` and their expected canonical
artifact bytes. These run in the default `make test` and are the only tier that
gates CI.

**Pinned build corpus** is listed in `corpus.json` with `tier: "build"`. Real
repositories small enough to clone and compile in a few minutes. Go integration
tests build these and assert properties of the resulting graph. They run on
demand, not in the default test target.

**Analysis corpus** is the rest of `corpus.json`. Large real projects used for
metric evaluation. Too expensive to build per test run, so they are built rarely
and their exported artifacts are what most consumers actually read.

Fixtures scoped to one package belong in that package's own `testdata/`
directory, per Go convention. This directory is only for cases that cross
package boundaries or leave the repository.

## corpus.json

JSON rather than YAML so both consumers can parse it without a dependency. Go
reads JSON in the standard library, and `schema/` is deliberately zero-dep.

```json
{
  "version": 1,
  "repos": [
    {
      "id": "example-service",
      "url": "https://github.com/owner/name",
      "sha": "0000000000000000000000000000000000000000",
      "tier": "build",
      "languages": ["java"],
      "build": "maven",
      "setup": "recipes/example-service.sh",
      "quality": { "basis": "defect-density", "label": "high", "source": "..." },
      "notes": ""
    }
  ]
}
```

`sha` is a full commit SHA, never a branch or a tag. Tags get moved and branches
advance, and either one makes a golden drift with no local change to blame. An
unpinned corpus reintroduces exactly the nondeterminism the canonical encoding
exists to prevent.

`setup` points at a build recipe. Build-level analysis needs the target to
actually compile, which means a per-repo toolchain and dependency resolution
step. That recipe is the real maintenance cost of this corpus and it rots
independently of anything in this repository. Add a repo to the build tier only
when someone is willing to keep its recipe working.

## Division with research/

The corpus is shared at the pointer level, not the artifact level.

Go builds the repositories and produces artifacts. `research/` reads exported
artifacts and never builds anything. If the harness grew a build orchestrator it
would need the same per-language toolchains as the product, and the rule that
Python is never a runtime dependency would stop meaning much.

Practically: run the analyzer over a corpus entry, export the artifact, and point
the harness at that. The harness treats artifacts as read-only input.

## Quality labels

`quality.basis` is required and names the external evidence for the label.
Intuition is not a basis.

If repositories are labeled good or bad by feel and the metrics then separate
good from bad, that result is circular and says nothing. Ground labels in
something outside the graph: measured defect density, a documented rewrite or
architecture migration, published architecture documentation that conformance can
be checked against. Where no such evidence exists, record the label as a
hypothesis and say so in `notes`.

## Goldens

Golden artifact bytes are compared exactly, not field by field. The canonical
encoding is what the cache addresses, so a difference in serialization is a real
failure even when the decoded values are equal.

Regenerate goldens explicitly. Never regenerate them automatically on failure.
