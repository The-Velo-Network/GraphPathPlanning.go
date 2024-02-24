package aStar_test

import (
	"github.com/GraphPathPlanning.go/aStar"
	"github.com/GraphPathPlanning.go/positionGraph"
	"gonum.org/v1/gonum/mat"
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
func CreateTestGraph2() *positionGraph.PositionGraph {
	// Constants
	g := positionGraph.NewPositionGraph()

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
