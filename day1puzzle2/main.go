package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func GetIndexesOfSubString(s string, subString string) (indexes []int) {
	i := strings.Index(s, subString)
	if i == -1 {
		return
	}

	indexes = append(indexes, i)
	result := GetIndexesOfSubString(s[i+len(subString):], subString)
	for _, j := range result {
		indexes = append(indexes, i+len(subString)+j)
	}
	return
}

func GetCalibrationValue(s string) (int, error) {
	validNumbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	strings2numbers := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
	}
	firstNumberIndex := len(s) + 1
	firstNumber := ""
	lastNumberIndex := -1
	lastNumber := ""
	for _, validNumber := range validNumbers {
		indexes := GetIndexesOfSubString(s, validNumber)

		if len(indexes) == 0 {
			continue
		}

		if indexes[0] < firstNumberIndex {
			firstNumberIndex = indexes[0]
			firstNumber = strings2numbers[validNumber]
		}

		if indexes[len(indexes)-1] > lastNumberIndex {
			lastNumberIndex = indexes[len(indexes)-1]
			lastNumber = strings2numbers[validNumber]
		}
	}

	completeNumber, err := strconv.Atoi(firstNumber + lastNumber)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("error combining "))
	}
	fmt.Printf("Stringa %v --> %v\n", s, completeNumber)

	return completeNumber, nil

}
