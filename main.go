package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// Struct to hold the scanned variabels and types
type variable struct {
	name string
	typ  string
}

// function that prepares the output files.
func outputFile(name string) (*os.File, error) {
	var outputFile *os.File
	var err error

	// Default to stdout
	if name == "" {
		outputFile = os.Stdout
	} else {
		outputFile, err = os.Create(name)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	return outputFile, nil
}

// Function that opens the input file or creates a new one if it doesn't exist
func inputFile() (*os.File, error) {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File not found, creating new file")
		inputFile, err = os.Create("input.txt")
		if err != nil {
			return nil, err
		}
	}
	return inputFile, nil
}

func main() {
	//Prepare the in and output files for reading and writing
	input, err := inputFile()
	if err != nil {
		log.Fatal(err.Error())
	}

	outputHeader, err := outputFile("output.h")
	if err != nil {
		log.Fatal(err.Error())
	}

	outputCPP, err := outputFile("output.c++")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Fprint(outputHeader, "#pragma once\n\n")

	//Create variables to hold the class name and the variabels and types
	var classname string
	var parsedVar []variable
	var parsedMethod []variable

	//Create the scanner to read the input file
	scanner := bufio.NewScanner(input)
	//Handles the first line of to determine class name
	for scanner.Scan() {
		//copy the first line of the input.txt file into classname
		classname = scanner.Text()
		break
	}

	//Scan the input file for variabels and types
	for scanner.Scan() {
		// check if the line contains -
		if strings.Contains(scanner.Text(), "-") {
			//split the line into variabel and type seperator by ; or :
			split := strings.Split(scanner.Text(), ";")
			if len(split) == 1 {
				split = strings.Split(scanner.Text(), ":")
			}
			//Magic to remove - and whitespace from split
			r := strings.NewReplacer("-", "", " ", "")
			split[0] = r.Replace(split[0])
			split[1] = r.Replace(split[1])

			parsedVar = append(parsedVar, variable{split[0], split[1]})
		}

		if strings.Contains(scanner.Text(), "+") {
			//split the line into variabel and type seperator by ; or :
			split := strings.Split(scanner.Text(), ";")
			if len(split) == 1 {
				regex := regexp.MustCompile(`\B:`)
				split = regex.Split(scanner.Text(), -1)
			}

			//Magic to remove - and whitespace from split
			r := strings.NewReplacer("+", "", " ", "")
			split[0] = r.Replace(split[0])
			split[1] = r.Replace(split[1])

			fmt.Println(split[0])
			subString := strings.SplitAfter(split[0], "(")

			fmt.Println(subString[1])

			if strings.Contains(subString[1], ",") {
				subString = strings.Split(subString[1], ",")
				fmt.Println("==", subString)

				for i := 0; i < len(subString); i++ {

					if strings.Contains(subString[i], ":") {
						subString[i] = strings.Replace(subString[i], ":", " ", -1)

						for j := 0; j < len(subString)-1; j++ {

							//split the string by whitespace and swap index 0 with index 1
							t := strings.Split(subString[j], " ")

							subString[j+1] = t[j]
							subString[j] = t[j+1]

						}

					}

				}
			} else if strings.Contains(subString[1], ":") {
				subString = strings.Split(subString[1], ":")

				//loop through the substring and exchange index 0 with index 1 and vice versa

				for i := 0; i < len(subString)-1; i++ {
					t := subString[i]
					subString[i] = subString[i+1]
					subString[i+1] = t
				}

			}

			fmt.Println(subString)
			fmt.Println(parsedMethod)
			fmt.Println(split[0])

			//go into split index 0 and remove everything inside the brackets
			regex := regexp.MustCompile(`\(.+\)`)

			split[0] = regex.ReplaceAllString(split[0], "")
			split[0] = fmt.Sprintf("%s(%s)", split[0], strings.Join(subString, ","))

			fmt.Println(split[0])

			parsedMethod = append(parsedMethod, variable{split[0], split[1]})

		}
	}
	//Error checking on scanner
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//Function call where all the magic writing to file happens for the header file
	writeToHeader(outputHeader, classname, parsedVar, parsedMethod)
	err = outputHeader.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	//function call where all the magic writing to file happens for the cpp file
	writeToCpp(outputCPP, classname, parsedVar, parsedMethod)
	err = outputCPP.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = input.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func writeToHeader(file *os.File, classname string, parsedVar []variable, parsedMethod []variable) {
	fmt.Fprintf(file, "#pragma once%s\n\n", "")
	fmt.Fprintf(file, "#include <string>%s\n\n", "")
	fmt.Fprintf(file, "using namespace std;%s\n\n", "")
	fmt.Fprintf(file, "class %s\n", classname)
	fmt.Fprintf(file, "{%s\n", "")
	fmt.Fprintf(file, "private:%s\n", "")

	//loop over the array and print the variabels and types
	for i := 0; i < len(parsedVar); i++ {
		fmt.Fprintf(file, "	%s %s;\n", parsedVar[i].typ, parsedVar[i].name)
	}
	fmt.Fprintf(file, "%s\n", "")
	fmt.Fprintf(file, "public:%s\n", "")

	// loop over the array and print the variabels and types
	for i := 0; i < len(parsedMethod); i++ {
		fmt.Fprintf(file, "	%s %s;\n", parsedMethod[i].typ, parsedMethod[i].name)
	}
	fmt.Fprintf(file, "};%s\n", "")
}

func writeToCpp(file *os.File, classname string, parsedVar []variable, parsedMethod []variable) {
	fmt.Fprintf(file, "#include <%s.h> \n\n", classname)

	fmt.Fprintf(file, "using namespace std;%s\n\n", "")

	for i := 0; i < len(parsedMethod); i++ {
		fmt.Fprintf(file, "%s %s::%s\n", parsedMethod[i].typ, classname, parsedMethod[i].name)
		fmt.Fprintf(file, "{%s\n", "")

		fmt.Fprintf(file, "//implement return value%s\n", "")

		fmt.Fprintf(file, "}%s\n\n", "")
	}
}

/* fmt.Println("pre split == ", split[0])

limiter := regexp.MustCompile(`\(([^()]*)\)`)

result := limiter.ReplaceAllString(split[0], "$1")

//split the result into variabel and type seperator by :
subString := strings.Split(result, ":")
//swap the value to the right and left of the : inside the () to match the format of the header file

fmt.Println("==", subString[0])

regex := regexp.MustCompile(`\b:`)
if regex.MatchString(split[0]) {
	fmt.Println("??", subString[0])
	split[0] = subString[1] + " " + "(" + subString[0] + ")"

	/* 	for i := 0; i < len(split); i++ {
		fmt.Println("??", split[i])
	} */

//fmt.Println(split[0]) */
