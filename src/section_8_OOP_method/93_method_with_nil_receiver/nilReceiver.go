package main

import "fmt"

type BinaryTree struct {
	value       int
	left, right *BinaryTree
}

func (b *BinaryTree) sum() int {
	if b == nil {
		return 0
	}
	return b.value + b.left.sum() + b.right.sum()
}

func main() {

	// 2 3 5 1

	bt := BinaryTree{
		value: 2,
		left:  &BinaryTree{value: 3, left: &BinaryTree{value: 1, left: nil, right: nil}},
		right: &BinaryTree{value: 5, left: nil, right: nil},
	}
	fmt.Println(bt.sum()) // 11
}
