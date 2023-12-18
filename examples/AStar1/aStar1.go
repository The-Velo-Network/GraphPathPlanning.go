package main

/*
aStar1.go
Description:
	This script will demonstrate a very simple example of doing A* path planning.
*/

import (
	"gonum.org/v1/gonum/graph/simple"
)

func main() {
	// Create graph
	g := simple.NewUndirectedGraph()

	// Add a few nodes
	sStart := g.NewNode()
	g.AddNode(sStart)

	sGoal := g.NewNode()
	g.AddNode(sGoal)

	s1 := g.NewNode()
	g.AddNode(s1)

	s2 := g.NewNode()
	g.AddNode(s2)

	s3 := g.NewNode()
	g.AddNode(s3)

	// Add edges
	g.SetEdge(simple.WeightedEdge{
		F: sStart, T: s1, W: 1.0,
	})
	g.SetEdge(simple.WeightedEdge{
		F: sStart, T: s2, W: 1.0,
	})
	g.SetEdge(simple.WeightedEdge{F: sStart, T: s3, W: 12.0})
	g.SetEdge(simple.WeightedEdge{F: sStart, T: sGoal, W: 100.0})
	g.SetEdge(simple.WeightedEdge{F: s1, T: s3, W: 1})
	g.SetEdge(simple.WeightedEdge{F: s2, T: s3, W: 1})
	g.SetEdge(simple.WeightedEdge{F: s3, T: sGoal, W: 1})

}
