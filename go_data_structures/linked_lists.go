package go_data_structures

type SingleLinkNode[T any] struct {
	Data T
	Next *SingleLinkNode[T]
}

type SinglyLinkedList[T any] struct {
	head *SingleLinkNode[T]
	tail *SingleLinkNode[T] // Tail is added for efficient InsertAtEnd operations.
}

type SinglyLinkedListInterface[T any] interface {
	InsertAfter(val T, prev *SingleLinkNode[T]) // create new node with data val after node prev, increase size by 1. O(1)
	RemoveAfter(prev *SingleLinkNode[T])        // remove node after node prev, decrease size by 1. O(1)
	InsertAtFront(val T)                        // create node with data val at front of list, increase size by 1. O(1)
	RemoveAtFront()                             // remove node at front of list, decrease size by 1. O(1)
	InsertAtEnd(val T)                          // create node with data val at end of list, increase size by 1. O(n)
	RemoveAtEnd()                               // remove node at end of list, decrease size by 1. O(n)
	Head() *SingleLinkNode[T]                   // return first node in list. O(1)
	Tail() *SingleLinkNode[T]                   // return last node in list. O(n)
	Empty() bool                                // returns whether list is empty. O(1)
	Size() int                                  // returns number of elements in list. O(1) or O(n) depending on implementation
	PeekFront() T
    PushBack(val T)
    PopFront() T
    IsEmpty() bool
    Count() int
}

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

func (sll *SinglyLinkedList[T]) InsertAfter(val T, prev *SingleLinkNode[T]) {
	if prev == nil {
		// Insert at front if prev is nil or list is empty.
		sll.InsertAtFront(val)
		return
	}
	newNode := &SingleLinkNode[T]{Data: val, Next: prev.Next}
	prev.Next = newNode
	if prev.Next == nil {
		// Update tail if inserted at the end.
		sll.tail = newNode
	}
}

func (sll *SinglyLinkedList[T]) RemoveAfter(prev *SingleLinkNode[T]) {
	if prev == nil || prev.Next == nil {
		return // Nothing to remove.
	}
	removing := prev.Next
	prev.Next = removing.Next
	if removing.Next == nil {
		// Update tail if removed node was the last one.
		sll.tail = prev
	}
}

func (sll *SinglyLinkedList[T]) InsertAtFront(val T) {
	newNode := &SingleLinkNode[T]{Data: val, Next: sll.head}
	sll.head = newNode
	if sll.tail == nil {
		// If list was empty, new node is also the tail.
		sll.tail = newNode
	}
}

func (sll *SinglyLinkedList[T]) RemoveAtFront() {
	if sll.head == nil {
		return // List is empty.
	}
	sll.head = sll.head.Next
	if sll.head == nil {
		// List became empty, so update tail.
		sll.tail = nil
	}
}

func (sll *SinglyLinkedList[T]) InsertAtEnd(val T) {
	newNode := &SingleLinkNode[T]{Data: val, Next: nil}
	if sll.tail == nil {
		// List is empty.
		sll.head = newNode
		sll.tail = newNode
	} else {
		sll.tail.Next = newNode
		sll.tail = newNode
	}
}

func (sll *SinglyLinkedList[T]) RemoveAtEnd() {
	if sll.head == nil {
		return // List is empty.
	}
	if sll.head.Next == nil {
		// List has only one element.
		sll.head = nil
		sll.tail = nil
		return
	}
	// Iterate to find the second last node.
	curr := sll.head
	for curr.Next != sll.tail {
		curr = curr.Next
	}
	curr.Next = nil
	sll.tail = curr
}

func (sll *SinglyLinkedList[T]) Head() *SingleLinkNode[T] {
	return sll.head
}

func (sll *SinglyLinkedList[T]) Tail() *SingleLinkNode[T] {
	return sll.tail
}

func (sll *SinglyLinkedList[T]) Empty() bool {
	return sll.head == nil
}

// Size() would require iteration over the list to count elements, as size tracking is commented out.

type DoubleLinkNode[T any] struct {
	Value T
	Next *DoubleLinkNode[T]
	Prev *DoubleLinkNode[T]
}

type DoublyLinkedList[T any] struct {
	head *DoubleLinkNode[T]
	tail *DoubleLinkNode[T]
	size int
}

type DoublyLinkedListInterface[T any] interface {
	InsertAfter(val T, prev *DoubleLinkNode[T])  // create new node with data val after node prev, increase size by 1 O(1)
	InsertBefore(val T, next *DoubleLinkNode[T]) // create new node with data val before node next, increase size by 1 O(1)
	Remove(node *DoubleLinkNode[T])              // remove node , decrease size by 1 O(1)
	InsertAtFront(val T)                         // create node with data val at front of list, increase size by 1. O(1)
	RemoveAtFront()                              // remove node at front of list, decrease size by 1. O(1)
	InsertAtEnd(val T)                           // create node with data val at end of list, increase size by 1. O(n)
	RemoveAtEnd()                                // remove node at end of list, decrease size by 1. O(n)
	Head() *DoubleLinkNode[T]                    // return first node in list. O(1)
	Tail() *DoubleLinkNode[T]                    // return last node in list. O(1)
	Empty() bool                                 // returns whether list is empty. O(1)
	Size() int                                   // returns number of elements in list. O(1) or O(n) depending on implementation

}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

