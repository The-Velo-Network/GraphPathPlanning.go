package djikstra

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
	Graph            graph.WeightedUndirected
	CurrentGraphNode graph.Node
	PreviousInPlan   *PlanningNode
	CostToGo         float64 // The cost to go from the current node to the goal node
}

// =======
// Methods
// =======

/*
Cost
Description:

	Returns the cost of the current planning node based on the
	cost of the previous node and the edge between the previous
*/
func (pn *PlanningNode) Cost() float64 {
	return pn.CostToGo
}

/*
CalculateCostToGo
Description:

	Calculates the cost to go to the current node using:
	- the cost of the previous node,
	- the edge between the previous node and the current node, as well as
	- the heuristic cost from the current node to the goal node.
*/
func (pn *PlanningNode) CalculateCostToGo() float64 {
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
	return pn.PreviousInPlan.CostToGo + lastEdge.Weight()

}

/*
Expand
Description:

	"Expands" from the current graph node to all of the adjacent nodes.
	Returns a slice of PlanningNodes that represent the expanded nodes.
*/
func (pn *PlanningNode) Expand() []*PlanningNode {
	// Setup

	// Create a slice to hold the expanded nodes
	var expandedNodes []*PlanningNode
	neighboringNodes := pn.Graph.From(pn.CurrentGraphNode.ID())
	for neighboringNodes.Next() {
		// Get current node
		node := neighboringNodes.Node()

		// Create a new planning node for the expanded node
		expandedNode := &PlanningNode{
			Graph:            pn.Graph,
			CurrentGraphNode: node,
			PreviousInPlan:   pn,
		}

		// Calculate the costs for the expanded node
		expandedNode.CostToGo = expandedNode.CalculateCostToGo()

		// Add the expanded node to the slice
		expandedNodes = append(expandedNodes, expandedNode)
	}

	// Return
	return expandedNodes

}
