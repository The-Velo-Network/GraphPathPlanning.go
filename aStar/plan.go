package aStar

import (
	"github.com/GraphPathPlanning.go/planningHeap"
	"gonum.org/v1/gonum/graph"
)

/*
plan.go
Description:

	Defines how plans are generated with the A* algorithm.
*/

// =======
// Functions
// =======

/*
Plan
Description:

	Generates a plan using the A* algorithm.
	To move from node start to node end through the graph
	g.
*/
func Plan(
	g graph.WeightedUndirected,
	start, end int64,
	heuristic func(int64, graph.WeightedUndirected) float64,
) (*PlanningNode, error) {
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
			return pn, nil
		}

		// Otherwise, expand the node
		expandedNodes := pn.Expand(heuristic)

		// Add the expanded nodes to the heap
		for _, newPN := range expandedNodes {
			heap0.Push(newPN)
		}

	}

	return nil, nil

}
