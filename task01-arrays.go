package homework

func average(input [15]float32) (result float32) {
	const inputLen = len(input)
	var inputSum float64 = 0
	for _, inputElem := range input {
		inputSum = inputSum + float64(inputElem)
	}

	var averageVal = float32(inputSum / float64(inputLen))

	//Place your code here
	return averageVal
}
