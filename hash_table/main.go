package main

import (
	"fmt"
)

const ArraySize = 2

type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	val  string
	link *bucketNode
}

func main() {
	h := Init()

	h.insert("leo", "great1")
	h.insert("mariya", "great2")
	h.insert("geo", "great3")
	h.insert("test", "great4")
	h.insert("boi", "great5")
	h.insert("chicken", "great6")

	// h.Search("boi")
	
	// h.Search("boi")
	h.delete("chicken")
	h.Search("chicken")
	// fmt.Println(h.array[0].head.link)
}

//init
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

//Hashing algorithm
func (h *HashTable) hash(k string) int {
	sum := 0
	for _, v := range k {
		sum += int(v)
	}
	return sum % ArraySize
}

//insert into bucketnode
func (b *bucket) insert(k, val string) {
	newNode := &bucketNode{key: k, val: val}
	newNode.link = b.head
	b.head = newNode
}

//insert
func (h *HashTable) insert(key, val string) {
	hashedkey := h.hash(key)

	h.array[hashedkey].insert(key, val)
}

func (h *HashTable) Search(key string) {
	hashedkey := h.hash(key)

	currNode := h.array[hashedkey].head
	for currNode != nil {
		if currNode.key == key {
			fmt.Printf("Key: %s, Value: %s \n", currNode.key, currNode.val)
			return
		}
		currNode = currNode.link
	}
}
func (b *bucket) delete(key string) {
	// hashedkey := h.hash(key)
	if b.head.key ==key{
		b.head = b.head.link
		return
	}

	previousNode := b.head

	for previousNode.link != nil {
		if previousNode.link.key == key {
			previousNode.link = previousNode.link.link
			return
		}
		previousNode = previousNode.link
	}
}

func (h *HashTable) delete(key string) {
	hashedkey := h.hash(key)
	h.array[hashedkey].delete(key)
}
