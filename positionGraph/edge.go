package positionGraph

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/mat"
)

/*
edge.go
Description:

	Defines an edge for use in the position graph.
*/

// =======
// Objects
// =======

/*
PGEdge
*/
type PGEdge struct {
	graph *PositionGraph
	from  int64
	to    int64
}

// =======
// Methods
// =======

/*
From
Description:

	Returns the node that the edge is coming from.
*/
func (e *PGEdge) From() graph.Node {
	// Constants

	// Algorithm
	return e.graph.nodes[e.from]
}

/*
To
Description:

	Returns the node that the edge is going to.
*/
func (e *PGEdge) To() graph.Node {
	// Constants

	// Algorithm
	return e.graph.nodes[e.to]
}

/*
ReversedEdge
Description:

	Returns a new edge that is the reverse of the current edge.
*/
func (e *PGEdge) ReversedEdge() graph.Edge {
	// Constants

	// Algorithm
	return &PGEdge{
		graph: e.graph,
		from:  e.to,
		to:    e.from,
	}
}

/*
Weight
Description:

	Returns the weight of the edge.
*/
func (e *PGEdge) Weight() float64 {
	// Constants
	from := e.From().(*Node)
	to := e.To().(*Node)

	// Algorithm
	var distance *mat.VecDense = mat.NewVecDense(from.Position.Len(), nil)
	distance.SubVec(to.Position, to.Position)
	return mat.Norm(distance, 2)
}
