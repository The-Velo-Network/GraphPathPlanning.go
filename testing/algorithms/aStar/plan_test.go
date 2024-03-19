package aStar_test

import (
	aStar2 "github.com/GraphPathPlanning.go/algorithms/aStar"
	"github.com/GraphPathPlanning.go/gppErrors"
	"github.com/GraphPathPlanning.go/positionGraph"
	"gonum.org/v1/gonum/mat"
	"testing"
)

/*
plan_test.go
Description:

	This file is meant to test all methods defined in the plan
	file for A*.
*/

/*
TestPlan_FindPlan1
Description:

	This test is meant to verify that the FindPlan()
	method finds a correct plan for a simple graph
	(with 2 nodes and one edge between them).
*/
func TestPlan_FindPlan1(t *testing.T) {
	// Setup Graph
	g := positionGraph.NewPositionGraph()

	n1 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{1.0, 2.0}),
	)
	n2 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{2.0, 2.0}),
	)

	g.AddEdgeBetween(n1, n2)

	// Setup Path Planning Heuristic
	goalIdx := n2.ID()
	simpleHeuristic := func(currPN *aStar2.PlanningNode) float64 {
		// Setup
		currentIdx := currPN.CurrentGraphNode.ID()
		wu := currPN.Graph

		goalNode := wu.Node(goalIdx).(*positionGraph.Node)
		currNode := wu.Node(currentIdx).(*positionGraph.Node)

		// Algorithm
		var diff mat.VecDense
		diff.SubVec(goalNode.Position, currNode.Position)
		return mat.Norm(&diff, 2)
	}

	// Apply Plan method
	p1, err := aStar2.FindPlan(g, n1.ID(), n2.ID(), simpleHeuristic)
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
	g := positionGraph.NewPositionGraph()

	n1 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{1.0, 2.0}),
	)
	n2 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{2.0, 2.0}),
	)

	// Setup Path Planning Heuristic
	goalIdx := n2.ID()
	simpleHeuristic := func(currPN *aStar2.PlanningNode) float64 {
		// Setup
		currentIdx := currPN.CurrentGraphNode.ID()
		wu := currPN.Graph

		goalNode := wu.Node(goalIdx).(*positionGraph.Node)
		currNode := wu.Node(currentIdx).(*positionGraph.Node)

		// Algorithm
		var diff mat.VecDense
		diff.SubVec(goalNode.Position, currNode.Position)
		return mat.Norm(&diff, 2)
	}

	// Apply Plan method
	_, err := aStar2.FindPlan(g, n1.ID(), n2.ID(), simpleHeuristic)
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
