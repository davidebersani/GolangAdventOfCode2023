package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	N_RED := 12
	N_GREEN := 13
	N_BLUE := 14

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
		parts := strings.Split(line, ": ")
		// Check cubes
		if CheckCubes(parts[1], N_RED, N_GREEN, N_BLUE) {
			gameNumber, err := strconv.Atoi(strings.Split(parts[0], " ")[1])
			if err != nil {
				log.Fatal("Error converting string", parts[0], "to number")
			}
			sum += gameNumber
		}
	}
	fmt.Println("Sum is", sum)
}

func CheckCubes(s string, N_RED int, N_GREEEN int, N_BLUE int) bool {
	extractions := strings.Split(s, "; ")
	for _, p := range extractions {
		counts := strings.Split(p, ", ")
		for _, c := range counts {
			color := strings.Split(c, " ")[1]
			n, err := strconv.Atoi(strings.Split(c, " ")[0])
			if err != nil {
				log.Fatal("Error extracting count from", c)
			}
			if (color == "red" && n > N_RED) || (color == "green" && n > N_GREEEN) || (color == "blue" && n > N_BLUE) {
				return false
			}
		}
	}
	return true
}
