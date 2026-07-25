// The declared architecture. arch.yaml is hand authored and schema versioned,
// so parsing rejects unknown fields rather than ignoring them; a typo in a rule
// must fail loudly instead of silently disabling the rule.

package boundary

import "errors"

// ErrNotImplemented marks the scaffold.
var ErrNotImplemented = errors.New("boundary: not implemented")

// RuleKind is the declared relationship a rule asserts between modules.
type RuleKind string

const (
	// RuleAllow permits From to depend on To.
	RuleAllow RuleKind = "allow"
	// RuleDeny forbids From from depending on To.
	RuleDeny RuleKind = "deny"
	// RuleOnlyThrough permits the dependency only via the modules in Through.
	RuleOnlyThrough RuleKind = "only_through"
	// RuleLayer forbids any dependency from a lower layer to a higher one.
	RuleLayer RuleKind = "layer"
)

// Module is a named set of globs. Globs are repo relative.
type Module struct {
	Name  string
	Globs []string
}

// Rule is one declared constraint between modules.
type Rule struct {
	Kind    RuleKind
	From    string
	To      string
	Through []string
	Reason  string
}

// Config is a parsed arch.yaml.
type Config struct {
	SchemaVersion int
	Modules       []Module
	Rules         []Rule
	Layers        []string
}

// Parse reads an arch.yaml. It rejects unknown fields and unknown rule kinds.
func Parse(b []byte) (*Config, error) {
	return nil, ErrNotImplemented
}

// Resolver maps symbols to module membership after glob expansion. Build it once
// per run; it is read only afterwards.
type Resolver struct {
	cfg *Config
}

// NewResolver expands cfg against the given repo relative file set. Overlapping
// globs are a config error, not a precedence question.
func NewResolver(cfg *Config, files []string) (*Resolver, error) {
	return nil, ErrNotImplemented
}

// Module returns the module owning the given repo relative file.
func (r *Resolver) Module(file string) (string, bool) {
	return "", false
}
