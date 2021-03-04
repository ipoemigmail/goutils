package collection

func Range(from int, to int) []int {
	if from > to {
		return make([]int, 0)
	}
	r := make([]int, to-from+1)
	for i := 0; i < len(r); i++ {
		r[i] = from + i
	}
	return r
}

func RangeStep(from int, to int, step int) []int {
	if from > to {
		return make([]int, 0)
	}
	length := (to - from + 1) / step
	if (to-from+1)%step != 0 {
		length += 1
	}
	r := make([]int, length)
	for i := 0; i < len(r); i++ {
		r[i] = from + i*step
	}
	return r
}
