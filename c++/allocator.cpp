#include <iostream>
using namespace std;

int* allocate(int x);

int main() {
    int elements = 0;
    cout << "number of elements to allocate: "; cin >> elements;

    int* array = allocate(elements);

    //display new array
    for(int i=0; i<elements; i++) {
        cout << array[i];
    }
   
    return 0;
}

int* allocate(int x) {
    int* dynamicArray = new int[x];

    return dynamicArray;
}
