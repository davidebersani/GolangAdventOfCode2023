package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetMinSet(gameStr string) map[string]int {
	minSet := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, extraction := range strings.Split(gameStr, "; ") {
		for _, count := range strings.Split(extraction, ", ") {
			pp := strings.Split(count, " ")
			n, err := strconv.Atoi(pp[0])
			if err != nil {
				log.Fatal("Error converting number", pp[0])
			}

			if minSet[pp[1]] < n {
				minSet[pp[1]] = n
			}
		}
	}
	return minSet
}

func CalculatePowerOfMinSet(gameStr string) (power int) {
	minSet := GetMinSet(gameStr)
	return minSet["red"] * minSet["blue"] * minSet["green"]
}

func main() {
	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Printf("Error opening the file: %v", err)
		return
	}

	powerSum := 0
	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		gameStr := strings.Split(line, ": ")[1]
		powerSum += CalculatePowerOfMinSet(gameStr)
	}
	fmt.Println("Sum of the power is", powerSum)
}
