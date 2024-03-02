package positionGraph

import "fmt"

/*
errors.go
Description:

	Defines the errors for the position graph.
*/

// =======
// Errors
// =======

type NodeNotFoundError struct {
	ID int64
}

func (e NodeNotFoundError) Error() string {
	return fmt.Sprintf(
		"Node with ID %v not found",
		e.ID,
	)
}
