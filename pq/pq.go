// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
	"math"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
// func main() {
// 	// Some items and their priorities.
// 	items := map[string]int{
// 		"banana": 3, "apple": 2, "pear": 4,
// 	}

// 	// Create a priority queue, put the items in it, and
// 	// establish the priority queue (heap) invariants.
// 	pq := make(PriorityQueue, len(items))
// 	i := 0
// 	for value, priority := range items {
// 		pq[i] = &Item{
// 			value:    value,
// 			priority: priority,
// 			index:    i,
// 		}
// 		i++
// 	}
// 	heap.Init(&pq)

// 	// Insert a new item and then modify its priority.
// 	item := &Item{
// 		value:    "orange",
// 		priority: 1,
// 	}
// 	heap.Push(&pq, item)
// 	pq.update(item, item.value, 5)

// 	// Take the items out; they arrive in decreasing priority order.
// 	for pq.Len() > 0 {
// 		item := heap.Pop(&pq).(*Item)
// 		fmt.Printf("%.2d:%s ", item.priority, item.value)
// 	}
// }

type pair struct {
	dst    string
	weight int
}

func main() {

	inf := math.MaxInt32

	Nodes := map[string]Item{
		"S": {"S", 0, 0},
		"A": {"A", inf, 0},
		"B": {"B", inf, 0},
		"C": {"C", inf, 0},
		"D": {"D", inf, 0},
		"F": {"F", inf, 0},
		"G": {"G", inf, 0},
		"H": {"H", inf, 0},
		"I": {"I", inf, 0},
		"J": {"J", inf, 0},
		"K": {"K", inf, 0},
		"L": {"L", inf, 0},
		"E": {"E", inf, 0},
	}

	// Graph := map[string][]pair{
	// 	"S": []pair{pair{"A", 7}, pair{"B", 2}, pair{"C", 3}},
	// 	"A": []pair{pair{"S", 7}, pair{"B", 3}, pair{"D", 4}},
	// 	"B": []pair{pair{"S", 2}, pair{"A", 3}, pair{"D", 4}},
	// 	"C": []pair{pair{"S", 3}, pair{"L", 2}},
	// 	"D": []pair{pair{"A", 4}, pair{"B", 4}, pair{"F", 5}},
	// 	"F": []pair{pair{"D", 5}, pair{"H", 3}},
	// 	"G": []pair{pair{"H", 2}, pair{"E", 2}},
	// 	"H": []pair{pair{"B", 1}, pair{"F", 3}, pair{"G", 2}},
	// 	"I": []pair{pair{"L", 4}, pair{"J", 6}, pair{"K", 4}},
	// 	"J": []pair{pair{"L", 4}, pair{"I", 6}, pair{"K", 4}},
	// 	"K": []pair{pair{"I", 4}, pair{"J", 4}, pair{"E", 5}},
	// 	"L": []pair{pair{"C", 2}, pair{"I", 4}, pair{"J", 4}},
	// 	"E": []pair{pair{"G", 2}, pair{"K", 5}},
	// }

	// visited := map[string]bool{}
	// prev := map[string]string{}

	pq := make(PriorityQueue, len(Nodes))
	i := 0
	for _, n := range Nodes {
		pq[i] = &n
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}

	// for len(pq) > 0 {
	// 	// pop and visit the top node
	// 	node := heap.Pop(&pq).(*Item)
	// 	neighbours := Graph[node.value]
	// 	for _, n := range neighbours {

	// 	}
	// }

}
