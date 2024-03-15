package gppErrors

import (
	"fmt"
	"gonum.org/v1/gonum/graph"
)

/*
no_path_found.go
Description:

	The error returned when there is no path found in the
	graph.
*/

// Types
// =====

type NoPathFound struct {
	Graph graph.Graph
}

// Methods
// =======

func (npf NoPathFound) Error() string {
	return fmt.Sprintf("no path found in graph %v", npf.Graph)
}
