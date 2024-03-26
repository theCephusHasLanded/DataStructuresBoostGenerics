package go_data_structures

import (
	"golang.org/x/exp/constraints"
)

type MaxHeap[T constraints.Ordered] struct {
	items []T
}

func NewMaxHeap[T constraints.Ordered]() *MaxHeap[T] {
	return &MaxHeap[T]{}
}

func (h *MaxHeap[T]) Peek() T {
	if len(h.items) == 0 {
		var zero T
		return zero
	}
	return h.items[0]
}

func (h *MaxHeap[T]) RemoveTop() {
	if len(h.items) == 0 {
		return
	}
	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.siftDown(0)
}

func (h *MaxHeap[T]) Insert(val T) {
	h.items = append(h.items, val)
	h.siftUp(len(h.items) - 1)
}

func (h *MaxHeap[T]) Remove(val T) {
	for i, item := range h.items {
		if item == val {
			h.items[i] = h.items[len(h.items)-1]
			h.items = h.items[:len(h.items)-1]
			h.siftDown(i)
			return
		}
	}
}

func (h *MaxHeap[T]) Empty() bool {
	return len(h.items) == 0
}

func (h *MaxHeap[T]) Size() int {
	return len(h.items)
}

func (h *MaxHeap[T]) Sorted() []T {
	originalItems := make([]T, len(h.items))
	copy(originalItems, h.items)

	var sorted []T
	for len(h.items) > 0 {
		sorted = append(sorted, h.Peek())
		h.RemoveTop()
	}

	h.items = originalItems
	return sorted
}

func (h *MaxHeap[T]) siftUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.items[index] <= h.items[parentIndex] {
			break
		}
		h.items[index], h.items[parentIndex] = h.items[parentIndex], h.items[index]
		index = parentIndex
	}
}

func (h *MaxHeap[T]) siftDown(index int) {
	lastIndex := len(h.items) - 1
	for index < lastIndex {
		leftChildIndex := 2*index + 1
		rightChildIndex := 2*index + 2
		largestIndex := index

		if leftChildIndex <= lastIndex && h.items[leftChildIndex] > h.items[largestIndex] {
			largestIndex = leftChildIndex
		}
		if rightChildIndex <= lastIndex && h.items[rightChildIndex] > h.items[largestIndex] {
			largestIndex = rightChildIndex
		}
		if largestIndex == index {
			break
		}
		h.items[index], h.items[largestIndex] = h.items[largestIndex], h.items[index]
		index = largestIndex
	}
}
