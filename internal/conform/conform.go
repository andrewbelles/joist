// Checking, ranking and the disagreement quadrant. Every function here is a pure
// function of its inputs. Anything that reads the clock, the network or ambient
// repo state belongs in a caller, not here.

package conform

import (
	"errors"

	"github.com/andrewbelles/call-graph-tooling/internal/boundary"
	"github.com/andrewbelles/call-graph-tooling/internal/graph"
	"github.com/andrewbelles/call-graph-tooling/schema"
)

// ErrNotImplemented marks the scaffold.
var ErrNotImplemented = errors.New("conform: not implemented")

// Severity orders violations for reporting. It does not decide whether CI fails;
// that is the caller's policy, driven by the low confidence flag and the ratchet.
type Severity int

const (
	SeverityInfo Severity = iota
	SeverityWarn
	SeverityError
)

// Violation is one edge that contradicts a declared rule.
type Violation struct {
	Rule     boundary.RuleKind
	From     string
	To       string
	Edge     graph.Edge
	Severity Severity
	Message  string
}

// Checker evaluates a resolved graph against declared rules.
type Checker interface {
	Check(g *graph.Graph, r *boundary.Resolver) ([]Violation, error)
}

// Rank orders violations by the cost of severing the offending edge, measured in
// raw call sites. Ranking is stable: equal costs keep symbol ID order so the
// artifact does not churn between runs.
func Rank(v []Violation) []Violation {
	return nil
}

// Cell is one entry of the disagreement quadrant. A pair with high co-change and
// no static edge is a hidden coupling finding.
type Cell struct {
	A        string
	B        string
	CoChange float64
	Static   bool
	Hidden   bool
}

// Disagree crosses structural edges with the co-change matrix. It takes both as
// plain inputs so this stays a pure function.
func Disagree(g *graph.Graph, coChange map[[2]string]float64) []Cell {
	return nil
}

// Result is everything one deterministic run produces, before it is written.
type Result struct {
	Graph      *graph.Graph
	Violations []Violation
	Cells      []Cell
}

// Artifact converts a Result into the published format. This is the only place
// the internal model crosses into the schema module, so a change to either side
// shows up here.
func Artifact(r *Result, commitSHA string, gen schema.Generator) (*schema.Artifact, error) {
	return nil, ErrNotImplemented
}
