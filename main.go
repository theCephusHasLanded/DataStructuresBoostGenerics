package main

import (
	"fmt"
	"/Users/cephus/Desktop/DataStructuresBoost/go_data_structures"

)

func main() {
	// Create a new BinarySearchTree instance.
    bst := NewBinarySearchTree[string, int]()
	fmt.Println(bst)
	// Test Insert.
	bst.Insert("apple", 5)
	fmt.Println("Inserted ('apple', 5)")

	bst.Insert("banana", 2)
	fmt.Println("Inserted ('banana', 2)")

	bst.Insert("orange", 8)
	fmt.Println("Inserted ('orange', 8)")

	// Test Contains.
	fmt.Printf("Contains 'apple'? %v\n", bst.Contains("apple"))
	fmt.Printf("Contains 'grape'? %v\n", bst.Contains("grape"))

	// Test InOrderTraversal.
	inOrder := bst.InOrderTraversal()
	fmt.Println("InOrder Traversal:", inOrder)

	// Test PreOrderTraversal.
	preOrder := bst.PreOrderTraversal()
	fmt.Println("PreOrder Traversal:", preOrder)

	// Test PostOrderTraversal.
	postOrder := bst.PostOrderTraversal()
	fmt.Println("PostOrder Traversal:", postOrder)

	// Test LevelOrderTraversal.
	levelOrder := bst.LevelOrderTraversal()
	fmt.Println("LevelOrder Traversal:", levelOrder)

	// Bonus: Test Remove (if implemented).
	bst.Remove("banana")
	fmt.Println("After Removing 'banana':")
	fmt.Println("InOrder Traversal:", bst.InOrderTraversal())
}
