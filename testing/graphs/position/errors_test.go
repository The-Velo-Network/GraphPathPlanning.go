package position_graph_test

/*
errors_test.go
Description:

	This method tests the errors defined in errors.go
*/

import (
	"fmt"
	position_graph "github.com/GraphPathPlanning.go/graphs/position"
	"strings"
	"testing"
)

func TestErrors_NodeNotFoundError1(t *testing.T) {
	// Constants
	tempError := position_graph.NodeNotFoundError{
		ID: 1,
	}

	// Test
	if !strings.Contains(
		tempError.Error(),
		fmt.Sprintf("Node with ID %v not found", tempError.ID),
	) {
		t.Errorf("unexpected error: %v", tempError)
	}
}
