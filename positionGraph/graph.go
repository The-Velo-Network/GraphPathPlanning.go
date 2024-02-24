package positionGraph

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/iterator"
	"gonum.org/v1/gonum/mat"
)

/*
graph.go
Description:

	Defines a graph for use in the position graph.
*/

// =======
// Objects
// =======

/*
PositionGraph
*/
type PositionGraph struct {
	nodes map[int64]*Node
	edges map[int64]*PGEdge
}

// =======
// Methods
// =======

/*
NewPositionGraph
Description:

	Creates a new PositionGraph.
*/
func NewPositionGraph() *PositionGraph {
	// Constants

	// Algorithm
	return &PositionGraph{
		nodes: make(map[int64]*Node),
		edges: make(map[int64]*PGEdge),
	}
}

/*
Node
Description:

	Returns the node with the given ID.
*/
func (pg *PositionGraph) Node(id int64) graph.Node {
	// Constants

	// Algorithm
	out, ok := pg.nodes[id]
	if !ok {
		panic(NodeNotFoundError{id})
	}

	return out
}

/*
Nodes
Description:

	Returns the nodes in the graph.
*/
func (pg *PositionGraph) Nodes() graph.Nodes {
	// Constants

	// Algorithm
	var out []graph.Node
	for _, n := range pg.nodes {
		out = append(out, n)
	}

	return iterator.NewOrderedNodes(out)
}

/*
From
Description:

	Returns the nodes that can be reached from the node
	with the given ID.
*/
func (pg *PositionGraph) From(id int64) graph.Nodes {
	// Input Processing
	if _, ok := pg.nodes[id]; !ok {
		panic(NodeNotFoundError{id})
	}

	// Constants

	// Algorithm
	var out []graph.Node
	for _, e := range pg.edges {
		if e.From().ID() == id {
			out = append(out, e.To())
		}
	}

	return iterator.NewOrderedNodes(out)
}

/*
HasEdgeBetween
Description:

	Returns whether or not there is an edge between the
	two nodes with the given IDs.
*/
func (pg *PositionGraph) HasEdgeBetween(from, to int64) bool {
	// Constants

	// Algorithm
	for _, e := range pg.edges {
		if e.From().ID() == from && e.To().ID() == to {
			return true
		}
	}

	return false
}

/*
Edge
Description:

	Returns the edge between the two nodes with the given IDs.
*/
func (pg *PositionGraph) Edge(from, to int64) graph.Edge {
	// Constants

	// Algorithm
	for _, e := range pg.edges {
		if e.From().ID() == from && e.To().ID() == to {
			return e
		}
	}

	return nil
}

/*
WeightedEdge
Description:

	Returns the weighted edge between the two nodes with the given IDs.
*/
func (pg *PositionGraph) WeightedEdge(from, to int64) graph.WeightedEdge {
	return pg.Edge(from, to).(*PGEdge)
}

/*
Weight
Description:

	Returns the weight of the edge between the two nodes with the given IDs.

From the gonum documentation:

	// Weight returns the weight for the edge between
	// x and y with IDs xid and yid if Edge(xid, yid)
	// returns a non-nil Edge.
	// If x and y are the same node or there is no
	// joining edge between the two nodes the weight
	// value returned is implementation dependent.
	// Weight returns true if an edge exists between
	// x and y or if x and y have the same ID, false
	// otherwise.
*/
func (pg *PositionGraph) Weight(xid, yid int64) (float64, bool) {
	// Collect edge
	tempEdge := pg.WeightedEdge(xid, yid)

	// Return weight, if possible
	if tempEdge == nil {
		// Note: This will also make any self-loops have this super high weight.
		return 1e10, false
	}

	return tempEdge.Weight(), true

}

/*
EdgeBetween
Description:

	Finds the edge between the given two ids, if it exists.
	Otherwise, returns nil.
*/
func (pg *PositionGraph) EdgeBetween(from, to int64) graph.Edge {
	// Collect edge from xid to yid
	edge1 := pg.Edge(from, to)
	if edge1 != nil {
		return edge1
	}

	// If that edge doesn't exist, then try the reversed version
	return pg.Edge(to, from)
}

/*
WeightedEdgeBetween
Description:

	Finds the weighted edge between the two ids, if it
	exists. Otherwise, returns nil.
*/
func (pg *PositionGraph) WeightedEdgeBetween(from, to int64) graph.WeightedEdge {
	// Use Edge Between and cast it to the correct type
	return pg.EdgeBetween(from, to).(*PGEdge)
}

/*
AddNode
Description:

	Adds a node to the graph.
*/
func (pg *PositionGraph) AddNode(n Node) {
	// Constants

	// Algorithm
	pg.nodes[n.ID()] = &n
}

/*
AddNoteAt
Description:

	Adds a node to the graph at a specific position.
*/
func (pg *PositionGraph) AddNodeAt(position *mat.VecDense) Node {
	// Constants

	// Create node
	var nextIndex int64 = int64(len(pg.nodes))
	n := Node{
		id:       nextIndex,
		Position: position,
	}

	// Add node
	pg.AddNode(n)

	// Return node
	return n
}

/*
AddEdge
Description:

	Adds an edge to the graph.
*/
func (pg *PositionGraph) AddEdge(e PGEdge) {
	// Constants

	// Algorithm
	var nextIndex int64 = int64(len(pg.edges))
	pg.edges[nextIndex] = &e
}

/*
AddEdgeBetween
Description:

	Adds an edge between two nodes in the graph.
*/
func (pg *PositionGraph) AddEdgeBetween(from Node, to Node) PGEdge {
	// Constants

	// Create edge
	e := PGEdge{
		graph: pg,
		from:  from.ID(),
		to:    to.ID(),
	}

	// Add edge
	pg.AddEdge(e)

	// Return edge
	return e
}
