package sort

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Bubble(data Interface) {
	s := data.Len()
	for s != 0 {
		n := 0
		for i := 1; i < s; i++ {
			if data.Less(i, i-1) {
				data.Swap(i, i-1)
				n = i
			}
		}
		s = n
	}
}

func Selection(data Interface) {
	for i := 1; i < data.Len(); i++ {
		for j := i; j < data.Len(); j++ {
			if data.Less(j, i-1) {
				data.Swap(i-1, j)
			}
		}
	}
}

func merge(left, right []int) []int {
	i := 0
	j := 0
	l := len(left)
	r := len(right)
	ret := make([]int, l+r)

	for i < l && j < r {
		if left[i] < right[j] {
			ret[i+j] = left[i]
			i++
		} else {
			ret[i+j] = right[j]
			j++
		}
	}

	if i < l {
		for i < l {
			ret[i+j] = left[i]
			i++
		}
	}

	if j < r {
		for j < r {
			ret[i+j] = right[j]
			j++
		}
	}

	return ret
}

func Merge(list []int) []int {
	s := len(list)

	if s <= 1 {
		return list
	}

	left := Merge(list[:s/2])
	right := Merge(list[s/2:])

	return merge(left, right)
}

func Quick(data Interface, s int, e int) {
	if e-s <= 1 {
		return
	}

	pivot := s

	i := s+1

	for j := s+1; j < e; j++ {
		if data.Less(j, pivot) {
			data.Swap(i, j)
			i++
		}
	}

	data.Swap(pivot, i-1)

	Quick(data, s, i-1)
	Quick(data, i, e)
}
