package main

import (
	"math/rand"
	"strings"
	"bufio"
	"fmt"
	"os"
    "time"
)

type Vertice struct {
	Name string
}

type Edge struct {
	Tail string
	Head string
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// Creates graph from input slice
// First element of slice is vertice name rest are adjent vertices
// returns graph as slice of vertices and slice of edges
func CreateGraph(lines []string) ([]*Vertice, []*Edge) {
	var vertices []*Vertice
	var edges []*Edge

	addedVertices := make(map[string]int)

	for _, line := range lines {
		words := strings.Fields(line)

		v := new(Vertice)
		v.Name = words[0]

		vertices = append(vertices, v)
		addedVertices[v.Name] = 1

		for i:=1; i<len(words); i++ {
			// Since graph is unidirected check it
			// TODO convert vertices slice to map that would be a smarter move
			_, ok := addedVertices[words[i]]
			if !ok {
				continue
			}

			e := new(Edge)

			e.Tail = v.Name
			e.Head = words[i]
			edges = append(edges, e)
		}
	}

	return vertices, edges
}

func random(length int) int {
	rand.Seed(time.Now().Unix()+rand.Int63n(1000))
    return rand.Intn(length)
}

func contract(randEdgeIndex int, vertices []*Vertice, edges []*Edge) ([]*Vertice, []*Edge) {
	// Get randomly selected edge
	randEdge := edges[randEdgeIndex]

	tailVerticeIndex := 0
	tailVerticeValue := randEdge.Tail
	for i, v := range vertices {
		if v.Name == tailVerticeValue {
			tailVerticeIndex = i
			break
		}
	}

	// Discard contracted vertice
	vertices[tailVerticeIndex], vertices[len(vertices)-1] = vertices[len(vertices)-1], vertices[tailVerticeIndex]
	vertices = vertices[:len(vertices)-1]

	// Replace value of tail vertice in edges with the value of head vertice
	for i, _ := range edges {
		if edges[i].Tail == tailVerticeValue {
			edges[i].Tail = randEdge.Head
		}
		if edges[i].Head == tailVerticeValue {
			edges[i].Head = randEdge.Head
		}
	}

	// Remove self-loops
	edgeLen := len(edges)
	for i := 0; i < edgeLen; {
		if edges[i].Head == edges[i].Tail {
			edges[i], edges[edgeLen-1] = edges[edgeLen-1], edges[i]
			edgeLen--
		} else {
			i++
		}
	}
	edges = edges[:edgeLen]

	return vertices, edges
}

func MinCut(vertices []*Vertice, edges []*Edge) int {
	for len(vertices) > 2 {
		// Randomly select edge
		randEdgeIndex := random(len(edges)-1)
		vertices, edges = contract(randEdgeIndex, vertices, edges)
	}

	return len(edges)
}

func main() {
	lines, _ := ReadLines("/tmp/kargerMinCut.txt")

	min := 999
	for i:=0; i<1000; i++ {
		vertices, edges := CreateGraph(lines)
		tmp := MinCut(vertices, edges)
		if tmp < min {
			min = tmp
		}
	}

	fmt.Println(min)
}
