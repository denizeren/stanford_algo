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
