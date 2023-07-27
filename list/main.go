package main

import "fmt"

type Node struct {
	data string
	link *Node
}
type Linkedlist struct {
	head   *Node
	tail   *Node
	length int
}

func (l *Linkedlist) prepend(val string) {
	node := &Node{data: val}
	temp := l.head
	l.head = node

	if l.length == 0 {
		l.tail = node
	}

	l.head.link = temp
	l.length++
}

func (l *Linkedlist) append(val string) {
	node := &Node{data: val}

	if l.length == 0 {
		l.head = node
	} else {
		l.tail.link = node
	}

	l.tail = node
	l.length++
}

func (l *Linkedlist) insert(val string, index int) {
	if index >= l.length {
		l.append(val)
		return
	}
	if l.length == 0 {
		l.prepend(val)
		return
	}

	node := &Node{data: val}
	curr := l.head
	for i := 0; i < l.length; i++ {
		if i == (index - 1) {
			node.link = curr.link

			curr.link = node
			break
		}
		curr = curr.link
	}
}

func (l *Linkedlist) delete(index int) {
	if index >= l.length {
		fmt.Println("Out of bound")
		return
	}

	if l.length == 1 {
		l.head = nil
		l.tail = nil
		l.length = 0
	}
	curr := l.head
	for i := 0; i < l.length; i++ {
		if i == (index - 1) {
			nextNode := curr.link

			curr.link = nextNode.link //curr.link becomes nil on tail

			nextNode.link = nil
			if index == (l.length-1) {
				l.tail = curr
			}
			break

		}
		curr = curr.link
	}
}

func main() {
	list := Linkedlist{}
	//@ add element at the front
	list.prepend("hi")
	list.prepend("bye")
	list.prepend("goodbye")

	getval(list.head)

	//@ add element at the end
	// list.append("i get it now")
	// fmt.Println("--------------")
	// getval(list.head)

	// fmt.Println("--------------")
	// //@ insert elements in between
	// list.insert("sike", 2)
	// getval(list.head)

	//@delete val
	list.delete(2)
	fmt.Println("--------------")
	getval(list.head)
	// fmt.Println(list.tail.data)
}

func getval(h *Node) {
	fmt.Println(h.data)
	if h.link == nil {
		return
	}
	getval(h.link)
}
