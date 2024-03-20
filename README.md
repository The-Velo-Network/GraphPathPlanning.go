[![codecov](https://codecov.io/gh/The-Velo-Network/GraphPathPlanning.go/graph/badge.svg?token=hwCaHaHIgi)](https://codecov.io/gh/The-Velo-Network/GraphPathPlanning.go)

# GraphPathPlanning.go
A collection of single- or multi-objective path planning algorithms.

## Installation

```bash
go get github.com/The-Velo-Network/GraphPathPlanning.go
```

## Usage

Once you have defined the graph you want to plan over, you can use one of the built-in
planners with a single call. For example, to use Djikstra's algorithm:
```go
package main

import (
	"fmt"
	position_graph "github.com/The-Velo-Network/GraphPathPlanning.go/graphs/position"
	"github.com/GraphPathPlanning.go/planning/djikstra"
	"gonum.org/v1/gonum/mat"
)

func main() {

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
		panic(fmt.Errorf("Error finding plan: %w", err))
	}
}
```

To use a planner like the A* algorithm, you can use the `FindPlan` method in the same way:
```go
package main

import (
	"fmt"
	position_graph "github.com/The-Velo-Network/GraphPathPlanning.go/graphs/position"
	"github.com/GraphPathPlanning.go/planning/aStar"
	"gonum.org/v1/gonum/mat"
)

func main() {
	// Setup Graph
	g := position_graph.New()

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

		goalNode := wu.Node(goalIdx).(*positionGraph2.Node)
		currNode := wu.Node(currentIdx).(*positionGraph2.Node)

		// Algorithm
		var diff mat.VecDense
		diff.SubVec(goalNode.Position, currNode.Position)
		return mat.Norm(&diff, 2)
	}

	// Apply Plan method
	p1, err := aStar.FindPlan(g, n1.ID(), n2.ID(), simpleHeuristic)
	if err != nil {
		t.Errorf("there was a problem finding the plan: %v", err)
	}
}

```

### More Detailed Usage

We've built the module to be flexible in that you can plan with ANY object that
implements the `Graph` interface provided by [gonum/graph](https://pkg.go.dev/gonum.org/v1/gonum/graph).

To see an example of how you can create your own graph, look at example graph definitions
like the `PositionGraph` in `graphs/position`.


## Related Work

- [paths](https://github.com/SolarLune/paths) by SolarLune is a great
    collection of pathfinding algorithms for 2D grids. (This library is meant to
    be used on arbitrary graphs (not just grids) and can work in any dimension.)
- [golangGeojsonDjikstra](https://github.com/pitchinnate/golangGeojsonDijkstra) by pitchinnate
    contains an implementation of Djikstra's algorithm for GeoJSON data in databases.