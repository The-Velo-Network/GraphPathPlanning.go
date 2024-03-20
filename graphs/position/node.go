package position_graph

import "gonum.org/v1/gonum/mat"

/*
node.go
Description:

	Defines a node for use in the position graph.
*/

// =======
// Objects
// =======

/*
Node
Description:

	An implementation of a Node for use in the position graph.
*/
type Node struct {
	id       int64
	Position *mat.VecDense
}

// =======
// Methods
// =======

/*
ID
Description:

	Returns the ID of the node.
*/
func (n *Node) ID() int64 {
	// Constants

	// Algorithm
	return n.id
}
