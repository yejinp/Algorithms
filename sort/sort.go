package mySort

type cmpFunc func(interface{}, interface{}) int

func QuickSort(slice []interface{}, cmp cmpFunc) {

}

func InsertionSort(slice []interface{}, cmp cmpFunc) {

}

func merge1(data []interface{}, left, mid, right int, cmp cmpFunc) {
	help := make([]interface{}, right-left+1)
	i, middle := 0, mid+1
	for left <= mid && middle <= right {
		//	fmt.Println(left, middle, mid, right, data[left], data[middle])
		if cmp(data[left], data[middle]) < 0 {
			help[i] = data[left]
			left++
		} else {
			help[i] = data[middle]
			middle++
		}
		i++
	}

	if left <= mid {
		copy(help[i:], data[left:])
	} else {
		copy(help[i:], data[middle:])
	}
	copy(data[left:], help)
}

func merge(data []interface{}, left, mid, right int, cmp cmpFunc) {
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
