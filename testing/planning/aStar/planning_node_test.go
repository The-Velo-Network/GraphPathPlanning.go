package aStar_test

import (
	positionGraph2 "github.com/GraphPathPlanning.go/graphs/position"
	"github.com/GraphPathPlanning.go/planning/aStar"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
	"math"
	"testing"
)

/*
planning_node_test.go
Description:

	Defines tests for the PlanningNode struct.
*/

/*
CreateTestGraph2
Description:

	Creates a simple test graph for use in these tests.
*/
func CreateTestGraph2() *positionGraph2.PositionGraph {
	// Constants
	g := positionGraph2.New()

	// Algorithm

	n1 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{0.0, 0.0}),
	)
	n2 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{1.0, 0.0}),
	)
	g.AddEdgeBetween(n1, n2)

	return g
}

/*
CreateTestGraph_PlanningNodeTest2
Description:

	Creates a simple test graph for use in these tests.
*/
func CreateTestGraph_PlanningNodeTest2() *positionGraph2.PositionGraph {
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
		mat.NewVecDense(2, []float64{2.0, 1.0}),
	)
	n4 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{3.0, 3.0}),
	)

	// Add nodes
	g.AddEdgeBetween(n1, n2)
	g.AddEdgeBetween(n2, n3)
	g.AddEdgeBetween(n3, n4)

	return g
}

/*
TestPlanningNode_Cost1
Description:

	Tests the Cost() method of the PlanningNode struct.
	Specifically, this tests that the cost of an initial planning node
	is 0. Such a node is one that has no previous node.
*/
func TestPlanningNode_Cost1(t *testing.T) {
	// Constants
	g2 := CreateTestGraph2()

	// Algorithm
	nodes := g2.Nodes()
	nodes.Next()
	//targetNode, ok := g1.NodeWithID(nodes.Node().ID())
	//if !ok {
	//	t.Errorf("Failed to retrieve node with ID %v", targetID)
	//}

	n2 := nodes.Node()
	pn := aStar.PlanningNode{
		Graph:            g2,
		CurrentGraphNode: n2,
		PreviousInPlan:   nil,
	}

	if pn.Cost() != 0.0 {
		t.Errorf("Expected cost to be 0.0, but got %f", pn.Cost())
	}
}

/*
TestPlanningNode_Cost2
Description:

	Tests the Cost() method of the PlanningNode struct.
	Specifically, this tests that the cost of a planning node
	is the cost of the previous node plus the cost of the edge
	between the previous node and the current node.
*/
func TestPlanningNode_Cost2(t *testing.T) {
	// Constants
	g2 := CreateTestGraph2()

	pn1 := aStar.PlanningNode{
		Graph:            g2,
		CurrentGraphNode: g2.Node(0),
	}

	pn2 := aStar.PlanningNode{
		Graph:            g2,
		CurrentGraphNode: g2.Node(1),
		PreviousInPlan:   &pn1,
	}

	// Algorithm
	if pn2.Cost() > 0.0 {
		t.Errorf("Expected cost to be greater than 0.0, but got %f", pn2.Cost())
	}
}

/*
TestPlanningNode_UpdateCosts1
Description:

	Tests the CalculateCostToGo() method of the PlanningNode struct.
*/
func TestPlanningNode_UpdateCosts1(t *testing.T) {
	// Setup
	g := CreateTestGraph_PlanningNodeTest2()

	// Setup
	pn1 := aStar.PlanningNode{
		Graph:            g,
		CurrentGraphNode: g.Node(0),
	}
	goalIdx := int64(3)

	simpleHeuristic := func(currPN *aStar.PlanningNode) float64 {
		// Setup
		currentIdx := currPN.CurrentGraphNode.ID()

		wu := currPN.Graph
		goalNode := wu.Node(goalIdx).(*positionGraph2.Node)
		currNode := wu.Node(currentIdx).(*positionGraph2.Node)

		// Algorithm
		var diff mat.VecDense
		diff.SubVec(goalNode.Position, currNode.Position)
		return mat.Norm(&diff, 2)
	}

	// Algorithm
	pn1.UpdateCosts(simpleHeuristic)

	// Check that the costs is zero for this planning node.
	if pn1.CostToGo != 0.0 {
		t.Errorf("Expected cost to go to be 0.0, but got %f", pn1.CostToGo)
	}

	// Heursitic Cost should be greater than 0
	if pn1.Cost() <= pn1.CostToGo {
		t.Errorf(
			"expected plan cost to be greater than %v, but got %f",
			pn1.CostToGo,
			pn1.Cost(),
		)
	}

}

