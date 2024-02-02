#include <iostream>
#include <stdio.h>
#include <stdlib.h>
#include <map>

using namespace std;

struct MONTHLY_BUDGET{
	double housing;
	double utilities;
	double household_expenses;
	double transportation;
	double food;
	double medical;
	double insurance;
	double entertainment;
	double clothing;
	double misc;

	void init(
		double h, double u,
		double he, double t,
		double f, double me,
		double i, double e,
		double c, double mi
	){
		housing = h;
		utilities = u;
	  household_expenses = he;
		transportation = t;
		food = f;
		medical = me;
		insurance = i;
		entertainment = e;
		clothing = c;
		misc = mi;
	};

	double sum(){
		double sum = housing + utilities + household_expenses+ transportation + food + medical 
								 + insurance + entertainment + clothing + misc;
		return sum;
	}

	
};

struct MONTHLY_BUDGET amt_spent();
void report(MONTHLY_BUDGET budget, MONTHLY_BUDGET spent);

int main(){
	MONTHLY_BUDGET budget;
	budget.init(500.00,150.00,65.00,50.00,250.00,30.00,100.00,150.00,75.00,50.00);
	MONTHLY_BUDGET spent = amt_spent();

	report(budget, spent);
}

struct MONTHLY_BUDGET amt_spent(){
	MONTHLY_BUDGET spent;
	double h, u, he, t, f, me, i, e, c, mi;

	cout << "===BUDGET CALCULATOR===\nEnter the amount spent this month per category:\n"; 
	cout << "housing: ";cin >> h;
	cout << "utilities: ";cin >> u;
	cout << "household_expenses: ";cin >> he;
	cout << "transportation: ";cin >> t;
	cout << "food: ";cin >> f;
	cout << "medical: ";cin >> me;
	cout << "insurance: ";cin >> i;
	cout << "entertainment: ";cin >> e;
	cout << "clothing: ";cin >> c;
	cout << "misc: ";cin >> mi;
	
	spent.init(h, u, he, t, f, me, i, e, c, mi);
	return spent;
}

void report(MONTHLY_BUDGET budget, MONTHLY_BUDGET spent){
	cout << "\nMONTHLY BUDGET REPORT: \n";
	map<string, double> categories = {
		{"housing", budget.housing - spent.housing},
		{"utilities", budget.utilities - spent.utilities},
		{"household_expenses", budget.household_expenses - spent.household_expenses},
		{"transportation", budget.transportation - spent.transportation},
		{"food", budget.food - spent.food},
		{"medical", budget.medical - spent.medical},
		{"insurance", budget.insurance - spent.insurance},
		{"entertainment", budget.entertainment - spent.entertainment},
		{"clothing", budget.clothing - spent.clothing},
		{"misc", budget.misc - spent.misc},
	};

	for(auto const& cat : categories) {
		string field = cat.first;
		double amt = cat.second;
		if (amt < 0) {
			printf("%s is OVER the budget by $%.2f\n", field.c_str(), abs(amt));
		} else if (amt == 0) {
			printf("%s is exactly on budget\n", field.c_str());
		} else {
			printf("%s is UNDER the budget by $%.2f\n", field.c_str(), amt);
		}
	}

	double sum_budget = budget.sum(), sum_spent = spent.sum();
	if (sum_budget > sum_spent) {
		printf("\nstudent has spent %.2f under the monthly budget.\n", sum_budget-sum_spent);
	} else if (sum_budget == sum_spent) {
		printf("\nstudent has spent the exact monthly budget.\n");
	} else {
		printf("\nstudent has spent %.2f over the monthly budget.\n", sum_spent-sum_budget);
	}
}
