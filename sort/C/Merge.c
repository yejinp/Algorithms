#include <stdlib.h>
#include <unistd.h>
#include <stdio.h>

int b[] = {0, 0, 0, 0, 0, 0, 0, 0, 0};

int merge(int array[], int lo, int mid, int hi) {
	int i, j, k;
	for(i = lo; i <= hi; i++)
		b[i] = array[i];
	
	i = lo;
	j = mid + 1;
	for( k = lo; k <= hi; k++) {
		if(i > mid) {
			array[k] = b[j++];
		}else if(j > hi) {
			array[k] = b[i++];
		}else if(b[j] < b[i]) {
			array[k] = b[j++];
		}else {
			array[k] = b[i++];
		}

	}
	return 0;
}

void sort(int array[], int lo, int up) 
{
	int mid = (lo + up) / 2;
	if(up <= lo) {
		return;
	}

	sort(array, lo, mid);
	sort(array, mid + 1, up);
	merge(array, lo, mid, up);

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
