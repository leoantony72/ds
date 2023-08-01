//adjancey list
package main

import "fmt"

type Graph struct {
	vertices []*Vertex
}

//represent a graph vertex
type Vertex struct {
	key      int
	adjacent []*Vertex
}

//contain
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

//print vertices
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("%v", v.key)
		}
	}
}

//add vertex
func (g *Graph) addVertex(k int) {
	if contains(g.vertices, k) {
		return
	}
	g.vertices = append(g.vertices, &Vertex{key: k})
}

//add edges
func (g *Graph) addEdges(from, to int) {
	//get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	//check error - edge already exist
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	}else if contains(fromVertex.adjacent,to){
		err := fmt.Errorf("Existing edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	}else {
		//add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

func main() {
	g := &Graph{}

	for i := 0; i < 5; i++ {
		g.addVertex(i)
	}

	g.addEdges(1, 2)
	g.addEdges(1, 2)
	// g.addEdges(6, 2)
	g.Print()
}
