package mincut

import (
	"testing"
	"strconv"
)

func testVerticeEq(a, b []*Vertice) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v.Name != b[i].Name {
			return false
		}
	}

	return true
}

func testEdgeEq(a, b []*Edge) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v.Head != b[i].Head || v.Tail != b[i].Tail {
			return false
		}
	}

	return true
}

func TestCreateGraph(t *testing.T) {
	lines := []string{
		"1 2 5 6",
		"2 1 3 5 6",
		"3 4 7 8",
		"4 3 7 8",
		"5 1 2 6",
		"6 1 2 5 7",
		"7 3 4 8",
		"8 3 4 7",
	}
	vertices := []*Vertice {
		&Vertice{"1"},
		&Vertice{"2"},
		&Vertice{"3"},
		&Vertice{"4"},
		&Vertice{"5"},
		&Vertice{"6"},
		&Vertice{"7"},
		&Vertice{"8"},
	}
	edges := []*Edge {
		&Edge{"2", "1"} ,
		&Edge{"4", "3"} ,
		&Edge{"5", "1"} ,
		&Edge{"5", "2"} ,
		&Edge{"6", "1"} ,
		&Edge{"6", "2"} ,
		&Edge{"6", "5"} ,
		&Edge{"7", "3"} ,
		&Edge{"7", "4"} ,
		&Edge{"8", "3"} ,
		&Edge{"8", "4"} ,
		&Edge{"8", "7"} ,
	}

	v, e := CreateGraph(lines)

	if !testVerticeEq(vertices, v) {
		t.Error(
			"Graph reading failed.",
			"Vertices are created incorrectly.",
		)
	}

	if !testEdgeEq(edges, e) {
		t.Error(
			"Graph reading failed.",
			"Edges are created incorrectly.",
		)
	}
}

func TestMincut(t *testing.T) {
	lines := []string{
		"1 2 3 5 6",
		"2 1 3 5 6",
		"3 2 4 7 8",
		"4 3 7 8",
		"5 1 2 6",
		"6 1 2 5 7",
		"7 3 4 6 8",
		"8 3 4 7",
	}

	min := 999
	for i:=0; i<1000; i++ {
		vertices, edges := CreateGraph(lines)
		tmp := MinCut(vertices, edges)
		if tmp < min {
			min = tmp
		}
	}

	if min != 2 {
		t.Error(
			"Minimum cut calculation failed.",
			"Expected 2.",
			"Got ", strconv.Itoa(min),
		)
	}
}
