package bubbleSort

type cmpFunc func(interface{}, interface{}) int

func BubbleSort(data []interface{}, cmp cmpFunc) {
	length := len(data)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if 0 < cmp(data[j], data[j+1]) {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}
