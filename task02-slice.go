package homework

func reverse(input []int64) (result []int64) {

	//Place your code here
	result = make([]int64, len(input))
	copy(result, input)

	iLimit := len(input) / 2
	for i := 0; i < iLimit; i++ {
		idx := len(input) - 1 - i
		result[i], result[idx] = result[idx], result[i]
	}

	return result
}
