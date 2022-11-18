package heapSort

type cmpFunc func(interface{}, interface{}) int

func maxHeapify(data []interface{}, i int, cmp cmpFunc) {
	left, right := i*2, i*2+1
	largest := i

	if left < len(data) && cmp(data[largest], data[left]) < 0 {
		largest = left
	}
	if right < len(data) && cmp(data[largest], data[right]) < 0 {
		largest = right
	}

	if largest != i {
		data[largest], data[i] = data[i], data[largest]
		maxHeapify(data, largest, cmp)
	}
}

func buildMaxHeapify(data []interface{}, cmp cmpFunc) {
	for i := len(data) / 2; 0 <= i; i-- {
		maxHeapify(data, i, cmp)
	}
}

func HeapSort(data []interface{}, cmp cmpFunc) {
	buildMaxHeapify(data, cmp)
	for i := 1; i < len(data); i++ {
		data[len(data)-i], data[0] = data[0], data[len(data)-i]
		maxHeapify(data[:len(data)-i], 0, cmp)
	}
}
