package graph

// Graph implementation
type Graph struct {
	nodes []*Node
}

func (g *Graph) AddNode(n *Node) {
	g.nodes = append(g.nodes, n)
}

func (f *Graph) AddEdges(x *Node, y *Node) {
	x.AddAdjecent(y)
	y.AddAdjecent(x)
}

func (f *Graph) BFS(node *Node) {
	q := Queue{}
	q.Push(node)
	node.SetVisited()

	for len(q.data) > 0 {
		n := q.Pop()
		for _, x := range n.adjacents {
			if x.Visited() == false {
				x.SetVisited()
				q.Push(x)
			}
		}
	}
}

func (f *Graph) DFS(node *Node) {
	s := Stack{}
	s.Push(node)
	node.SetVisited()

	for len(s.data) > 0 {
		n := s.Pop()
		for _, x := range n.adjacents {
			if x.Visited() == false {
				x.SetVisited()
				s.Push(x)
			}
		}
	}
}

type Node struct {
	value     int
	visited   bool
	adjacents []*Node
}

func (n *Node) Equals(node *Node) bool {
	if n.value == node.value && n.visited == node.visited {
		// TODO check when order is different
		for k, v := range n.adjacents {
			if v != node.adjacents[k] {
				return false
			}
		}
		return true
	}

	return false
}

func (n *Node) AddAdjecent(node *Node) {
	n.adjacents = append(n.adjacents, node)
}

func (n *Node) Visited() bool {
	return n.visited
}

func (n *Node) SetVisited() {
	n.visited = true
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
