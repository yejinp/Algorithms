#include <stdlib.h>
#include <unistd.h>
#include <stdio.h>

int Shell(int array[], int size) 
{
	int i = 0, j, N = size;
	int h, temp;

	while( h < N/3)
		h = h*3 + 1;

	while(  h >= 1) {
		for(i = h; i < N; i++) {
			for(j = i; j >= h && array[j] < array[j-1]; j-= h) {
				temp = array[j];
				array[j] = array[j-h];
				array[j-h] = temp;
			}
		}
		h = h / 3;
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
	Shell(array, size);
	print_array(array, size);
}
