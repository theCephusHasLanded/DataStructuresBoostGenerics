package go_data_structures


type QueueInterface[T any] interface {
	Front() T      // returns first item in queue. O(1)
	Enqueue(val T) // adds val to end of queue, increases size by 1. O(1)
	Dequeue()      // removes item from front of queue, decreases size by 1. O(1)
	Empty() bool   // returns whether queue is empty. O(1)
	Size() int     // returns number of elements in queue. O(1) or O(n) depending on implementation
}

type Queue[T any] struct {
	list SinglyLinkedListInterface[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{list: NewSinglyLinkedList[T]()} // Assuming a constructor for the singly linked list
}

func (q *Queue[T]) Front() T {
	return q.list.PeekFront()
}

func (q *Queue[T]) Enqueue(val T) {
	q.list.PushBack(val)
}

func (q *Queue[T]) Dequeue() {
	q.list.PopFront()
}

func (q *Queue[T]) Empty() bool {
	return q.list.IsEmpty()
}

func (q *Queue[T]) Size() int {
	return q.list.Count()
}

type AlternateQueue[T any] struct {
	arr []T
}

func NewAlternateQueue[T any]() *AlternateQueue[T] {
	return &AlternateQueue[T]{arr: make([]T, 0)}
}

func (q *AlternateQueue[T]) Front() T {
	if len(q.arr) == 0 {
		var zero T // Return zero value of T if queue is empty
		return zero
	}
	return q.arr[0]
}

func (q *AlternateQueue[T]) Enqueue(val T) {
	q.arr = append(q.arr, val)
}

func (q *AlternateQueue[T]) Dequeue() {
	if len(q.arr) == 0 {
		return // Prevent panic if the queue is empty
	}
	q.arr = q.arr[1:]
}

func (q *AlternateQueue[T]) Empty() bool {
	return len(q.arr) == 0
}

func (q *AlternateQueue[T]) Size() int {
	return len(q.arr)
}
