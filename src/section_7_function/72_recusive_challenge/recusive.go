package main

import "fmt"

// func countdown(c int) {
// 	fmt.Println(c)
// 	if c == 0 {
// 		return
// 	}
// 	countdown(c - 1)
// }

type Node struct {
	value string
	nodes []Node
}

func main() {

	// fmt.Println(time.Now().Clock())
	// countdown(5)
	p := Node{value: "p"}
	g := Node{value: "g"}
	b := Node{value: "b", nodes: []Node{p, g}}
	q := Node{value: "q"}
	s := Node{value: "s"}
	k := Node{value: "k"}
	r := Node{value: "r", nodes: []Node{q}}
	a := Node{value: "a", nodes: []Node{r, s}}
	root := Node{value: "a", nodes: []Node{b, a, k}}
	fmt.Println(root)
	result := outline([]string{}, &root)
	fmt.Println(result)

}

func outline(stack []string, n *Node) [][]string {
	stack = append(stack, n.value)
	result := [][]string{}
	if len(n.nodes) == 0 {
		// fmt.Println(stack)
		result = append(result, stack)
	}
	for _, v := range n.nodes {
		result = append(result, outline(stack, &v)...)
	}
	return result
}
