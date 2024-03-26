package go_data_structures

import (
	"golang.org/x/exp/constraints"
)

// Pair holds a key-value pair.
// Pair has been defined in a previous file!

// BinaryTreeNode represents a node in the binary search tree.
type BinaryTreeNode[T constraints.Ordered, V any] struct {
    Data  Pair[T, V]       // Corrected to store a Pair[T, V] directly
    Left  *BinaryTreeNode[T, V]
    Right *BinaryTreeNode[T, V]
}

// BinaryTreeInterface defines the operations for a binary tree.
type BinaryTreeInterface[T constraints.Ordered, V any] interface {
    Root() *BinaryTreeNode[T, V]
    Insert(key T, val V)
    Contains(key T) bool
    Remove(key T)
    Empty() bool
    InOrderTraversal() []Pair[T, V] // Corrected return type
    PreOrderTraversal() []Pair[T, V] // Corrected return type
    PostOrderTraversal() []Pair[T, V] // Corrected return type
    LevelOrderTraversal() []Pair[T, V] // Corrected return type
}

// BinarySearchTree implements a binary search tree.
type BinarySearchTree[T constraints.Ordered, V any] struct {
	root *BinaryTreeNode[T, V]
	size int
}

// NewBinarySearchTree creates a new instance of a binary search tree.
func NewBinarySearchTree[T constraints.Ordered, V any]() *BinarySearchTree[T, V] {
	return &BinarySearchTree[T, V]{}
}

// Insert inserts a new key-value pair into the binary search tree.
func (bst *BinarySearchTree[T, V]) Insert(key T, val V) {
    if bst.root == nil {
        bst.root = &BinaryTreeNode[T, V]{Data: Pair[T, V]{Key: key, Value: val}}
        bst.size++
        return
    }
    bst.root = insertNode(bst.root, Pair[T, V]{Key: key, Value: val})
    bst.size++
}


func insertNode[T constraints.Ordered, V any](node *BinaryTreeNode[T, V], data Pair[T, V]) *BinaryTreeNode[T, V] {

	if node == nil {
        return &BinaryTreeNode[T, V]{Data: data}
	}
	if data.Key < node.Data.Key {
        node.Left = insertNode(node.Left, data)
    } else if data.Key > node.Data.Key {
        node.Right = insertNode(node.Right, data)
	}
	// Duplicate keys are not inserted in this implementation.
	return node
}

// Contains checks if a key exists in the binary search tree.
func (bst *BinarySearchTree[T, V]) Contains(key T) bool {
	return containsNode(bst.root, key)
}

func containsNode[T constraints.Ordered, V any](node *BinaryTreeNode[T, V], key T) bool {
    if node == nil {
        return false
    }
    if key == node.Data.Key {
        return true
    } else if key < node.Data.Key {
        return containsNode(node.Left, key)
    } else {
        return containsNode(node.Right, key)
    }
}


func removeNode[T constraints.Ordered, V any](node *BinaryTreeNode[T, V], key T) *BinaryTreeNode[T, V] {
    if node == nil {
        return nil
    }
    if key < node.Data.Key {
        node.Left = removeNode(node.Left, key)
    } else if key > node.Data.Key {
        node.Right = removeNode(node.Right, key)
    } else {
        // Node with only one child or no child
        if node.Left == nil {
            return node.Right
        } else if node.Right == nil {
            return node.Left
        }
        // Node with two children: Get the inorder successor (smallest in the right subtree)
        minNode := findMin(node.Right)
        node.Data.Key, node.Data.Value = minNode.Data.Key, minNode.Data.Value
        node.Right = removeNode(node.Right, minNode.Data.Key)
    }
    return node
}

func findMin[T constraints.Ordered, V any](node *BinaryTreeNode[T, V]) *BinaryTreeNode[T, V] {
    current := node
    for current.Left != nil {
        current = current.Left
    }
    return current
}


// Traversal helper function
func traverse[T constraints.Ordered, V any](node *BinaryTreeNode[T, V], order string, result *[]Pair[T, V]) {
    if node == nil {
        return
    }
    switch order {
    case "inorder":
        traverse(node.Left, order, result)
        *result = append(*result, node.Data) // node.Data is already a Pair[T, V]
        traverse(node.Right, order, result)
    case "preorder":
        *result = append(*result, node.Data) // Preorder: Node -> Left -> Right
        traverse(node.Left, order, result)
        traverse(node.Right, order, result)
    case "postorder":
        traverse(node.Left, order, result) // Postorder: Left -> Right -> Node
        traverse(node.Right, order, result)
        *result = append(*result, node.Data)
    }
}


// InOrderTraversal returns an in-order traversal of the BST as a slice of Pair[T, V].
func (bst *BinarySearchTree[T, V]) InOrderTraversal() []Pair[T, V] {
    var result []Pair[T, V]
    traverse(bst.root, "inorder", &result)
    return result
}

// PreOrderTraversal returns a pre-order traversal of the BST as a slice of Pair[T, V].
func (bst *BinarySearchTree[T, V]) PreOrderTraversal() []Pair[T, V] {
    var result []Pair[T, V]
    traverse(bst.root, "preorder", &result)
    return result
}

// PostOrderTraversal returns a post-order traversal of the BST as a slice of Pair[T, V].
func (bst *BinarySearchTree[T, V]) PostOrderTraversal() []Pair[T, V] {
    var result []Pair[T, V]
    traverse(bst.root, "postorder", &result)
    return result
}

// LevelOrderTraversal adapted to handle Pair[T, V].
func (bst *BinarySearchTree[T, V]) LevelOrderTraversal() []Pair[T, V] {
    var result []Pair[T, V]
    if bst.root == nil {
        return result
    }
    queue := []*BinaryTreeNode[T, V]{bst.root}
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        result = append(result, current.Data) // Directly use Data which is Pair[T, V]
        if current.Left != nil {
            queue = append(queue, current.Left)
        }
        if current.Right != nil {
            queue = append(queue, current.Right)
        }
    }
    return result
}
