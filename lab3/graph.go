package main

import (
	"fmt"
	"math/rand"
	"slices"
)

type Node struct {
	ID        int
	nodes_in  []int
	nodes_out []int
}

func (n *Node) addInConn(ID int) {
	n.nodes_in = append(n.nodes_in, ID)
}

func (n *Node) addOutConn(ID int) {
	n.nodes_out = append(n.nodes_out, ID)
}

func (n Node) printNode() {
	fmt.Println("\tnode[", n.ID, "]:")
	fmt.Println("\t\tin:", n.nodes_in)
	fmt.Println("\t\tout:", n.nodes_out)
}

type Graph struct {
	nodes []*Node
}

func (g *Graph) addNode(n *Node) {
	g.nodes = append(g.nodes, n)
}

func (g *Graph) initGraph(N int) {
	for i := 0; i < N; i++ {
		n := Node{i, []int{}, []int{}}
		g.addNode(&n)
	}
	p := float64(4) / float64(N)

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			curr_p := rand.Float64()
			if curr_p <= p {
				addEdge(g.nodes[i], g.nodes[j])
			}
		}
	}

}

func (g Graph) printGraph() {
	fmt.Println("Graph:")
	for i := range g.nodes {
		g.nodes[i].printNode()
	}
	fmt.Println()
}

func (g Graph) getGraphDegreeInMap() map[int][]int {
	mIn := make(map[int][]int)
	for i := range g.nodes {
		deg := len(g.nodes[i].nodes_in)
		mIn[deg] = append(mIn[deg], i)
	}
	return mIn
}
func (g Graph) getGraphDegreeOutMap() map[int][]int {
	mOut := make(map[int][]int)
	for i := range g.nodes {
		deg := len(g.nodes[i].nodes_out)
		mOut[deg] = append(mOut[deg], i)
	}
	return mOut
}

func (g Graph) Floyd_Warshall_alg() [][]int {
	d := make([][]int, len(g.nodes), len(g.nodes))
	for i := range d {
		d[i] = make([]int, len(g.nodes), len(g.nodes))
	}
	for i := range g.nodes {
		for j := range g.nodes {
			if i == j {
				d[i][j] = 0
			} else {
				if slices.Contains(g.nodes[i].nodes_in, j) {
					d[i][j] = 1
				} else {
					d[i][j] = 1000
				}
			}
		}
	}

	for k := range g.nodes {
		for i := range g.nodes {
			for j := range g.nodes {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}
	return d
}

func addEdge(n_in *Node, n_out *Node) {
	n_in.addInConn(n_out.ID)
	n_out.addOutConn(n_in.ID)
}
