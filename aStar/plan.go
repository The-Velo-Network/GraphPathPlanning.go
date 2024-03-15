package aStar

import (
	"github.com/GraphPathPlanning.go/gppErrors"
	"github.com/GraphPathPlanning.go/planningHeap"
	"gonum.org/v1/gonum/graph"
	"slices"
)

/*
plan.go
Description:

	Defines how plans are generated with the A* algorithm.
*/

// ================
// Type Definitions
// ================

type Plan struct {
	Sequence []graph.Node // The sequence of nodes in the path (start @ 0, and end @ len(Sequence) - 1
	CostToGo float64
}

// =======
// Functions
// =======

/*
FindPlan
Description:

	Generates a plan using the A* algorithm.
	To move from node start to node end through the graph
	g.
*/
func FindPlan(
	g graph.WeightedUndirected,
	start, end int64,
	heuristic func(*PlanningNode) float64,
) (*Plan, error) {
	// Constants

	// Create initial planning node and heap
	pn0 := &PlanningNode{
		Graph:            g,
		CurrentGraphNode: g.Node(start),
		PreviousInPlan:   nil,
		CostToGo:         0.0,
		HeuristicCost:    0.0,
	}

	var heap0 planningHeap.PlanningHeap
	heap0 = append(heap0, pn0)

	// Algorithm
	for len(heap0) > 0 {
		// Pop the top node off the heap
		pn := heap0.Pop().(*PlanningNode)

		// If we have reached the end, return the plan
		if pn.CurrentGraphNode.ID() == end {
			return UnrollPlanFrom(pn), nil
		}

		// Otherwise, expand the node
		expandedNodes := pn.Expand(heuristic)

		// Add the expanded nodes to the heap
		for _, newPN := range expandedNodes {
			heap0.Push(newPN)
		}

	}

	return nil, gppErrors.NoPathFound{g}

}

/*
UnrollPlanFrom()
Description:

	Unrolls a plan from a given planning node.
	Hopefully, you try this only on a planning node
	that reaches the end.
*/
func UnrollPlanFrom(pn *PlanningNode) *Plan {
	// Check to see if plan is empty
	if pn == nil {
		return nil
	}

	// Iterate through each of the nodes in the plan
	current := pn
	var reversedPlan []graph.Node
	for current != nil {
		// Add current node (in graph) to plan
		reversedPlan = append(
			reversedPlan,
			current.Graph.Node(current.CurrentGraphNode.ID()),
		)

		// Update current
		current = current.PreviousInPlan
	}

	// Return result
	forwardPlan := reversedPlan
	slices.Reverse(forwardPlan)

	return &Plan{
		Sequence: forwardPlan,
		CostToGo: pn.CostToGo,
	}
}
