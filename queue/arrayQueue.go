package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(val int) {
	q.items = append(q.items, val)
}

func (q *Queue) Dequeue() int {
	val := q.items[0]

	q.items = q.items[1:]
	fmt.Println(q.items)
	return val
}

func main() {
	q := &Queue{}

	q.Enqueue(12)
	q.Enqueue(1)
	q.Enqueue(313)
	q.Enqueue(65)

	fmt.Println(q.items)

	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
}
