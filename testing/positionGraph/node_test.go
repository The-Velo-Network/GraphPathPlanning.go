package positionGraph_test

import (
	position_graph "github.com/GraphPathPlanning.go/graphs/position"
	"gonum.org/v1/gonum/mat"
	"testing"
)

/*
node_test.go
Description:

	Defines a node for use in the position graph.
*/

/*
TestNode_ID1
Description:

	Tests the ID method of the Node object.
*/
func TestNode_ID1(t *testing.T) {
	// Constants
	pg := position_graph.New()
	n := pg.AddNodeAt(
		mat.NewVecDense(3, []float64{1.0, 2.0, 3.0}),
	)

	// Algorithm
	if n.ID() != 0 {
		t.Errorf("Expected 1, got %v", n.ID())
	}
}
