// The in memory graph. Adjacency is kept sorted at all times because traversal
// order determines serialization order, which determines the artifact hash.

package graph

// NodeID is a stable SCIP symbol string. It is the only identity used anywhere;
// file paths and line numbers are attributes, never keys.
type NodeID string

// Kind separates edges recovered from compiled output from edges synthesized by
// a framework adapter. Framework edges exist in no source text.
type Kind int

const (
	KindCall Kind = iota
	KindImport
	KindFramework
	KindStructural
)

// Node is one addressable program element.
type Node struct {
	ID     NodeID
	Name   string
	Kind   string
	File   string
	Line   int
	Module string
}

// Edge is a directed dependency. Sites is the raw call site count and stands in
// for the cost of severing the edge.
type Edge struct {
	From  NodeID
	To    NodeID
	Kind  Kind
	Sites int
}

// Graph holds nodes and their outgoing edges. Construct it with New and mutate
// it only through AddNode and AddEdge so ordering invariants hold.
type Graph struct {
	nodes map[NodeID]Node
	out   map[NodeID][]Edge
}

// New returns an empty graph.
func New() *Graph {
	return &Graph{nodes: make(map[NodeID]Node), out: make(map[NodeID][]Edge)}
}

// AddNode inserts or replaces a node.
func (g *Graph) AddNode(n Node) {
}

// AddEdge inserts an edge, keeping the adjacency slice sorted. Adding an edge
// whose endpoints are not yet present is allowed; nodes may arrive later.
func (g *Graph) AddEdge(e Edge) {
}

// Node returns the node for id.
func (g *Graph) Node(id NodeID) (Node, bool) {
	n, ok := g.nodes[id]
	return n, ok
}

// Out returns the outgoing edges of id in sorted order. The result aliases
// internal storage and must not be modified.
func (g *Graph) Out(id NodeID) []Edge {
	return g.out[id]
}

// Nodes returns every node ID in sorted order.
func (g *Graph) Nodes() []NodeID {
	return nil
}

// Weight scores an edge for ranking and layout. Framework edges carry real
// coupling but no call sites, so they cannot be weighted by Sites alone.
func Weight(e Edge) float64 {
	return 0
}
