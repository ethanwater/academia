#include <iostream>
using namespace std;

int* expand(int array[], int size);

int main() {
    int array[4] = {1, 4, 7, 8};
    int size = sizeof(array)/sizeof(int);

    int* expandedArray = expand(array, size);
}

int* expand(int array[], int size) {
    int newSize = size*2;

    int* newArray = new int[newSize];
    for(int i=0; i<size; i++) {
        newArray[i] = array[i];
    }
    
    //for testing purposes to display the new expanded array
    //
    //for(int i=0; i<newSize; i++) {
    //    cout << newArray[i];
    //}

    return newArray;
}
