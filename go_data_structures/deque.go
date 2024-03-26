package go_data_structures

  type DequeInterface[T any] interface {
    Front() T            // Returns the first item in deque. O(1)
    Back() T             // Returns the last item in deque. O(1)
    PushFront(val T)     // Adds val to the front of the queue, increases size by 1. O(1)
    PushBack(val T)      // Adds val to the back of the queue, increases size by 1. O(1)
    PopFront() (T, bool) // Removes and returns the item from the front of the queue. Returns false if empty.
    PopBack() (T, bool)  // Removes and returns the item from the back of the queue. Returns false if empty.
    Empty() bool         // Returns whether the queue is empty. O(1)
    Size() int           // Returns the number of elements in the queue. O(1)
}



// Deque implements a double-ended queue using a doubly linked list.
type Deque[T any] struct {
  list *DoublyLinkedList[T] // Assume DoublyLinkedList implements all needed methods
}

// NewDeque creates a new Deque instance.
func NewDeque[T any]() *Deque[T] {
  return &Deque[T]{list: NewDoublyLinkedList[T]()}
}

func (dq *Deque[T]) PopFront() (T, bool) {
  if dq.Empty() {
      var zero T // Zero value for type T
      return zero, false // Indicate failure due to empty deque
  }
  // Assume list has a method PopFront that returns (T, bool)
  return dq.list.PopFront()
}

func (dq *Deque[T]) Front() (T, bool) {
  if dq.Empty() {
      var zero T // Zero value for type T
      return zero, false // Indicate failure due to empty deque
  }
  // Assume list has a method Front that returns T
  return dq.list.Front()
}


func (dq *Deque[T]) Back() (T, bool) {
  if dq.Empty() {
      var zero T // Return zero value of T if deque is empty
      return zero, false
  }
  return dq.list.Back()
}


func (dq *Deque[T]) PushFront(val T) {
  dq.list.PushFront(val)
}

func (dq *Deque[T]) PushBack(val T) {
  dq.list.PushBack(val)
}

func (dq *Deque[T]) PopBack() {
  if !dq.Empty() {
      dq.list.PopBack()
  }
}

func (dq *Deque[T]) Empty() bool {
  return dq.list.Empty()
}

func (dq *Deque[T]) Size() int {
  return dq.list.Size()
}
