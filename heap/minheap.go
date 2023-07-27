package main

import "fmt"

type Minheap struct {
	array []int
}

func main() {
	h := Minheap{}

	buildheap := []int{23, 12, 123, 436, 2, 14, 1}

	for _, v := range buildheap {
		h.Insert(v)
	}
	fmt.Println(h)
	fmt.Println(h.Extract())
	fmt.Println(h)



}

func (h *Minheap) Insert(val int) {
	h.array = append(h.array, val)

	h.Minheapifyup(len(h.array) - 1)
}

func (h *Minheap) Minheapifyup(i int) {
	for h.array[Parent(i)] > h.array[i] {
		h.swap(Parent(i), i)
		i = Parent(i)
	}
}

func (h *Minheap) Extract() int {
	val := h.array[0]
	l := len(h.array) - 1
	h.array[0] = h.array[l]
	h.array = h.array[:l]
	h.minheapifydown(0)
	return val
}

func (h *Minheap) minheapifydown(i int) {
	lastindex := len(h.array) - 1
	childtoCompare := 0
	l, r := left(i), right(i)

	for l <= lastindex {
		if l == lastindex {
			childtoCompare = l
		} else if h.array[l] < h.array[r] {
			childtoCompare = l
		} else {
			childtoCompare = r
		}

		if h.array[childtoCompare] < h.array[i] {
			h.swap(childtoCompare, i)
			i = childtoCompare
			l,r = left(i),right(r)
		}else{
			return
		}
	}
}

func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}

func (h *Minheap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func Parent(index int) int {
	return (index - 1) / 2
}