// PopFront removes and returns the first element from the list.
func (dll *DoublyLinkedList[T]) PopFront() (T, bool) {
    if dll.head == nil { // Check if the list is empty
        var zero T
        return zero, false
    }
    value := dll.head.Value // Store the value to return
    dll.head = dll.head.Next // Move the head pointer to the next node
    if dll.head != nil {
        dll.head.Prev = nil // Remove backward link from the new head
    } else {
        dll.tail = nil // If the list is now empty, also clear the tail
    }
    dll.size-- // Decrement the size of the list
    return value, true
}

func (dll *DoublyLinkedList[T]) Front() (T, bool) {
    if dll.head == nil {
        var zero T
        return zero, false
    }
    return dll.head.Value, true
}

func (dll *DoublyLinkedList[T]) Back() (T, bool) {
    if dll.tail == nil {
        var zero T
        return zero, false
    }
    return dll.tail.Value, true
}

func (dll *DoublyLinkedList[T]) PushFront(val T) {
    newNode := &DoubleLinkNode[T]{Value: val, Next: dll.head}
    if dll.head != nil {
        dll.head.Prev = newNode
    } else {
        dll.tail = newNode // List was empty
    }
    dll.head = newNode
    dll.size++
}

func (dll *DoublyLinkedList[T]) PushBack(val T) {
    newNode := &DoubleLinkNode[T]{Value: val, Prev: dll.tail}
    if dll.tail != nil {
        dll.tail.Next = newNode
    } else {
        dll.head = newNode // List was empty
    }
    dll.tail = newNode
    dll.size++
}

func (dll *DoublyLinkedList[T]) PopBack() (T, bool) {
    if dll.tail == nil {
        var zero T
        return zero, false
    }
    value := dll.tail.Value
    dll.tail = dll.tail.Prev
    if dll.tail == nil {
        dll.head = nil // List became empty
    } else {
        dll.tail.Next = nil
    }
    dll.size--
    return value, true
}


func (dll *DoublyLinkedList[T]) InsertAfter(val T, prev *DoubleLinkNode[T]) {
    newNode := &DoubleLinkNode[T]{Value: val, Prev: prev}
    if prev != nil {
        newNode.Next = prev.Next
        prev.Next = newNode
        if newNode.Next != nil {
            newNode.Next.Prev = newNode
        } else {
            dll.tail = newNode // Update tail if newNode is now last
        }
    } else {
        // If prev is nil, insert at front
        dll.InsertAtFront(val)
    }
}


func (dll *DoublyLinkedList[T]) InsertBefore(val T, next *DoubleLinkNode[T]) {
	newNode := &DoubleLinkNode[T]{Value: val, Next: next}
	if next != nil {
		newNode.Prev = next.Prev
		next.Prev = newNode
		if newNode.Prev != nil {
			newNode.Prev.Next = newNode
		} else {
			dll.head = newNode // Update head if newNode is now first
		}
	} else {
		// If next is nil, insert at end
		dll.InsertAtEnd(val)
	}
}

func (dll *DoublyLinkedList[T]) Remove(node *DoubleLinkNode[T]) {
	if node == nil {
		return
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		dll.head = node.Next // Update head if removing first node
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		dll.tail = node.Prev // Update tail if removing last node
	}
}

func (dll *DoublyLinkedList[T]) InsertAtFront(val T) {
	newNode := &DoubleLinkNode[T]{Value: val, Next: dll.head}
	if dll.head != nil {
		dll.head.Prev = newNode
	}
	dll.head = newNode
	if dll.tail == nil {
		dll.tail = newNode // List was empty before insertion
	}
}

func (dll *DoublyLinkedList[T]) InsertAtEnd(val T) {
	newNode := &DoubleLinkNode[T]{Value: val, Prev: dll.tail}
	if dll.tail != nil {
		dll.tail.Next = newNode
	}
	dll.tail = newNode
	if dll.head == nil {
		dll.head = newNode // List was empty before insertion
	}
}

func (dll *DoublyLinkedList[T]) RemoveAtFront() {
	if dll.head == nil {
		return // List is empty
	}
	dll.head = dll.head.Next
	if dll.head != nil {
		dll.head.Prev = nil
	} else {
		dll.tail = nil // List became empty
	}
}

func (dll *DoublyLinkedList[T]) RemoveAtEnd() {
	if dll.tail == nil {
		return // List is empty
	}
	dll.tail = dll.tail.Prev
	if dll.tail != nil {
		dll.tail.Next = nil
	} else {
		dll.head = nil // List became empty
	}
}

func (dll *DoublyLinkedList[T]) Head() *DoubleLinkNode[T] {
	return dll.head
}

func (dll *DoublyLinkedList[T]) Tail() *DoubleLinkNode[T] {
	return dll.tail
}

func (dll *DoublyLinkedList[T]) Empty() bool {
	return dll.head == nil
}

func (dll *DoublyLinkedList[T]) Size() int {
    return dll.size // Return the size of the doubly linked list.
}

