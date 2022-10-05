package mySort

type cmpFunc func(interface{}, interface{}) int

func QuickSort(slice []interface{}, cmp cmpFunc) {

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
