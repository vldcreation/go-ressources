#include <stdio.h>
#include <stdlib.h>

int binary_search(int* array, int value, int low, int high) {
    int mid;
    if (high < low) {
        return -1; 
    } else {
        mid = low + (high - low) / 2;
        if (array[mid] > value) {
            return binary_search(array, value, low, mid - 1); 
        } else if (array[mid] < value) {
            return binary_search(array, value, mid + 1, high);
        } else {
            return mid;
        }
    }
}

int main() {
    FILE *inputFile = fopen("case2_c.in", "r");
    if (inputFile == NULL) {
        fprintf(stderr, "Error opening input file.\n");
        return 1;
    }

    FILE *outputFile = fopen("case2_c.out", "w");
    if (outputFile == NULL) {
        fprintf(stderr, "Error creating output file.\n");
        fclose(inputFile);
        return 1;
    }

    int i, value, answer, num_elements = 1;
    int array[10001]; 
    int query[10001];

    while (fscanf(inputFile, "%d", &value) > 0) {
        if (num_elements <= 10000) {
            printf("array [%d]%d\n", num_elements-10000, value);
            array[num_elements-1] = value;
        } else {
            printf("query [%d]%d\n", num_elements-10001, value);
            query[num_elements-10001] = value;
        }
        // if (num_elements >= 10000) {
        //     // fprintf(stderr, "Error: Too many elements in input.\n");
        //     fclose(inputFile);
        //     fclose(outputFile);
        //     // return 1;
        //     break;
        // }
        num_elements++;
    }

    printf("array[0]: %d\n", array[0]);
    printf("array[9999]: %d\n", array[9999]);
    printf("query[0]: %d\n", query[0]);
    printf("query[9999]: %d\n", query[9999]);
    // Rewind the input file to read search values 
    rewind(inputFile); 

    // Read values to search for and perform binary search
    for(int i=0; i<10000; i++){
        answer = binary_search(array, query[i], 0, 9999);
        fprintf(outputFile, "%d\n", answer);
        // if (answer == -1) {
        //     fprintf(outputFile, "-1\n");
        // }
    }

    fclose(inputFile);
    fclose(outputFile);
    return 0;
}