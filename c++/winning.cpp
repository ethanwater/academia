#include <iostream>
#include <string>
#include <vector>

double getSales(std::string divisionName) {
    double divSales;
    std::cout << "enter quarterly sales figure: ";
    std::cin >> divSales;
    if (divSales < 0) {
        std::cin >> divSales;
    }

    return divSales;
}

void findHighest(std::vector<double> sales) {
    double highestSales = sales[0];
    int index = 0;
    for(int i=0; i<4; i++) {
        if (sales[i] > highestSales) {
            highestSales = sales[i];
            index = i;
        }
    }

    switch(index) {
        case 0:
            std::cout << "highest grossing division: NE" << std::endl;
            break;
        case 1:
            std::cout << "highest grossing division: SE" << std::endl;
            break;
        case 2:
            std::cout << "highest grossing division: NW" << std::endl;
            break;
        case 3:
            std::cout << "highest grossing division: SW" << std::endl;
            break;
        default:
            std::cout << "null";
            break;
    }

    std::cout << "highest sales: " << highestSales;
}

int main() {
    std::string divisionName;
    std::vector<double> sales;

    for (int i=0; i<4; i++) {
        std::cout << "enter division name:";
        std::cin >> divisionName;

        sales.push_back(getSales(divisionName));
    }
    
    findHighest(sales);
}
