#include <iostream>
#include <string>
#include <algorithm>
using namespace std;

double sum(double[]);
double high(double[]);
double low(double[]);
string determineMonth(double, double[]);

const int MONTHS = 12;
int main() {
    double rainfall[MONTHS];
    
    cout << "enter rainfall per month:" << endl;
    for (int i=0; i<MONTHS; i++) {
        cin >> rainfall[i];
        if (rainfall[i] < 0) {
            cout << "please reenter, negative values are invalid: " << endl;
            cin >> rainfall[i];
        }
    }

    double total_rain = sum(rainfall);
    cout << "annual rainfall: " << total_rain << endl;
    cout << "average rainfall per month: " << total_rain / MONTHS << endl;
    cout << "highest rainfall month: " << determineMonth(high(rainfall), rainfall) << endl;
    cout << "lowest rainfall month: " << determineMonth(low(rainfall), rainfall) << endl;
}

double sum(double rainfall[]) {
    double rainSum = 0;
    for (int i=0; i<MONTHS; i++) {
        rainSum += rainfall[i];
    }
    return rainSum;
}

double high(double rainfall[]) {
    double high = 0;
    for(int i=0; i<MONTHS; i++) {
        if (rainfall[i] > high) {
            high = rainfall[i];
        }
    }

    return high;
}

double low(double rainfall[]) {
    double low = rainfall[0];
    for(int i=0; i<MONTHS; i++) {
        if (rainfall[i] < low) {
            low = rainfall[i];
        }
    }
    
    return low;
}

string determineMonth(double amount, double rainfall[]) {
    int month_amt = distance(rainfall, find(rainfall, rainfall + MONTHS, amount));
    string month;

    switch(month_amt)
    {
        case 0:
            month = "january";
            break;
        case 1:
            month = "february";
            break;
        case 2:
            month = "march";
            break;
        case 3:
            month = "april";
            break;
        case 4:
            month = "may";
            break;
        case 5:
            month = "june";
            break;
        case 6:
            month = "july";
            break;
        case 7:
            month = "august";
            break;
        case 8:
            month = "september";
            break;
        case 9:
            month = "october";
            break;
        case 10:
            month = "november";
            break;
        case 11:
            month = "december";
            break;
        default:
            month = "null";
            break;
    }
    
    return month;
}


