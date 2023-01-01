package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type FuckDetHeleMand struct {
	Variabel string
	Type     string
}

func main() {
	// Open the input file
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inputFile.Close()

	outputPath := "output.h"
	var outputFile *os.File

	// Default to stdout
	if outputPath == "" {
		outputFile = os.Stdout
	} else {
		outputFile, err = os.Create(outputPath)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	fmt.Fprintf(outputFile, "#pragma once%s\n\n", "")

	fmt.Fprintf(outputFile, "#include <string>%s\n", "")

	fmt.Fprintf(outputFile, "using namespace std;%s\n\n", "")

	//read the first line of the input.txt file

	scanner := bufio.NewScanner(inputFile)

	var classname string
	//Handles the first line of to determine class name
	for scanner.Scan() {
		//copy the first line of the input.txt file into classname
		classname = scanner.Text()
		break
	}

	//create an array of FuckDetHeleMand structs
	var fuckDetHeleMand []FuckDetHeleMand
	var fuckDetHeleMand2 []FuckDetHeleMand
	for scanner.Scan() {
		// check if the line contains -
		if strings.Contains(scanner.Text(), "-") {

			//split the line into variabel and type seperator by ; or :
			split := strings.Split(scanner.Text(), ";")
			if len(split) == 1 {
				split = strings.Split(scanner.Text(), ":")
			}
			//remove - and whitespace from split
			split[0] = strings.Replace(split[0], "-", "", -1)
			split[0] = strings.Replace(split[0], " ", "", -1)

			//remove whitespace from split[1]
			split[1] = strings.Replace(split[1], " ", "", -1)

			fuckDetHeleMand = append(fuckDetHeleMand, FuckDetHeleMand{split[0], split[1]})
		}

		if strings.Contains(scanner.Text(), "+") {
			//split the line into variabel and type seperator by ; or :
			split := strings.Split(scanner.Text(), ";")
			if len(split) == 1 {
				split = strings.Split(scanner.Text(), ":")
			}
			//remove - and whitespace from split
			split[0] = strings.Replace(split[0], "+", "", -1)
			split[0] = strings.Replace(split[0], " ", "", -1)

			//remove whitespace from split[1]
			split[1] = strings.Replace(split[1], " ", "", -1)

			fuckDetHeleMand2 = append(fuckDetHeleMand2, FuckDetHeleMand{split[0], split[1]})
		}

	}
	fmt.Fprintf(outputFile, "class %s\n", classname)
	fmt.Fprintf(outputFile, "{%s\n", "")
	fmt.Fprintf(outputFile, "private:%s\n", "")

	//loop over the array and print the variabels and types
	for i := 0; i < len(fuckDetHeleMand); i++ {
		fmt.Fprintf(outputFile, "	%s %s;\n", fuckDetHeleMand[i].Type, fuckDetHeleMand[i].Variabel)
	}
	fmt.Fprintf(outputFile, "%s\n", "")
	fmt.Fprintf(outputFile, "private:%s\n", "")

	// loop over the array and print the variabels and types
	for i := 0; i < len(fuckDetHeleMand2); i++ {
		fmt.Fprintf(outputFile, "	%s %s;\n", fuckDetHeleMand2[i].Type, fuckDetHeleMand2[i].Variabel)
	}
	fmt.Fprintf(outputFile, "};%s\n", "")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(fuckDetHeleMand)
	fmt.Println(fuckDetHeleMand2)

	err = outputFile.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	//Create c++ file

	var outputFile2 *os.File
	outputPath2 := "output.cpp"

	if outputPath2 == "" {
		outputFile2 = os.Stdout
	} else {
		outputFile2, err = os.Create(outputPath2)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	fmt.Fprintf(outputFile2, "#include <string>%s\n\n", "")

	fmt.Fprintf(outputFile2, "#include <%s.h> \n\n", classname)

	fmt.Fprintf(outputFile2, "using namespace std;%s\n\n", "")

	for i := 0; i < len(fuckDetHeleMand2); i++ {
		fmt.Fprintf(outputFile2, "%s %s::%s\n", fuckDetHeleMand2[i].Type, classname, fuckDetHeleMand2[i].Variabel)
		fmt.Fprintf(outputFile2, "{%s\n", "")

		fmt.Fprintf(outputFile2, "//implement return value%s\n", "")

		fmt.Fprintf(outputFile2, "}%s\n\n", "")
	}

	/* char Car::getCategory()
	{
		return category_;
	} */

	err = outputFile2.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

}

/* Car
- idNumber; int
- category: char
- model: string
- doors: int
- fuelType: char
- gearType: char
- pricePrDay: double
- isAvailable: bool
+ getCategory(): char
+ getIsAvailable(): bool
+ setIsAvailable(bool): void
+ getIdNumber(): int
+ setIdNumber(int): void
+ print(): void */

/* class Car
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

public:
    Car() = default;
    Car(int id, char cat, string model, int doors, char fuel, char gear, double price);

    char getCategory();
    bool getIsAvailable();
    void setIsAvailable(bool);
    int getIdNumber();
    void setIdNumber(int);
    void print();
}; */
