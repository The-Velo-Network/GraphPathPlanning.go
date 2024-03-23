package djikstra_test

import (
	"github.com/GraphPathPlanning.go/gppErrors"
	position_graph "github.com/GraphPathPlanning.go/graphs/position"
	"github.com/GraphPathPlanning.go/planning/djikstra"
	"gonum.org/v1/gonum/mat"
	"testing"
)

/*
TestPlan_FindPlan1
Description:

	This test is meant to verify that the FindPlan()
	method finds a correct plan for a simple graph
	(with 2 nodes and one edge between them).
*/
func TestPlan_FindPlan1(t *testing.T) {
	// Setup Graph
	g := position_graph.New()

	n1 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{1.0, 2.0}),
	)
	n2 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{2.0, 2.0}),
	)

	g.AddEdgeBetween(n1, n2)

	// Apply Plan method
	p1, err := djikstra.FindPlan(g, n1.ID(), n2.ID())
	if err != nil {
		t.Errorf("there was a problem finding the plan: %v", err)
	}

	if len(p1.Sequence) != 2 {
		t.Errorf("there should only be two nodes in the plan, found %v", len(p1.Sequence))
	}

	for idx := int64(0); int(idx) < len(p1.Sequence); idx++ {
		if p1.Sequence[idx].ID() != idx {
			t.Errorf(
				"expected first node in plan to be node %v; received node %v",
				idx,
				p1.Sequence[idx].ID(),
			)
		}
	}
}

/*
TestPlan_FindPlan2
Description:

	This test is meant to verify that the FindPlan()
	method properly returns an error when no plan exists.
*/
func TestPlan_FindPlan2(t *testing.T) {
	// Setup Graph
	g := position_graph.New()

	n1 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{1.0, 2.0}),
	)
	n2 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{2.0, 2.0}),
	)

	// Apply Plan method
	_, err := djikstra.FindPlan(g, n1.ID(), n2.ID())
	if err == nil {
		t.Errorf("no error was thrown, but one should have been!")
	} else {
		expectedError := gppErrors.NoPathFound{
			Graph: g,
		}
		if err.Error() != expectedError.Error() {
			t.Errorf(
				"expected error \"%v\"; received \"%v\"",
				expectedError,
				err,
			)
		}
	}
}

/*
CreateREADMEGraph
Description:

	Creates a graph that is used in the README for the
	package.
*/
func CreateREADMEGraph() *position_graph.PositionGraph {
	// Constants
	g := position_graph.New()

	// Algorithm

	n1 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{-4.0, -1.0}),
	)
	n2 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{-4.0, +1.0}),
	)
	n3 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{-3.0, -1.0}),
	)
	n4 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{+3.0, +1.0}),
	)
	n5 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{0.0, 0.0}),
	)
	n6 := g.AddNodeAt(mat.NewVecDense(2, []float64{0.0, -1.5}))
	n7 := g.AddNodeAt(mat.NewVecDense(2, []float64{1.0, +1.0}))
	n8 := g.AddNodeAt(mat.NewVecDense(2, []float64{2.0, -1.5}))
	n9 := g.AddNodeAt(mat.NewVecDense(2, []float64{3.0, +1.0}))
	n10 := g.AddNodeAt(mat.NewVecDense(2, []float64{3.0, -1.0}))
	n11 := g.AddNodeAt(mat.NewVecDense(2, []float64{4.0, +1.0}))
	n12 := g.AddNodeAt(mat.NewVecDense(2, []float64{4.0, -1.0}))

	// Create all edges
	g.AddEdgeBetween(n1, n2)
	g.AddEdgeBetween(n1, n3)
	g.AddEdgeBetween(n2, n3)
	g.AddEdgeBetween(n2, n5)
	g.AddEdgeBetween(n3, n4)
	g.AddEdgeBetween(n3, n5)
	g.AddEdgeBetween(n3, n6)
	g.AddEdgeBetween(n4, n5)
	g.AddEdgeBetween(n4, n7)
	g.AddEdgeBetween(n5, n6)
	g.AddEdgeBetween(n5, n7)
	g.AddEdgeBetween(n5, n8)
	g.AddEdgeBetween(n5, n10)
	g.AddEdgeBetween(n7, n8)
	g.AddEdgeBetween(n7, n10)
	g.AddEdgeBetween(n8, n10)
	g.AddEdgeBetween(n9, n10)
	g.AddEdgeBetween(n9, n11)
	g.AddEdgeBetween(n10, n11)
	g.AddEdgeBetween(n10, n12)
	g.AddEdgeBetween(n11, n12)

	return g
}

/*
TestPlan_FindPlan3
Description:

	Testing the FindPlan method with a graph that was
	used to define the graph in the README.
*/
func TestPlan_FindPlan3(t *testing.T) {
	// Setup
	g := CreateREADMEGraph()

	// Algorithm
	p1, err := djikstra.FindPlan(g, 10, 0)
	if err != nil {
		t.Errorf("there was a problem finding the plan: %v", err)
	}

	if len(p1.Sequence) != 5 {
		t.Errorf(
			"there should only be 5 nodes in the plan, found %v",
			len(p1.Sequence),
		)
	}

	expected := []int64{10, 9, 4, 2, 0}
	for idx := 0; int(idx) < len(p1.Sequence); idx++ {
		if p1.Sequence[idx].ID() != expected[idx] {
			t.Errorf(
				"expected first node in plan to be node %v; received node %v",
				expected[idx],
				p1.Sequence[idx].ID(),
			)
		}
	}
}
