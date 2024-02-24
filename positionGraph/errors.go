package positionGraph

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
	return "Node with ID " + string(e.ID) + " not found"
}
