package graph

// Graph implementation
type Graph struct {
	nodes []*Node
}

type Node struct {
	value     int
	visited   bool
	adjacents []*Node
}

func (n *Node) Equals(node *Node) bool {
	if n.value == node.value && n.visited == node.visited {
		for k, v := range n.adjacents {
			if v != node.adjacents[k] {
				return false
			}
		}
		return true
	}

	return false
}

// Queue implementation
type Queue struct {
	data []*Node
}

func (q *Queue) Push(node *Node) {
	q.data = append(q.data, node)
}

func (q *Queue) Pop() *Node {
	if len(q.data) < 1 {
		return nil
	}
	node := q.data[0]
	q.data = q.data[1:]

	return node
}

// Stack implementation
type Stack struct {
	data []*Node
}

func (s *Stack) Push(node *Node) {
	s.data = append(s.data, node)
}

func (s *Stack) Pop() *Node {
	if len(s.data) < 1 {
		return nil
	}
	node := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return node
}
