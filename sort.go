package sort

import "fmt"

func Bubble(list []int) {
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
}

func Selection(list []int) {
	for i := 1; i < len(list); i++ {
		for j := i; j < len(list); j++ {
			if list[j] < list[i-1] {
				list[i-1], list[j] = list[j], list[i-1]
			}
		}
	}
}
