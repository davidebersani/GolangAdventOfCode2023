package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func isItemInList[T comparable](item T, list []T) bool {
	for _, o := range list {
		if o == item {
			return true
		}
	}
	return false
}

func convertToint(list []string) (result []int) {
	for _, n := range list {
		nint, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal("error converting number ", n)
		}
		result = append(result, nint)
	}
	return
}

func removeEmptyItems(list []string) (result []string) {
	for _, item := range list {
		if strings.Trim(item, " ") == "" {
			continue
		}
		result = append(result, item)
	}
	return
}

func CheckCard(card string) int {
	points := -1
	allNumbers := strings.Split(card, ": ")[1]
	winningNumbersStrings := strings.Split(strings.Split(allNumbers, " | ")[0], " ")
	winningNumbers := convertToint(removeEmptyItems(winningNumbersStrings))
	userNumbers := removeEmptyItems(strings.Split(strings.Split(allNumbers, " | ")[1], " "))
	fmt.Println(winningNumbers, userNumbers)
	for _, n := range userNumbers {
		nint, err := strconv.Atoi(strings.Trim(n, " "))
		if err != nil {
			log.Fatal("error converting number ", n)
		}
		if isItemInList[int](nint, winningNumbers) {
			points++
		}
	}
	if points == -1 {
		return 0
	}

	return int(math.Pow(2, float64(points)))
}

func main() {
	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Printf("Error opening the file: %v", err)
		return
	}

	sum := 0
	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		sum += CheckCard(line)
		fmt.Println(sum)
	}
	fmt.Println("Sum of is", sum)
}
