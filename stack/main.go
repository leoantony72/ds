package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) push(val int) {
	s.items = append(s.items, val)
}

func (s *Stack) pop()int{
	l := len(s.items) - 1
	val := s.items[l]
	s.items = s.items[:l]
	return val
}

func main(){
	s := &Stack{}

	s.push(1)
	s.push(100)
	s.push(324)
	s.push(54)

	fmt.Println(s.items)
	
	s.pop()
	fmt.Println(s.items)
}
