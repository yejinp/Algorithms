#include <stdlib.h>
#include <unistd.h>
#include <stdio.h>

int insertion(int array[], int size) 
{
	int i = 0, j, N = size;
	int min, temp;
	for(i = 1; i < N; i++) {
		min = i;
		for(j = i; j > 0 && array[j] < array[j-1]; j--) {
			temp = array[j];
			array[j] = array[j-1];
			array[j-1] = temp;
		}
	}

	return 0;
}

int print_array(int array[], int size)
{
	int i = 0; 
	for(i = 0; i < size; i ++) {
		printf("%5d ", array[i]);
	}

	printf("\n");
}

int main() {

	int array[] = {32, 67, 12, 9, 98, 78, 53, 0, 9};
	int size = sizeof(array)/sizeof(array[0]);
	print_array(array, size);
	insertion(array, size);
	print_array(array, size);
}
