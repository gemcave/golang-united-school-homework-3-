package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var sum float32

	for _, item := range input {
		sum += item
	}

	return sum / float32(len(input))
}
