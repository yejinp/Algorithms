package quickSort

type cmpFunc func(interface{}, interface{}) int

func partition(data []interface{}, left, right int, cmp cmpFunc) int {
	pivot := data[left]
	i, j := left, right

	for i < j {
		for i < j && cmp(pivot, data[j]) < 0 {
			j--
		}
		if i < j {
			data[i] = data[j]
			i++
		}

		for i < j && cmp(data[i], pivot) < 0 {
			i++
		}
		if i < j {
			data[j] = data[i]
			j--
		}
	}
	data[i] = pivot
	return i
}

func partition2(data []interface{}, left, right int, cmp cmpFunc) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if cmp(data[i], data[pivot]) < 0 {
			data[i], data[index] = data[index], data[i]
			index += 1
		}
	}
	data[pivot], data[index-1] = data[index-1], data[pivot]
	return index - 1
}

func qSort(data []interface{}, start, end int, cmp cmpFunc) {
	if start >= end {
		return
	}
	p := partition(data, start, end, cmp)
	qSort(data, start, p-1, cmp)
	qSort(data, p+1, end, cmp)
}

func QuickSort(data []interface{}, cmp cmpFunc) {
	qSort(data, 0, len(data)-1, cmp)
}
