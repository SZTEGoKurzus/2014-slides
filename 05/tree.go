package main

import "fmt"

// START1 OMIT
type Node struct {
	Value    int
	Children []*Node
}

func New(v int, children ...*Node) *Node {
	return &Node{Value: v, Children: children}
}

func main() {
	tree := New(0, New(1, New(2, New(3)), New(4), New(5)), New(6), New(7, New(8)))
	ch := make(chan int)
	go tree.Walk(ch)
	for i := range ch {
		fmt.Println(i)
	}
}

// START2 OMIT
func (n *Node) Walk(out chan<- int) {
	// END1 OMIT
	n.walk(out)
	close(out)
}

func (n *Node) walk(out chan<- int) {
	out <- n.Value
	for _, c := range n.Children {
		c.walk(out)
	}
}

// END2 OMIT