/*
TestPlanningNode_UpdateCosts2
Description:

	Tests the CalculateCostToGo() method of the PlanningNode struct.
	Specifically, we consider a planning node that has two previous nodes
	and we check that the cost to go is the cost of the previous node
	plus the cost of the edge between the previous node and the current node.
*/
func TestPlanningNode_UpdateCosts2(t *testing.T) {
	// Setup
	g := CreateTestGraph_PlanningNodeTest2()

	goalIdx := int64(3)
	simpleHeuristic := func(currPN *aStar.PlanningNode) float64 {
		// Setup
		currentIdx := currPN.CurrentGraphNode.ID()
		wu := currPN.Graph

		goalNode := wu.Node(goalIdx).(*positionGraph2.Node)
		currNode := wu.Node(currentIdx).(*positionGraph2.Node)

		// Algorithm
		var diff mat.VecDense
		diff.SubVec(goalNode.Position, currNode.Position)
		return mat.Norm(&diff, 2) + currPN.CostToGo
	}

	// - Create nodes
	pn1 := aStar.PlanningNode{
		Graph:            g,
		CurrentGraphNode: g.Node(0),
	}

	pn2 := aStar.PlanningNode{
		Graph:            g,
		CurrentGraphNode: g.Node(1),
		PreviousInPlan:   &pn1,
	}

	pn3 := aStar.PlanningNode{
		Graph:            g,
		CurrentGraphNode: g.Node(2),
		PreviousInPlan:   &pn2,
	}

	// Algorithm
	pn1.UpdateCosts(simpleHeuristic)
	pn2.UpdateCosts(simpleHeuristic)
	pn3.UpdateCosts(simpleHeuristic)

	// Check that the costs is zero for this planning node.
	if !floats.EqualApprox(
		[]float64{1.0 + math.Sqrt(2.0)},
		[]float64{pn3.CostToGo},
		0.0001,
	) {
		t.Errorf("Expected cost to go to be 1.0 + sqrt(2.0), but got %f", pn3.CostToGo)
	}

	// Heursitic Cost should be greater than CostToGo
	if pn3.HeuristicCost <= pn3.CostToGo {
		t.Errorf("Expected heuristic cost to be greater than 0.0, but got %f", pn2.HeuristicCost)
	}
}

/*
TestPlanningNode_Expand1
Description:

	This function tests that a node that is not connected to
	anything in the graph returns no planning nodes in the
	result of the expand function.
*/
func TestPlanningNode_Expand1(t *testing.T) {
	// Setup Graph
	g := positionGraph2.New()

	n1 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{1.0, 2.0}),
	)
	n2 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{1.0, 3.0}),
	)
	n3 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{2.0, 3.0}),
	)

	g.AddEdgeBetween(n2, n3)

	// Setup planning node to expand
	pn4 := aStar.PlanningNode{
		Graph:            g,
		CurrentGraphNode: &n1,
		PreviousInPlan:   nil,
		CostToGo:         0.0,
		HeuristicCost:    0.0,
	}

	// Setup heuristic
	goalIdx := n3.ID()
	simple_heuristic := func(currPN *aStar.PlanningNode) float64 {
		// Setup
		currentIdx := currPN.CurrentGraphNode.ID()
		wu := currPN.Graph

		goalNode := wu.Node(goalIdx).(*positionGraph2.Node)
		currNode := wu.Node(currentIdx).(*positionGraph2.Node)

		// Algorithm
		var diff mat.VecDense
		diff.SubVec(goalNode.Position, currNode.Position)
		return mat.Norm(&diff, 2) + currPN.CostToGo
	}

	// Compute expansion
	result := pn4.Expand(simple_heuristic)

	// Compute the number of nodes in the result
	if len(result) != 0 {
		t.Errorf(
			"expected for expansion to create 0 nodes, received %v.",
			len(result),
		)
	}

}

/*
TestPlanningNode_Expand2
Description:

	This function tests that a node that is connected to
	three other nodes in the graph returns three planning
	nodes in the result of the expand function.
*/
func TestPlanningNode_Expand2(t *testing.T) {
	// Setup Graph
	g := positionGraph2.New()

	n1 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{1.0, 2.0}),
	)
	n2 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{1.0, 3.0}),
	)
	n3 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{2.0, 3.0}),
	)
	n4 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{0.0, 0.0}),
	)
	n5 := g.AddNodeAt(
		mat.NewVecDense(2, []float64{-1.0, 0.2}),
	)

	g.AddEdgeBetween(n1, n2)
	g.AddEdgeBetween(n1, n3)
	g.AddEdgeBetween(n1, n4)
	g.AddEdgeBetween(n3, n5)

	// Setup planning node to expand
	pn4 := aStar.PlanningNode{
		Graph:            g,
		CurrentGraphNode: &n1,
		PreviousInPlan:   nil,
		CostToGo:         0.0,
		HeuristicCost:    0.0,
	}

	// Setup heuristic
	goalIdx := n3.ID()
	simple_heuristic := func(currPN *aStar.PlanningNode) float64 {
		// Setup
		currentIdx := currPN.CurrentGraphNode.ID()
		wu := currPN.Graph

		goalNode := wu.Node(goalIdx).(*positionGraph2.Node)
		currNode := wu.Node(currentIdx).(*positionGraph2.Node)

		// Algorithm
		var diff mat.VecDense
		diff.SubVec(goalNode.Position, currNode.Position)
		return mat.Norm(&diff, 2) + currPN.CostToGo
	}

	// Compute expansion
	result := pn4.Expand(simple_heuristic)

	// Compute the number of nodes in the result
	if len(result) != 3 {
		t.Errorf(
			"expected for expansion to create 0 nodes, received %v.",
			len(result),
		)
	}
}
