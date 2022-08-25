package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func count() func() int {
	counter := 0
	return func() int {
		counter++
		return counter - 1
	}
}

func PreOrder(tree *Node) {
	if tree == nil {
		return
	}

	fmt.Println(tree.val)
	PreOrder(tree.left)
	PreOrder(tree.right)
}

func InOrder(tree *Node) {
	if tree == nil {
		return
	}
	InOrder(tree.left)
	fmt.Println(tree.val)
	InOrder(tree.right)
}

func PostOrder(tree *Node) {
	if tree == nil {
		return
	}

	PostOrder(tree.left)
	PostOrder(tree.right)
	fmt.Println(tree.val)
}

func main() {
	getCount := count()
	root := Node{val: getCount()}

	root.left = &Node{val: getCount()}
	root.right = &Node{val: getCount()}

	PreOrder(&root) // 0 1 2
	fmt.Println("---")
	InOrder(&root) // 1 0 2
	fmt.Println("---")
	PostOrder(&root) // 1 2 0
	fmt.Println("---")

}
