package positionGraph_test

import (
	"github.com/GraphPathPlanning.go/positionGraph"
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
	pg := positionGraph.NewPositionGraph()
	n := pg.AddNodeAt(
		mat.NewVecDense(3, []float64{1.0, 2.0, 3.0}),
	)

	// Algorithm
	if n.ID() != 0 {
		t.Errorf("Expected 1, got %v", n.ID())
	}
}
