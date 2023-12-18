package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Position struct {
	x int
	y int
}

type Number struct {
	number            int
	start             Position
	end               Position
	adjacentPositions []Position
}

func (p Position) equals(p2 Position) bool {
	return p.x == p2.x && p.y == p2.y
}

func getPotentialsGears(rowNumber int, s string) []Position {
	positions := []Position{}
	for i, ch := range s {
		// Exclude points
		if ch == '*' {
			positions = append(positions, Position{x: rowNumber, y: i})
		}
	}
	return positions
}

func GetNumbers(x int, s string) []Number {
	result := []Number{}
	i := 0
	for i < len(s) {
		n, err := strconv.Atoi(string(s[i]))
		if err == nil {
			start := i
			end := i
			i++
			nString := string(s[start])
			for i < len(s) {
				_, err := strconv.Atoi(string(s[i]))
				if err == nil {
					nString += string(s[i])
					end++
					i++
				} else {
					break
				}
			}
			n, err = strconv.Atoi(nString)
			if err != nil {
				log.Fatal("Errore nel convertire il numero ", nString)
			}
			result = append(result, Number{
				number:            n,
				start:             Position{x: x, y: start},
				end:               Position{x: x, y: end},
				adjacentPositions: GetAdjacentPositions(x, start, end),
			})
		}
		i++
	}
	return result
}

func GetAdjacentPositions(x int, start int, end int) []Position {
	result := []Position{}
	for y := start - 1; y <= end+1; y++ {
		result = append(result, Position{x: x - 1, y: y})
		result = append(result, Position{x: x + 1, y: y})
		if y < start || y > end {
			result = append(result, Position{x: x, y: y})
		}
	}
	return result
}

func GetSumOfGearsPower(row int, s string, numbers []Number) int {
	potentialsGears := getPotentialsGears(row, s)
	fmt.Printf("Potentials gears found in row %v: %v\n", row, potentialsGears)
	sum := 0
	for _, g := range potentialsGears {
		adjacentNumbers := []Number{}
		for _, number := range numbers {
			for _, p := range number.adjacentPositions {
				if p.equals(g) {
					adjacentNumbers = append(adjacentNumbers, number)
					break
				}
			}
		}
		fmt.Println(adjacentNumbers)
		if len(adjacentNumbers) == 2 {
			fmt.Printf("Gear %v is valid because is adjacent at numbers %v in %v", g, adjacentNumbers)
			sum += adjacentNumbers[0].number * adjacentNumbers[1].number
		}
	}
	return sum
}

// Per ogni riga, get numbers e get adjacent positions
// Poi, per ogni riga, trova l'ingranaggio (*) e controlla se è adiacente a due numeri

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Printf("Error opening the file: %v", err)
		return
	}

	numbers := []Number{}

	fileScanner := bufio.NewScanner(readFile)
	i := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		numbers = append(numbers, GetNumbers(i, line)...)
		i++
	}
	readFile.Close()

	readFile, err = os.Open("input.txt")

	if err != nil {
		fmt.Printf("Error opening the file: %v", err)
		return
	}
	defer readFile.Close()

	fileScanner = bufio.NewScanner(readFile)
	i = 0
	sum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		sum += GetSumOfGearsPower(i, line, numbers)
		i++
	}
	fmt.Println("The sum is", sum)
}
