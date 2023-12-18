package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	line := ""
	sum := 0
	for fileScanner.Scan() {
		line = fileScanner.Text()
		calValue, err := GetCalibrationValue(line)
		if err != nil {
			log.Fatal("Error: ", err)
		}

		sum += calValue
	}

	fmt.Println("Sum is", sum)
}

func GetCalibrationValue(s string) (int, error) {
	firstNumberString := ""
	lastNumberString := ""
	for _, char := range s {
		_, err := strconv.Atoi(string(char))
		if err != nil {
			continue
		}

		if firstNumberString == "" {
			firstNumberString = string(char)
		}
		lastNumberString = string(char)
	}

	completeNumber, err := strconv.Atoi(firstNumberString + lastNumberString)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("impossible to convert the complete number combining %v and %v", firstNumberString, lastNumberString))
	}
	return completeNumber, nil
}
