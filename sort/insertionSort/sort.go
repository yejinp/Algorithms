package insertionSort

type cmpFunc func(interface{}, interface{}) int

func InsertionSort(data []interface{}, cmp cmpFunc) {

	n := len(data)
	for i := 1; i < n; i++ {
		key := data[i]
		j := i - 1
		for 0 <= j && cmp(key, data[j]) < 0 {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = key
	}
}
