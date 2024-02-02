#include <string>
#include <iostream>
using namespace std;

int main() {
    string contactList[11] = {"Becky Warren, 555-1223",
                              "Joe Looney, 555-0097",
                              "Geri Palmer, 555-8787",
                              "Lynn Presnell, 555-1212",
                              "Holly Gaddis, 555-8878",
                              "Sam Wiggins, 555-0998",
                              "Bob Kain, 555-8712",
                              "Tim Haynes, 555-7676",
                              "Warren Gaddis, 555-9037",
                              "Jean James, 555-4939",
                              "Ron Palmer, 555-2783"};

    string request;
    cout << "search for: "; cin >> request;

    for(int i=0; i<11; i++) {
        if( contactList[i].find(request) != string::npos) {
            cout << contactList[i];
        } 
    }
}
