package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	keys := make([]int, 0, len(input))

	for k, _ := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, i := range keys {
		result = append(result, input[i])
	}

	return result
}
