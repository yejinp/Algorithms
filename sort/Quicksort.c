#include <stdlib.h>
#include <unistd.h>
#include <stdio.h>

int partition(int array[], int lo, int hi) {
	int i, j, k, t, temp = array[lo];
	if(hi <= lo) {
		return lo;
	}
	
	j = hi + 1;
	i = lo;
	while(1) {
		while(array[++i] < temp) {
			if(i == hi ) {
				break;
			}
		}

		while(temp < array[--j]) {
			if(j == lo) {
				break;
			}
		}
		if(i >= j) {
			break;
		}
		t = array[i];
		array[i] = array[j];
		array[j] = t;
	}
	t = array[lo];
	array[lo] = array[j];
	array[j] = t;
	return j;
}

void sort(int array[], int lo, int hi) 
{
	int mid;
	if(hi <= lo) {
		return;
	}
	mid = partition(array, lo, hi);
	sort(array, lo, mid - 1);
	sort(array, mid + 1, hi);
}

int print_array(int array[], int size)
{
	int i = 0; 
	for(i = 0; i < size; i ++) {
		printf("%5d ", array[i]);
	}

	printf("\n");
	return 0;
}

int main() {

	int array[] = {32, 67, 12, 9, 98, 78, 53, 0, 9};
	int size = sizeof(array)/sizeof(array[0]);
	print_array(array, size);
	sort(array, 0, size - 1);
	print_array(array, size);
}
