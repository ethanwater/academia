#include <iostream>
#include <stdio.h>
#include <string.h>
#include <stdexcept>
using namespace std;

class Date {
	public:
		int day;
		int month;
		int year;

		void init(int d, int m, int y) {
			//error handling if user bypasses input restrictions
			try {
				if (d>31 || d<1 || m>12 || m<1) {
					throw 0;
				}
				day = d;
				month = m;
				year = y;
			}
			catch(int date_exception) {
				cout << "ERROR: date Provided is invalid\n";
				exit(1);
			}
		};

		string determine_month(int m) {
			string month;
			switch(m){
				case 1:
            month = "January";
            break;
        case 2:
            month = "February";
            break;
        case 3:
            month = "March";
            break;
        case 4:
            month = "April";
            break;
        case 5:
            month = "May";
            break;
        case 6:
            month = "June";
            break;
        case 7:
            month = "July";
            break;
        case 8:
            month = "August";
            break;
        case 9:
            month = "September";
            break;
        case 10:
            month = "October";
            break;
        case 11:
            month = "November";
            break;
        case 12:
            month = "December";
            break;
        default:
            month = "Null";
            break;
			}

			return month;
		};

		void universal() {
			string str_year = to_string(year).substr(2,2);
			printf("%u/%u/%s\n", month, day, str_year.c_str());
		};

		void american() {
			string m = determine_month(month);
			printf("%s %u, %u\n", m.c_str(), day, year);
		};

		void british() {
			string m = determine_month(month);
			printf("%u %s %u\n", day, m.c_str(), year);
		};
};

int main() {
	Date date;
	int day, month, year;
	cout << "===DATE FORMATTER===\n";
	cout << "enter the date you would like to format:\n"; 
	cout << "-day: ";cin >> day;
	while (day < 1 || day > 31) {
		cout << "enter a valid day: ";cin >> day;
	}
	cout << "-month: ";cin >> month;
	while (month < 1 || month > 12){
		cout << "enter a valid month: ";cin >> month;
	}
	cout << "-year: "; cin >> year;
	date.init(day, month, year);

	cout << "\nFORMATS:\n";
	cout << "1. universal: "; date.universal();
	cout << "2. american english: "; date.american();
	cout << "3. british english: "; date.british();
}
