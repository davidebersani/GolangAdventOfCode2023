package main

import (
	"bufio"
	"fmt"
	"log"
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
	count := 0
	allNumbers := strings.Split(card, ": ")[1]
	winningNumbersStrings := strings.Split(strings.Split(allNumbers, " | ")[0], " ")
	winningNumbers := convertToint(removeEmptyItems(winningNumbersStrings))
	userNumbers := removeEmptyItems(strings.Split(strings.Split(allNumbers, " | ")[1], " "))
	// fmt.Println(winningNumbers, userNumbers)
	for _, n := range userNumbers {
		nint, err := strconv.Atoi(strings.Trim(n, " "))
		if err != nil {
			log.Fatal("error converting number ", n)
		}
		if isItemInList[int](nint, winningNumbers) {
			count++
		}
	}

	return count
}

func main() {
	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Printf("Error opening the file: %v", err)
		return
	}

	count := 0
	card2copies := map[int]int{}
	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Split(strings.Split(line, ": ")[0], " ")
		cardNumberStr := parts[len(parts)-1]
		cardNumber, err := strconv.Atoi(cardNumberStr)
		if err != nil {
			log.Fatal("Error converting number ", cardNumber)
		}
		copies := card2copies[cardNumber]
		copies++ // +1 it's the card itself
		fmt.Printf("Scratching %v copies of card %v\n", copies, cardNumber)
		for c := 0; c < copies; c++ {
			res := CheckCard(line)
			// fmt.Printf("You won %v more scratchcards\n", res)
			count++
			for i := 1; i <= res; i++ {
				card2copies[cardNumber+i]++
			}
		}
	}
	fmt.Println("Sum of is", count)
}
