package position_graph_test

import (
	position_graph "github.com/GraphPathPlanning.go/graphs/position"
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
func CreateTestGraph_ForPositionGraph1() *position_graph.PositionGraph {
	// Constants
	g := position_graph.New()

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
TestPositionGraph_From2
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

/*
TestPositionGraph_From3
Description:

	Tests the From() method of the graph object when
	a node ID is provided, that does not exist.
*/
func TestPositionGraph_From3(t *testing.T) {
	// Setup
	g := CreateTestGraph_ForPositionGraph1()
	var badID int64 = 101

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic")
		}

		rAsError, tf := r.(position_graph.NodeNotFoundError)
		if !tf {
			t.Errorf(
				"the panic was caused by something that was not an error: %T",
				r,
			)
		}

		expectedError := position_graph.NodeNotFoundError{
			ID: badID,
		}
		if rAsError.Error() != expectedError.Error() {
			t.Errorf(
				"expected error \"%v\"; received \"%v\"",
				expectedError,
				rAsError,
			)
		}

	}()

	g.From(badID)

}

/*
TestPositionGraph_HasEdgeBetween1
Description:

	Tests that the function HasEdgeBetween() correctly denotes
	when an edge exists between two given nodes.
*/
func TestPositionGraph_HasEdgeBetween1(t *testing.T) {
	// Setup
	g := CreateTestGraph_ForPositionGraph1()
	var target0 int64 = 0
	var target1 int64 = 1

	// Check whether there exists an edge between target0 and target1
	if !g.HasEdgeBetween(target0, target1) {
		t.Errorf("there is an edge between %v and %v, but function says no!",
			target0, target1,
		)
	}

}

/*
TestPositionGraph_HasEdgeBetween2
Description:

	Tests that the function HasEdgeBetween() correctly denotes
	when an edge exists between two given nodes.
	In this case, we ask whether an edge exists between nodes
	1 and 0, when an edge was added to graph between 0 and 1.
	Because the graph is undirected, this function should still
	return true.
*/
func TestPositionGraph_HasEdgeBetween2(t *testing.T) {
	// Setup
	g := CreateTestGraph_ForPositionGraph1()
	var target0 int64 = 0
	var target1 int64 = 1

	// Check whether there exists an edge between target0 and target1
	if !g.HasEdgeBetween(target1, target0) {
		t.Errorf("there is an edge between %v and %v, but function says no!",
			target1, target0,
		)
	}

}

/*
TestPositionGraph_HasEdgeBetween3
Description:

	Tests that the HasEdgeBetween() method properly returns false,
	when no edge exists between two nodes.
*/
func TestPositionGraph_HasEdgeBetween3(t *testing.T) {
	// Setup
	g := CreateTestGraph_ForPositionGraph1()
	var target0 int64 = 0
	var target1 int64 = 4

	// Check whether there exists an edge between target0 and target1
	if g.HasEdgeBetween(target0, target1) {
		t.Errorf("there is NOT an edge between %v and %v, but function says there is!",
			target0, target1,
		)
	}
}

/*
TestPositionGraph_Weight1
Description:

	Tests that the Weight() method properly returns the weight associated
	with two nodes that exist in the graph.
*/
func TestPositionGraph_Weight1(t *testing.T) {
	// Setup
	g := CreateTestGraph_ForPositionGraph1()
	var target0 int64 = 0
	var target1 int64 = 1

	// Test
	wght, tf := g.Weight(target0, target1)
	if !tf {
		t.Errorf("edge exists, so tf should be true, but received false!")
	}

	if !mat.EqualApprox(
		mat.NewVecDense(1, []float64{1.0}),
		mat.NewVecDense(1, []float64{wght}),
		1.0e-5,
	) {
		t.Errorf("weight was %v; expected 1.0", wght)
	}
}

/*
TestPositionGraph_Weight2
Description:

	Tests that the function properly returns a false flag when
	asked to collect a Weight in the graph that doesn't exist.
*/
func TestPositionGraph_Weight2(t *testing.T) {
	// Setup
	g := CreateTestGraph_ForPositionGraph1()
	var target0 int64 = 0
	var target1 int64 = 4

	// Test
	_, tf := g.Weight(target0, target1)
	if tf {
		t.Errorf("expected weight function to return false flag for pair of nodes that ARE NOT connected;\n received true!")
	}
}

/*
TestPositionGraph_GetNodeAt1
Description:

	Tests whether we can retrieve a node by its position.
*/
func TestPositionGraph_GetNodeAt1(t *testing.T) {
	// Setup
	g := CreateTestGraph_ForPositionGraph1()
	x := mat.NewVecDense(2, []float64{1.0, 0.0})

	// Test
	n := g.GetNodeAt(x)

	if n.ID() != 1 {
		t.Errorf("Expected node %v to be at position %v; received node %v",
			1,
			x,
			n.ID(),
		)
	}
}

/*
TestPositionGraph_GetNodeAt2
Description:

	Tests whether GetNodeAt() returns nil when you try to retrieve a node
	from a position that is not reflected in the graph.
*/
func TestPositionGraph_GetNodeAt2(t *testing.T) {
	// Setup
	g := CreateTestGraph_ForPositionGraph1()
	x := mat.NewVecDense(2, []float64{101.0, 20.0})

	// Test
	n := g.GetNodeAt(x)

	if n != nil {
		t.Errorf(
			"expected node address to be nil; received %v",
			n,
		)
	}
}

/*
TestPositionGraph_RemoveEdge1
Description:

	Tests that when we remove an edge (the only edge between
	node 3 and node 4), the From() method later returns zero nodes.
*/
func TestPositionGraph_RemoveEdge1(t *testing.T) {
	// Setup
	g := CreateTestGraph_ForPositionGraph1()
	var id1, id2 int64 = 3, 4
	nFrom1 := g.From(id2).Len()
	if nFrom1 != 1 {
		t.Errorf(
			"expected for from to initially contain %v elements; received %v",
			1,
			nFrom1,
		)
	}

	// Test remove method
	g.RemoveEdge(id1, id2)

	// Check the nFrom again
	if g.From(id2).Len() != nFrom1-1 {
		t.Errorf(
			"expected there to be %v nodes connected to %v after edge removal; received %v",
			nFrom1-1,
			id2,
			g.From(id2).Len(),
		)
	}
}
