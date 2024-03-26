package go_data_structures



type StackInterface[T any] interface {
	Top() T      // returns top item in stack. O(1)
	Push(val T)  // adds val to top of stack, increases size by 1. O(1)
	Pop()        // removes item from top of stack, decreases size by 1. O(1)
	Empty() bool // returns whether stack is empty. O(1)
	Size() int   // returns number of elements in stack. O(1) or O(n) depending on implementation
}

type Stack[T any] struct {
	list SinglyLinkedListInterface[T]
}

type AlternateStack[T any] struct {
	arr []T // array implementation of a stack
}
