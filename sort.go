package sort

func Bubble(list []int) {
	s := len(list)
	for s!=0 {
		n := 0
		for i:=1; i<s; i++ {
			if list[i-1] > list[i] {
				list[i-1], list[i] = list[i], list[i-1]
				n = i
			}
		}
		s = n
	}
}
