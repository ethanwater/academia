#include <iostream>
#include <fstream>
#include <string>
#include <vector>
using namespace std;

double high(vector<double>);
double low(vector<double>);

int main() {
    ifstream numbers ("numbers.txt");
    double num;
    vector<double> vec;

    while (numbers >> num) {
        vec.push_back(num);
    }
    
    double size = static_cast<double>(vec.size());
    double total = 0;
    for(int i=0; i<size; i++){
        total += vec[i];
    }
    
    cout << "high: " << high(vec) << endl;
    cout << "low: " << low(vec) << endl;
    cout << "total: " << total << endl;
    cout << "average: " << total / size << endl;
}

double high(vector<double> numbers) {
    double size = static_cast<double>(numbers.size());
    double high = 0;
    for(int i=0; i<size; i++) {
        if (numbers[i] > high) {
            high = numbers[i];
        }
    }

    return high;
}

double low(vector<double> numbers) {
    double size = static_cast<double>(numbers.size());
    double low = numbers[0];
    for(int i=0; i<size; i++) {
        if (numbers[i] < low) {
            low = numbers[i];
        }
    }

    return low;
}
