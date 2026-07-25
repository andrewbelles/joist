// Artifact types. One artifact describes one commit SHA and is immutable once
// written. Every field here is part of the published format, so adding or
// changing one requires a Version bump.

package schema

// Version is the artifact format version. Consumers reject an artifact whose
// Version they do not recognize rather than guessing at missing fields.
const Version = 1

// Artifact is the complete output of a deterministic run. Nothing in it comes
// from the exploratory path.
type Artifact struct {
	Version       int          `json:"version"`
	CommitSHA     string       `json:"commit_sha"`
	Generator     Generator    `json:"generator"`
	LowConfidence bool         `json:"low_confidence"`
	Symbols       []Symbol     `json:"symbols"`
	Edges         []Edge       `json:"edges"`
	Boundaries    []Boundary   `json:"boundaries"`
	Violations    []Violation  `json:"violations"`
	CoChange      []CoChange   `json:"co_change"`
	History       HistoryStats `json:"history"`
}

// Generator identifies what produced an artifact. It participates in cache keys,
// so a change to any field here must invalidate previously cached results.
type Generator struct {
	Tool      string            `json:"tool"`
	Version   string            `json:"version"`
	Indexers  map[string]string `json:"indexers"`
	Toolchain map[string]string `json:"toolchain"`
}

// Symbol is one addressable program element. ID is a stable SCIP symbol string
// and is the only identity used across artifacts.
type Symbol struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Kind   string `json:"kind"`
	File   string `json:"file"`
	Line   int    `json:"line"`
	Module string `json:"module"`
}

// EdgeKind separates edges recovered from compiled output from edges synthesized
// by a framework adapter. Consumers weight the two differently.
type EdgeKind string

const (
	EdgeCall       EdgeKind = "call"
	EdgeImport     EdgeKind = "import"
	EdgeFramework  EdgeKind = "framework"
	EdgeStructural EdgeKind = "structural"
)

// Edge is a directed dependency between two symbols. Sites is the raw call site
// count and is the cost of severing the edge.
type Edge struct {
	From   string   `json:"from"`
	To     string   `json:"to"`
	Kind   EdgeKind `json:"kind"`
	Sites  int      `json:"sites"`
	Source string   `json:"source"`
}

// Boundary is a module after glob resolution. Members holds symbol IDs, sorted.
type Boundary struct {
	Name    string   `json:"name"`
	Globs   []string `json:"globs"`
	Members []string `json:"members"`
}

// RuleKind is the declared relationship a rule asserts between boundaries.
type RuleKind string

const (
	RuleAllow       RuleKind = "allow"
	RuleDeny        RuleKind = "deny"
	RuleOnlyThrough RuleKind = "only_through"
	RuleLayer       RuleKind = "layer"
)

// Violation is one edge that contradicts a declared rule. Rank is assigned by
// the ranker and orders the list; lower is more severe.
type Violation struct {
	Rule     RuleKind `json:"rule"`
	From     string   `json:"from"`
	To       string   `json:"to"`
	Edge     Edge     `json:"edge"`
	Rank     int      `json:"rank"`
	Severity string   `json:"severity"`
	Message  string   `json:"message"`
}

// CoChange is one pair of files that change together in history, normalized
// against an independent change null model. Score above zero means the pair
// changes together more than chance.
type CoChange struct {
	A      string  `json:"a"`
	B      string  `json:"b"`
	Count  int     `json:"count"`
	Score  float64 `json:"score"`
	Hidden bool    `json:"hidden"`
}

// HistoryStats describes the history window an artifact was mined from. It is
// recorded so a consumer can tell a weak signal from a short window.
type HistoryStats struct {
	Commits   int    `json:"commits"`
	FirstSHA  string `json:"first_sha"`
	LastSHA   string `json:"last_sha"`
	HalfLife  string `json:"half_life"`
	Truncated bool   `json:"truncated"`
}
