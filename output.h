#pragma once

#include <string>
using namespace std;

class Car
{
private:
	int idNumber_;
	char category_;
	string model_;
	int doors_;
	char fuelType_;
	char gearType_;
	double pricePrDay_;
	bool isAvailable_;

private:
	char getCategory();
	bool getIsAvailable();
	void setIsAvailable(bool);
	int getIdNumber();
	void setIdNumber(int);
	void print();
};
