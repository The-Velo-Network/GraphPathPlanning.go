package positionGraph_test

import (
	"github.com/GraphPathPlanning.go/positionGraph"
	"gonum.org/v1/gonum/mat"
	"testing"
)

/*
edge_test.go
Description:

	Tests the PositionGraph object.
*/

/*
CreateTestGraph_ForPositionGraph1
Description:

	Creates a simple test graph with six nodes and four edges
*/
func CreateTestGraph_ForPositionGraph1() *positionGraph.PositionGraph {
	// Constants
	g := positionGraph.NewPositionGraph()

	// Algorithm

	n1 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{0.0, 0.0}),
	)
	n2 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{1.0, 0.0}),
	)
	n3 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{0.0, 1.0}),
	)
	n4 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{1.0, 1.0}),
	)
	n5 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{2.0, 0.0}),
	)
	g.AddNodeAt(
		mat.NewVecDense(2, []float64{2.0, 1.0}),
	)
	g.AddEdgeBetween(n1, n2)
	g.AddEdgeBetween(n1, n3)
	g.AddEdgeBetween(n2, n4)
	g.AddEdgeBetween(n4, n5)
	//g.AddEdgeBetween(n4, n6)

	return g
}

/*
TestPositionGraph_Node1
Description:

	Tests that the Node method properly panics
	if asked for a node that does not exist.
*/
func TestPositionGraph_Node1(t *testing.T) {
	// Setup
	g := CreateTestGraph_ForPositionGraph1()

	// Algorithm
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	g.Node(100)
}

/*
TestPositionGraph_From1
Description:

	Tests the From method of the graph object.
	This method should produce ALL nodes that are coming from the
	node with ID 0. There should be two such nodes.
*/
func TestPositionGraph_From1(t *testing.T) {
	// Constants
	g := CreateTestGraph_ForPositionGraph1()

	// Algorithm
	from := g.From(0)

	if from.Len() != 2 {
		t.Errorf("Expected 2, got %v", from.Len())
	}
}

/*
TestPositionGraph_From1
Description:

	Tests the From method of the graph object.
	This method will try to find all nodes that are coming
	from node 5 which should be 0.
*/
func TestPositionGraph_From2(t *testing.T) {
	// Setup
	g := CreateTestGraph_ForPositionGraph1()

	// Algorithm
	from := g.From(5)

	if from.Len() != 0 {
		t.Errorf("Expected 0, got %v", from.Len())
	}
}
