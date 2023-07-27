package main

import "fmt"

type Maxheap struct {
	array []int
}

func main() {
	heap := Maxheap{}
	fmt.Println(heap)

	buildheap := []int{432,75,213,70,20,4,30,10,32,12,700}

	for _, v := range buildheap {
		heap.Insert(v)
	}
	fmt.Println(heap)
	fmt.Println(heap.Extract())
	fmt.Println(heap)
	
}

func (h *Maxheap) Insert(val int) {
	h.array = append(h.array, val)
	h.maxHeapifyup(len(h.array) - 1)

}

func (h *Maxheap) Extract() int{
	val := h.array[0]

	if len(h.array) == 0 {
		fmt.Println("heap is empty")
	}
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]

	h.maxHeapifyDown(0)
	return val
}

func (h *Maxheap) maxHeapifyDown(index int) {
	lastindex := len(h.array) - 1
	l, r := leftchild(index), rightchild(index)
	childtocompare := 0
	for l <= lastindex {
		if l == lastindex {
			childtocompare = l
		} else if h.array[l] > h.array[r] {
			childtocompare = l
		} else {
			childtocompare = r
		}
		if h.array[childtocompare] > h.array[index] {
			h.swap(childtocompare, index)
			index = childtocompare
			l, r = leftchild(index), rightchild(index)
		} else {
			return
		}
	}

}

func (h *Maxheap) maxHeapifyup(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *Maxheap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func parent(index int) int {
	return (index - 1) / 2
}
func leftchild(i int) int {
	return 2*i + 1
}
func rightchild(i int) int {
	return 2*i + 2
}
