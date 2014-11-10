package sort

func Bubble(list []int) []int {
	s := len(list)
	for s != 0 {
		n := 0
		for i := 1; i < s; i++ {
			if list[i-1] > list[i] {
				list[i-1], list[i] = list[i], list[i-1]
				n = i
			}
		}
		s = n
	}

	return list
}

func Selection(list []int) []int {
	for i := 1; i < len(list); i++ {
		for j := i; j < len(list); j++ {
			if list[j] < list[i-1] {
				list[i-1], list[j] = list[j], list[i-1]
			}
		}
	}

	return list
}

func merge(left []int, right []int) []int {
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

func Quick(list []int) {
	if len(list) <= 1 {
		return
	}

	pivot := list[0]

	i := 1

	for j := 1; j<len(list); j++ {
		if list[j] < pivot {
			list[i], list[j] = list[j], list[i]
			i++
		}
	}

	list[0], list[i-1] = list[i-1], list[0]

	Quick(list[:i-1])
	Quick(list[i:])
}
