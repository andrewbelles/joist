# Roadmap

Author: Claude Opus 5 

Current state: contracts and repo shape only. No analysis is implemented.

D0 is decisions. S1 through S9 are implementation. Each stage names how it is
verified, so a stage is either done or not.

## D0: decide before writing implementation

### D0.0 Spike the indexers

Run `scip-typescript`, `scip-java` and `scip-python` by hand on five candidate
repos and read the output. No joist code. Everything below depends on what this
shows.

Two things to establish:

SCIP is an occurrence index, not a call graph. It records that a symbol is
referenced at a position and that a definition spans a range. A call edge is
derived by finding which definition range encloses each reference. Confirm that
derivation works on real code before committing to a graph model. Edge cases are
nested closures, class field initializers, top level statements, generated code.

Raw index size on a large repo. That number decides D0.1.

### D0.1 Artifact against graph store

The design has one artifact keyed by commit SHA and an embedded graph store. Their
relationship is undefined. If the artifact holds the full symbol graph it may be
too large on a real repo. If it holds a summary, then blast radius queries need
the store and the claim that every consumer reads one file no longer holds.

Likely answer: the artifact holds boundaries, violations, co-change and
diagnostics, plus the content hash of a separately addressed graph blob. Decide
before S1.

### D0.2 Node granularity and the symbol to file join

The graph is symbol level. Git history is path level. Co-change, the disagreement
quadrant and the DSM all need the two reconciled. Decide which level violations
and the DSM operate at, and how symbol edges aggregate upward. Blocks S6 and the
viewer.

### D0.3 Canonical encoding

JSON with sorted keys is debuggable and keeps `schema/` dependency free. Protobuf
is compact but adds a dependency to the module that is deliberately clean. Decide
before writing `schema.Encode`.

### D0.4 What a document is

The manifest maps document to blob hash, so this sets incremental granularity.
Source file is the obvious answer. Confirm it holds for languages where a
compilation unit spans files.

### D0.5 arch.yaml semantics

Rule precedence when several match. Default posture, deny by default or allow by
default. Whether overlapping globs are an error, which `boundary/config.go`
currently asserts and may be too strict.

### D0.6 Fidelity model

Replace the binary `LowConfidence` flag with a per-indexer tier. `scip-python` and
`scip-java` do not deserve equal trust and reporting them identically misleads.

### D0.7 First language

One language to excel at, four as baseline. TypeScript keeps corpus cost low since
it needs dependency install but no compile. Java suits the architecture governance
buyer at higher build cost.

## S1: SCIP to graph to artifact

Ingest one indexer, derive call edges, build `graph.Graph`, serialize canonically.
No cache, no rules, no history.

Verify: two runs on the same pinned commit produce byte identical artifacts, on
two different machines. Golden test on a synthetic fixture.

**Live testing starts here.** From S1 the tool runs against real open source
repositories, and every later stage is testable against the same corpus.

## S2: viewer on a real artifact

DSM and ranked list rendered from S1 output.

Verify: `joist viewer` renders a real repo. Eyeballing whether the structure is
plausible is the fastest correctness signal available at this point.

## S3: MCP server

who-calls and blast-radius-of as tools, the artifact as a resource.

Verify: point an agent at it and ask real questions about a real repo. First
honest test of the core thesis.

## S4: boundary and conform

arch.yaml parsing, glob resolution, rule evaluation, violation ranking. `joist
check` starts doing work.

Verify: hand author an arch.yaml for a corpus repo. Known violations appear and
known good edges do not.

## S5: cache, manifest, incremental rebuild

Content addressing, per-SHA manifests, manifest diffing, rebuild tiers.

Verify: an incremental artifact is byte identical to a full rebuild at the same
SHA. That equality is the entire correctness argument for the cache. Also record
full against incremental rebuild time on a large repo.

## S6: history and disagreement

Git mining, co-change against an independent change null model, time decay, rename
tracking, the disagreement quadrant.

Verify: known hidden couplings in a repo you understand well surface as findings.

## S7: first framework adapter

One adapter in the chosen ecosystem, out of process, emitting SCIP symbol strings.

Verify: synthesized edges join the precise graph rather than producing orphans.

## S8: remaining languages

The other four. Mostly invocation and build recipes by this point.

Verify: each language runs against its corpus tier with a recorded fidelity level.

## S9: explore and the research loop

Feed artifacts to `research/`, evaluate metrics, promote survivors into
`internal/explore`.

Verify: a metric moves from Python to Go only after it separates known good from
known bad repos on a basis recorded in `testdata/corpus.json`.

## Sequencing note

S5 is late, but D0.1 and D0.4 are not. Incrementality cannot be retrofitted into a
model that cannot take it. Build the vertical slice first, decide the cache design
up front.
