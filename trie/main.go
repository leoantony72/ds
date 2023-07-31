package main

import "fmt"

const AlphabeticSize = 26

type Node struct {
	char  [26]*Node
	isEnd bool
}

type Trie struct {
	root *Node
}

func newTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

//insert
//foo
func (t *Trie) insert(w string) {
	length := len(w)
	if length == 0 {
		return
	}
	currnode := t.root
	for i := 0; i < length; i++ {
		charIndex := w[i] - 'a'
		if currnode.char[charIndex] == nil {
			currnode.char[charIndex] = &Node{}
		}
		currnode = currnode.char[charIndex]
	}
	currnode.isEnd = true
}

func (t *Trie) search(w string) bool {
	length := len(w)
	currnode := t.root
	for i := 0; i < length; i++ {
		charIndex := w[i] - 'a'
		if currnode.char[charIndex] == nil {
			return false
		}
		currnode = currnode.char[charIndex]
	}
	if currnode.isEnd == true {
		return true
	}
	return false
}

func main() {
	//initialize new trie
	t := newTrie()
	t.insert("john")
	t.insert("david")
	t.insert("leo")
	t.insert("zoro")

	fmt.Println(t.root)
	fmt.Println(t.search("kinh"))
}
