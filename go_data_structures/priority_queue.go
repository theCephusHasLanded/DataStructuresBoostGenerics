package go_data_structures
import (
	"container/heap"
	"golang.org/x/exp/constraints"
)

type PriorityQueueInterface[T constraints.Ordered, V any] interface {
	Front() Pair[T, V]         // returns first item in queue. O(1)
	Enqueue(priority T, val V) // adds val to queue with priority, increases size by 1. O(1)
	Dequeue()                  // removes the highest priority item form queue, decreases size by 1. O(1)
	Empty() bool               // returns whether queue is empty. O(1)
	Size() int                 // returns number of elements in queue. O(1)
}



// PriorityQueue implements a priority queue using a heap.
type PriorityQueue[T constraints.Ordered, V any] struct {
	items []Pair[T, V]
}

// ensure PriorityQueue implements heap.Interface
var _ heap.Interface = (*PriorityQueue[int, any])(nil)

// Implement heap.Interface for PriorityQueue.
func (pq *PriorityQueue[T, V]) Len() int { return len(pq.items) }
func (pq *PriorityQueue[T, V]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq.items[i].Priority > pq.items[j].Priority
}
func (pq *PriorityQueue[T, V]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}
func (pq *PriorityQueue[T, V]) Push(x any) {
	item := x.(Pair[T, V])
	pq.items = append(pq.items, item)
}
func (pq *PriorityQueue[T, V]) Pop() any {
	old := pq.items
	n := len(old)
	item := old[n-1]
	pq.items = old[0 : n-1]
	return item
}

// NewPriorityQueue creates a new PriorityQueue.
func NewPriorityQueue[T constraints.Ordered, V any]() *PriorityQueue[T, V] {
	pq := &PriorityQueue[T, V]{}
	heap.Init(pq)
	return pq
}

// Front returns the first item in the queue without removing it.
func (pq *PriorityQueue[T, V]) Front() Pair[T, V] {
	return pq.items[0]
}

// Enqueue adds an item to the queue with a given priority.
func (pq *PriorityQueue[T, V]) Enqueue(priority T, val V) {
	heap.Push(pq, Pair[T, V]{Priority: priority, Value: val})
}

// Dequeue removes and returns the highest priority item from the queue.
func (pq *PriorityQueue[T, V]) Dequeue() {
	heap.Pop(pq)
}

// Empty returns whether the queue is empty.
func (pq *PriorityQueue[T, V]) Empty() bool {
	return pq.Len() == 0
}

// Size returns the number of elements in the queue.
func (pq *PriorityQueue[T, V]) Size() int {
	return pq.Len()
}
