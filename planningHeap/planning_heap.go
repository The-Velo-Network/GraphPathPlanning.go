package planningHeap

type PlanningHeap []PlanningNode

func (ph PlanningHeap) Len() int { return len(ph) }
func (ph PlanningHeap) Less(i, j int) bool {
	//Description:
	//	Returns true iff ph[i] < ph[j] in terms of the Cost

	return ph[i].Cost() < ph[j].Cost()
}
func (ph PlanningHeap) Swap(i, j int) {
	//Description:
	//	Swaps the elements h[i] and h[j] in the heap.

	ph[i], ph[j] = ph[j], ph[i]
}

func (ph *PlanningHeap) Push(x any) {
	//Description:
	//	Adds a planning node to the heap pointed to by ph.
	//	Should automatically sort using our Less() method.

	item := x.(PlanningNode)
	*ph = append(*ph, item)
}

func (ph *PlanningHeap) Pop() any {
	//Description:
	//	Retrieves the element with the lowest cost.
	//	Tries to avoid memory leaks.

	//Constants

	// Algorithm
	old := *ph
	n := len(old)

	tempPlanningNode := old[n-1]
	old[n-1] = nil // avoid memory leak. Why doesn't this work?
	*ph = old[0 : n-1]
	return tempPlanningNode
}
