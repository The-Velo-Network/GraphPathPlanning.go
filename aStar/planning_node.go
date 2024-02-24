package aStar

import (
	"gonum.org/v1/gonum/graph"
)

/*
node.go
Description:

	Defines an implementation of a PlanningNode for use in the A* algorithm.
*/

// =======
// Objects
// =======

/*
PlanningNode
Description:

	An implementation of a PlanningNode for use in the A* algorithm.
*/
type PlanningNode struct {
	Graph            graph.WeightedUndirected //TODO: Change to graph.WeightedUndirected (a more general type)
	CurrentGraphNode graph.Node
	PreviousInPlan   *PlanningNode
}

// =======
// Methods
// =======

/*
Cost
Description:

	Returns the cost of the current plan based on the .
*/
func (pn *PlanningNode) Cost() float64 {
	//Constants
	underlyingGraph := pn.Graph

	// Algorithm
	// If the previous node is nil, then we are at the start of the path.
	// In this case, we should return 0.
	if pn.PreviousInPlan == nil {
		return 0.0
	}

	// Otherwise, we should return the cost of the previous node
	// plus the cost of the edge between the previous node and the current node.
	lastEdge := underlyingGraph.WeightedEdgeBetween(
		pn.PreviousInPlan.CurrentGraphNode.ID(),
		pn.CurrentGraphNode.ID(),
	)
	return pn.PreviousInPlan.Cost() + lastEdge.Weight()

}
