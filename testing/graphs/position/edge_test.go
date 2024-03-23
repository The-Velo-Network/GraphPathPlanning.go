package position_graph_test

/*
edge_test.go
Description:

	Tests the edge object for the position graph.
*/

import (
	positionGraph2 "github.com/GraphPathPlanning.go/graphs/position"
	"gonum.org/v1/gonum/mat"
	"testing"
)

/*
CreateTestGraph_ForEdges1
Description:

	Creates a simple test graph with four nodes and three edges
	for use in these tests.
*/
func CreateTestGraph_ForEdges1() *positionGraph2.PositionGraph {
	// Constants
	g := positionGraph2.New()

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
	g.AddEdgeBetween(n1, n2)
	g.AddEdgeBetween(n1, n3)
	g.AddEdgeBetween(n2, n4)

	return g
}

/*
TestEdge_From1
Description:

	Tests the From method of the edge object.
	This method should return the node that the edge is coming from
	which should be the node with ID 0.
*/
func TestEdge_From1(t *testing.T) {
	// Constants
	g := CreateTestGraph_ForEdges1()
	edge := g.EdgeBetween(0, 1)

	// Algorithm
	if edge.From().ID() != 0 {
		t.Errorf("Expected 0, got %v", edge.From().ID())
	}
}

/*
TestEdge_To1
Description:

	Tests the To method of the edge object.
	This method should return the node that the edge is going to
	which should be the node with ID 1.
*/
func TestEdge_To1(t *testing.T) {
	// Constants
	g := CreateTestGraph_ForEdges1()
	edge := g.EdgeBetween(0, 1)

	// Algorithm
	if edge.To().ID() != 1 {
		t.Errorf("Expected 1, got %v", edge.To().ID())
	}
}

/*
TestEdge_ReversedEdge1
Description:

	Tests the ReversedEdge method of the edge object.
	Checks that the new edge has the correct from (1) and to (0) nodes.
*/
func TestEdge_ReversedEdge1(t *testing.T) {
	// Constants
	g := CreateTestGraph_ForEdges1()
	edge := g.EdgeBetween(0, 1)

	// Algorithm
	reversedEdge := edge.ReversedEdge()
	if reversedEdge.From().ID() != 1 {
		t.Errorf("Expected 1, got %v", reversedEdge.From().ID())
	}
	if reversedEdge.To().ID() != 0 {
		t.Errorf("Expected 0, got %v", reversedEdge.To().ID())
	}

}

/*
TestEdge_Weight1
Description:

	Tests the Weight method of the edge object.
	For the edge between 0 and 1, the weight should be 1.0.
*/
func TestEdge_Weight1(t *testing.T) {
	// Constants
	g := CreateTestGraph_ForEdges1()
	edge := g.EdgeBetween(0, 1).(*positionGraph2.PGEdge)

	// Algorithm
	if edge.Weight() != 1.0 {
		t.Errorf("Expected 1.0, got %v", edge.Weight())
	}
}
