package mySort

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

func InsertionSort(slice []interface{}, cmp cmpFunc) {

}

func merge(data []interface{}, left, mid, right int, cmp cmpFunc) {
	help := make([]interface{}, right-left+1)
	sleft, sright, i := left, mid+1, 0

	for ; sleft <= mid && sright <= right; i++ {
		if cmp(data[sleft], data[sright]) < 0 {
			help[i] = data[sleft]
			sleft++
		} else {
			help[i] = data[sright]
			sright++
		}
	}

	if sleft <= mid {
		copy(help[i:], data[sleft:])
	} else {
		copy(help[i:], data[sright:])
	}
	copy(data[left:], help)
}

func merge2(data []interface{}, left, mid, right int, cmp cmpFunc) {
	help := make([]interface{}, len(data))
	copy(help, data)

	i := left
	j := mid + 1
	for k := left; k <= right; k++ {
		if i > mid {
			data[k] = help[j]
			j++
		} else if j > right {
			data[k] = help[i]
			i++
		} else if cmp(help[j], help[i]) < 0 {
			data[k] = help[j]
			j++
		} else {
			data[k] = help[i]
			i++
		}

	}
}

func mSort(data []interface{}, left, right int, cmp cmpFunc) {
	if left == right {
		return
	}
	mid := (left + right) / 2
	mSort(data, left, mid, cmp)
	mSort(data, mid+1, right, cmp)
	merge(data, left, mid, right, cmp)
}

func MergeSort(data []interface{}, cmp cmpFunc) {
	mSort(data, 0, len(data)-1, cmp)
}

func SelectSort(slice []interface{}, cmp cmpFunc) {

}

func ShellSort(slice []interface{}, cmp cmpFunc) {

}

func HeapSort(slice []interface{}, cmp cmpFunc) {

}
