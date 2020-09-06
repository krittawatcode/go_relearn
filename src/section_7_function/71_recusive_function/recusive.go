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

/*
					a
				/	|	\
		b			a			k
	    /\			/\
	p		q	r		s
				|
				q

*/

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
	// outline(&root) // version I
	outline([]string{}, &root)

}

/* ----- version I ----- */
// func outline(n *Node) {
// 	if len(n.nodes) == 0 {
// 		fmt.Println(n.value)
// 	}
// 	// fmt.Println(n.value)
// 	for _, v := range n.nodes {
// 		// fmt.Println(v)
// 		outline(&v)
// 	}
// }

/* ----- version II ----- */
func outline(stack []string, n *Node) {
	stack = append(stack, n.value)
	if len(n.nodes) == 0 {
		// fmt.Println(n.value)
		fmt.Println(stack)
	}
	// fmt.Println(n.value)
	for _, v := range n.nodes {
		// fmt.Println(v)
		// outline(n.value, &v)
		outline(stack, &v)

	}
}
