#include <iostream>

int main() {
    int patientStatus = 0;
    double days = 0, rate = 0, medication = 0, services = 0;

    std::cout << "enter 0 if IN-PATIENT or 1 if OUT-PATIENT: ";
    std::cin >> patientStatus;

    std::cout << "charges for hospital services: ";
    std::cin >> services;
    std::cout << "charges for medication: ";
    std::cin >> medication;

    if (patientStatus == 0) {
        std::cout << "days in hospital: ";
        std::cin >> days;
        std::cout << "daily rate: ";
        std::cin >> rate;
    }

    std::cout << "total charges: " << (rate * days) + medication + services;    
}
