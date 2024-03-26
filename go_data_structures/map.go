package go_data_structures

import (
    "encoding/json"
    "fmt"
    "golang.org/x/exp/constraints"
    "hash/fnv"
)

// Pair holds a key-value pair.
// MapInterface defines operations for a map data structure.
type MapInterface[T constraints.Ordered, V any] interface {
    Contains(key T) bool
    Get(key T) (V, error)
    Set(key T, val V)
    Removes(key T)
    Empty() bool
    Size() int
    Values() []V
    Keys() []T
    Objects() []Pair[T, V]
}


// Map implements a map using a slice of binary search trees for collision resolution.
type Map [T constraints.Ordered, V any] struct {
    arr     []*BinarySearchTree[Pair[T, V]] // Use a slice of pointers to BinarySearchTree
    maxFill float32
    size    int
}

// NewMap creates a new Map instance.
func NewMap[T constraints.Ordered, V any](size int, maxFill float32) *Map[T, V] {
    return &Map[T, V]{
		arr:     make([]*BinarySearchTree[Pair[T, V]], size),
        maxFill: maxFill,
    }
}

// hash generates a hash code for the given key.
func (m *Map[T, V]) hash(key T) int {
    keyStr, _ := json.Marshal(key)
    h := fnv.New32a()
    h.Write(keyStr)
    return int(h.Sum32()) % len(m.arr)
}

// resize doubles the size of the underlying array and rehashes all keys.
func (m *Map[T, V]) resize() {
    newArrSize := len(m.arr) * 2
    newArr := make([]BinarySearchTree[Pair[T, V]], newArrSize)
    for _, tree := range m.arr {
        for _, pair := range tree.Objects() {
            newIndex := m.hash(pair.Key)
            newArr[newIndex].Insert(pair)
        }
    }
    m.arr = newArr
}

// Contains checks if the map contains the specified key.
func (m *Map[T, V]) Contains(key T) bool {
    index := m.hash(key)
    return m.arr[index].Contains(Pair[T, V]{Key: key})
}

// Get retrieves the value associated with the given key.
func (m *Map[T, V]) Get(key T) (V, error) {
    index := m.hash(key)
    if val, found := m.arr[index].Get(Pair[T, V]{Key: key}); found {
        return val.Value, nil
    }
    var zero V
    return zero, fmt.Errorf("key not found")
}

// Set inserts or updates the key-val pair in the map.
func (m *Map[T, V]) Set(key T, val V) {
    index := m.hash(key)
    if !m.arr[index].Contains(Pair[T, V]{Key: key}) {
        m.size++
        if float32(m.size)/float32(len(m.arr)) > m.maxFill {
            m.resize()
        }
    }
    m.arr[index].Insert(Pair[T, V]{Key: key, Value: val})
}

// Removes deletes the key-val pair from the map.
func (m *Map[T, V]) Removes(key T) {
    index := m.hash(key)
    if m.arr[index].Remove(Pair[T, V]{Key: key}) {
        m.size--
    }
}

// Empty checks if the map is empty.
func (m *Map[T, V]) Empty() bool {
    return m.size == 0
}

// Size returns the number of elements in the map.
func (m *Map[T, V]) Size() int {
    return m.size
}

// Values returns all values in the map.
func (m *Map[T, V]) Values() []V {
    var values []V
    for _, tree := range m.arr {
        for _, pair := range tree.Objects() {
            values = append(values, pair.Value)
        }
    }
    return values
}

// Keys returns all keys in the map.
func (m *Map[T, V]) Keys() []T {
    var keys []T
    for _, tree := range m.arr {
        for _, pair := range tree.Objects() {
            keys = append(keys, pair.Key)
        }
    }
    return keys
}

// Objects returns all key-value pairs in the map.
func (m *Map[T, V]) Objects() []Pair[T, V] {
    var objects []Pair[T, V]
    for _, tree := range m.arr {
        objects = append(objects, tree.Objects()...)
    }
    return objects
}

// BinarySearchTree and other necessary types or methods (such as `Insert`, `Remove`, `Contains`, `Objects` in BST) should be defined here.
