#include <iostream>
using namespace std;

unsigned int stringlen(string);

int main() {
	string input;
	cout << "enter string: ";
	cin >> input;
	
	printf("string length: %o\n", stringlen(input));
}

unsigned int stringlen(string input) {
	return input.length();
}
