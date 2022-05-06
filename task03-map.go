package homework

import (
	"sort"
)

func getSortMapKeys(input map[int]string) (result []int) {
	keys := make([]int, 0, len(input))

	for k := range input {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	return keys
}

func getSortMapValues(input map[int]string, keys []int) (values []string) {
	values = make([]string, 0, len(input))

	for sKey := range keys {
		k := keys[sKey]
		values = append(values, input[k])
	}

	return values
}

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	keys := getSortMapKeys(input)
	values := getSortMapValues(input, keys)

	return values
}
