#include <iostream>
#include <string>

struct MOVIE_DATA{
	std::string title;
	std::string director;
	unsigned short int year; 
	unsigned short int time;
	void display() {
		printf("The film %s directed by %s was released in %u and has a running time of %u minutes.\n", 
				title.c_str(), 
				director.c_str(), 
				year, 
				time);
	}
	void init(std::string t, std::string d, unsigned short int y, unsigned short int l) {
				title = t;
				director = d;
				year = y;
				time = l;
	}
};

int main() {
	struct MOVIE_DATA movie1;
	struct MOVIE_DATA movie2;
	movie1.init("Titanic", "James Cameron", 1997, 195);
	movie2.init("Interstellar","James Cameron", 2014, 169);	

	movie1.display();
	movie2.display();
}
