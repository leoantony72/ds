package main

import "fmt"

type Node struct {
	key   int
	left  *Node
	right *Node
}

//insert
func (b *Node) insert(k int) {
	if b.key < k {
		if b.right == nil {
			b.right = &Node{key: k}
		} else {
			b.right.insert(k)
		}
	} else if b.key > k {
		if b.left == nil {
			b.left = &Node{key: k}
		} else {
			b.left.insert(k)
		}
	}
}

func (b *Node) search(k int) bool {
	if b ==nil{
		return false
	}
	if b.key < k {
		//right
		return b.right.search(k)
	} else if b.key > k {
		//left
		return b.left.search(k)
	}
	return true
}

func (b *Node) traverse(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.key, n)
	b.traverse(n.left)
	b.traverse(n.right)
}

func main() {
	b := &Node{key: 10}

	b.insert(123215)
	b.insert(12)
	b.insert(121)
	b.insert(11)
	b.insert(9)

	b.traverse(b)

	fmt.Println(b.search(123215))
}
