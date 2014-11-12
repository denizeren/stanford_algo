package graph

import (
	"strconv"
	"testing"
)

func TestQueue(t *testing.T) {
	nodes := []*Node{
		&Node{1, false, nil},
		&Node{2, false, nil},
		&Node{3, false, nil},
		&Node{4, false, nil},
	}

	q := Queue{}

	for _, v := range nodes {
		q.Push(v)
	}

	for _, v := range nodes {
		n := q.Pop()

		if n.Equals(v) == false {
			t.Error(
				"Push function of queue failed.",
				"Expected node: ", strconv.Itoa(v.value),
				"Got node: ", strconv.Itoa(n.value),
			)
		}
	}
}

func TestStack(t *testing.T) {
	nodes := []*Node{
		&Node{1, false, nil},
		&Node{2, false, nil},
		&Node{3, false, nil},
	}

	s := Stack{}

	for _, v := range nodes {
		s.Push(v)
	}

	for i := len(nodes) - 1; i >= 0; i-- {
		n := s.Pop()

		if n.Equals(nodes[i]) == false {
			t.Error(
				"Push function of queue failed.",
				"Expected node: ", strconv.Itoa(nodes[i].value),
				"Got node: ", strconv.Itoa(n.value),
			)
		}
	}
}

func TestBFS(t *testing.T) {
	nodes := []*Node{
		&Node{1, false, nil},
		&Node{2, false, nil},
		&Node{3, false, nil},
		&Node{4, false, nil},
		&Node{5, false, nil},
		&Node{6, false, nil},
	}

	g := Graph{}

	for _, v := range nodes {
		g.AddNode(v)
	}

	g.AddEdges(nodes[0], nodes[1])
	g.AddEdges(nodes[0], nodes[2])
	g.AddEdges(nodes[1], nodes[2])
	g.AddEdges(nodes[1], nodes[3])
	g.AddEdges(nodes[2], nodes[3])
	g.AddEdges(nodes[2], nodes[4])
	g.AddEdges(nodes[3], nodes[5])
	g.AddEdges(nodes[4], nodes[5])

	g.BFS(nodes[0])

	for _, v := range nodes {
		if v.Visited() == false {
			t.Error(
				"BFS failed.",
				"Expected node ", strconv.Itoa(v.value),
				" to be visited, but it's not visited.",
			)
		}
	}
}

func TestDFS(t *testing.T) {
	nodes := []*Node{
		&Node{1, false, nil},
		&Node{2, false, nil},
		&Node{3, false, nil},
		&Node{4, false, nil},
		&Node{5, false, nil},
		&Node{6, false, nil},
	}

	g := Graph{}

	for _, v := range nodes {
		g.AddNode(v)
	}

	g.AddEdges(nodes[0], nodes[1])
	g.AddEdges(nodes[0], nodes[2])
	g.AddEdges(nodes[1], nodes[2])
	g.AddEdges(nodes[1], nodes[3])
	g.AddEdges(nodes[2], nodes[3])
	g.AddEdges(nodes[2], nodes[4])
	g.AddEdges(nodes[3], nodes[5])
	g.AddEdges(nodes[4], nodes[5])

	g.DFS(nodes[0])

	for _, v := range nodes {
		if v.Visited() == false {
			t.Error(
				"DFS failed.",
				"Expected node ", strconv.Itoa(v.value),
				" to be visited, but it's not visited.",
			)
		}
	}
}
