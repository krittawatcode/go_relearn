package main

import "fmt"

type BinaryTree struct {
	value int
	left  *BinaryTree
	right *BinaryTree
}

func main() {
	root := BinaryTree{value: 2}
	left := BinaryTree{value: 1}
	right := BinaryTree{value: 3}

	root.left = &left
	root.right = &right

	fmt.Println(root)

	showDF(&root)
}

func showDF(node *BinaryTree) {
	// fmt.Println(node.left.value)
	// fmt.Println(node.value)
	// fmt.Println(node.right.value)
	if node != nil {
		showDF(node.left)
		fmt.Println(node.value)
		showDF(node.right)
	}
}
