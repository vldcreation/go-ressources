#include <stdio.h>
#include <stdlib.h>
int binary_search(int* array, int value, int low, int high) {
    int mid;
    if (high < low) {
        return -1;
    } else {
        mid = low + (high - low)/2;
        if (array[mid] > value) {
            return binary_search(array, value, low, mid);
        } else if (array[mid] < value) {
            return binary_search(array, value, mid+1, high);
        } else {
            return mid;
        }
    }
}
main() {
    int i, value, answer;
    int array[10000];
    for (i=0; i<10000; i++) {
        scanf("%d", array+i);
    }
    for (i=0; i<10000; i++) {
        scanf("%d", &value);
        answer = binary_search(array, value, 0, 9999);
        printf("%d\n", answer);
    }
}
