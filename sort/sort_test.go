package sort

import "testing"

type IntSlice []int

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
	{[]int{-10, 1, -10, 3, -2, 4, 0}, []int{-10, -10, -2, 0, 1, 3, 4}},
	{[]int{6, 7, 1, 2, 3}, []int{1, 2, 3, 6, 7}},
	{[]int{6, 7, 1, 2}, []int{1, 2, 6, 7}},
	{[]int{6, 1, 2}, []int{1, 2, 6}},
	{[]int{5, 6, 7, 1, 2, 3}, []int{1, 2, 3, 5, 6, 7}},
	{[]int{5, 6, 7, 1, 2, 3, 4}, []int{1, 2, 3, 4, 5, 6, 7}},
}

func (list IntSlice) Less(i, j int) bool {
	if list[i] < list[j] {
		return true
	} else {
		return false
	}
}

func (list IntSlice) Len() int {
	return len(list)
}

func (list IntSlice) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
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
		unsorted := make(IntSlice, len(pair.unsorted))
		copy(unsorted, pair.unsorted)
		Bubble(unsorted)
		if !testSliceEq(unsorted, pair.sorted) {
			t.Error(
				"For", pair.unsorted,
				"expected", pair.sorted,
				"got", unsorted,
			)
		}
	}
}

func TestSelection(t *testing.T) {
	for _, pair := range tests {
		unsorted := make(IntSlice, len(pair.unsorted))
		copy(unsorted, pair.unsorted)
		Selection(unsorted)
		if !testSliceEq(unsorted, pair.sorted) {
			t.Error(
				"For", pair.unsorted,
				"expected", pair.sorted,
				"got", unsorted,
			)
		}
	}
}

func TestMerge(t *testing.T) {
	for _, pair := range tests {
		unsorted := make([]int, len(pair.unsorted))
		copy(unsorted, pair.unsorted)
		v := Merge(unsorted)
		if !testSliceEq(v, pair.sorted) {
			t.Error(
				"For", pair.unsorted,
				"expected", pair.sorted,
				"got", v,
			)
		}
	}
}

func TestQuick(t *testing.T) {
	for _, pair := range tests {
		unsorted := make(IntSlice, len(pair.unsorted))
		copy(unsorted, pair.unsorted)
		Quick(unsorted, 0, unsorted.Len())
		if !testSliceEq(unsorted, pair.sorted) {
			t.Error(
				"For", pair.unsorted,
				"expected", pair.sorted,
				"got", unsorted,
			)
		}
	}
}
