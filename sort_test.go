package sort

import "testing"

type testpair struct {
	unsorted []int
	sorted   []int
}

var tests = []testpair{
	{[]int{-3, -2, -1, 0, 1, 2, 3}, []int{-3, -2, -1, 0, 1, 2, 3}},
	{[]int{-3, -2, -2, 0, 1}, []int{-3, -2, -2, 0, 1}},
	{[]int{3, 2, 1, 0, -1, -2, -3}, []int{-3, -2, -1, 0, 1, 2, 3}},
	{[]int{3, 2, 1, 0, -2, -2, -3}, []int{-3, -2, -2, 0, 1, 2, 3}},
	{[]int{-10, 1, -10, 3, -2, 4}, []int{-10, -10, -2, 1, 3, 4}},
}

func testSliceEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestBubble(t *testing.T) {
	for _, pair := range tests {
		v := Bubble(pair.unsorted)
		if !testSliceEq(v, pair.sorted) {
			t.Error(
				"For", pair.unsorted,
				"expected", pair.sorted,
				"got", v,
			)
		}
	}
}

func TestSelection(t *testing.T) {
	for _, pair := range tests {
		v := Selection(pair.unsorted)
		if !testSliceEq(v, pair.sorted) {
			t.Error(
				"For", pair.unsorted,
				"expected", pair.sorted,
				"got", v,
			)
		}
	}
}

func TestMerge(t *testing.T) {
	for _, pair := range tests {
		v := Merge(pair.unsorted)
		if !testSliceEq(v, pair.sorted) {
			t.Error(
				"For", pair.unsorted,
				"expected", pair.sorted,
				"got", v,
			)
		}
	}
}
