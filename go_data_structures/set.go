package go_data_structures

import (
	"encoding/json"
	"golang.org/x/exp/constraints"
	"hash/fnv"
)

type SetInterface[T constraints.Ordered] interface {
	hash(val T) int      // helper function to assign val to an index in the underlying array
	resize()             // helper function to double the size of the underlying array
	Contains(val T) bool // returns whether set contains element val. O(1)
	Insert(val T)        // inserts val in set if not present, increases size by 1. O(1)
	Remove(val T)        // removes item from set if present, decreases size by 1. O(1)
	Empty() bool         // returns whether set is empty. O(1)
	Size() int           // returns number of elements in set. O(1)
	Values() []T         // returns all values in the set.
}

type Set[T constraints.Ordered, V] struct {
	arr     []BinarySearchTree[T, V]
	maxFill float32
	size    int
}

func NewSet[T constraints.Ordered]() *Set[T] {
	initialSize := 8 // Starting size of the array
	return &Set[T]{
		arr:     make([]BinarySearchTree[T], initialSize),
		maxFill: 0.75, // Example fill factor before resizing
	}
}

func (s *Set[T]) hash(val T) int {
	h := fnv.New32a()
	// Convert val to bytes. Note: This simplistic approach might need refinement for complex types T.
	valBytes, _ := json.Marshal(val)
	h.Write(valBytes)
	// Use the hash value to find an index within the array bounds.
	return int(h.Sum32()) % len(s.arr)
}

func (s *Set[T]) resize() {
	// Double the size of the array
	newArr := make([]BinarySearchTree[T], len(s.arr)*2)
	// Rehash and insert all existing elements into the new array
	for _, tree := range s.arr {
		for _, val := range tree.Values() {
			newIndex := int(fnv.New32a().Sum32()) % len(newArr)
			newArr[newIndex].Insert(val)
		}
	}
	s.arr = newArr
}

func (s *Set[T]) Contains(val T) bool {
	index := s.hash(val)
	return s.arr[index].Contains(val)
}

func (s *Set[T]) Insert(val T) {
	if s.Contains(val) {
		return
	}
	index := s.hash(val)
	s.arr[index].Insert(val)
	s.size++
	if float32(s.size)/float32(len(s.arr)) > s.maxFill {
		s.resize()
	}
}

func (s *Set[T]) Remove(val T) {
	if !s.Contains(val) {
		return
	}
	index := s.hash(val)
	s.arr[index].Remove(val)
	s.size--
}

func (s *Set[T]) Empty() bool {
	return s.size == 0
}

func (s *Set[T]) Size() int {
	return s.size
}

func (s *Set[T]) Values() []T {
	var values []T
	for _, tree := range s.arr {
		values = append(values, tree.Values()...)
	}
	return values
}
