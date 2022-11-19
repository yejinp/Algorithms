package selectionSort

type cmpFunc func(interface{}, interface{}) int

func SelectionSort(data []interface{}, cmp cmpFunc) {
	length := len(data)

	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if 0 < cmp(data[min], data[j]) {
				min = j
			}
		}
		data[i], data[min] = data[min], data[i]
	}
}
