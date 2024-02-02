#include <iostream>


int main() {
    int floors;
    double totalRooms;
    double totalOccupied;
    
    std::cout << "how many floors in the hotel?";
    std::cin >> floors;
    
    for (int i=0; i<floors; i++) {
        if (i == 13) {
            continue;
        }
        int rooms;
        int occupied;
    
        std::cout << "how many rooms on the floor?";
        std::cin >> rooms;
        std::cout << "how many occupied?";
        std::cin >> occupied;
    
        totalRooms += rooms;
        totalOccupied += rooms;
    }
    
    std::cout << "rooms: " << totalRooms;
    std::cout << "occupied: " << totalOccupied;
    std::cout << "available: " << totalRooms - totalOccupied;
    std::cout << "percentage of occupied rooms: " << totalOccupied / totalRooms;
}

