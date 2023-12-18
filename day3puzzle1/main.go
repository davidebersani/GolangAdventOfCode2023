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
	number int
	start  Position
	end    Position
}

func (p Position) equals(p2 Position) bool {
	return p.x == p2.x && p.y == p2.y
}

func GetSymbolPosition(rowNumber int, s string) []Position {
	positions := []Position{}
	for i, ch := range s {
		// Exclude points
		if ch == '.' {
			continue
		}
		// Exclude numbers
		_, err := strconv.Atoi(string(ch))
		if err == nil {
			continue
		}
		// It's a symbol
		positions = append(positions, Position{x: rowNumber, y: i})
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
			result = append(result, Number{number: n, start: Position{x: x, y: start}, end: Position{x: x, y: end}})
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

func GetSumOfValidNumbers(row int, s string, symbolPositions []Position) int {
	numbers := GetNumbers(row, s)
	fmt.Printf("Numers found in row %v: %v\n", row, numbers)
	sum := 0
	for _, n := range numbers {
		adjacentPositions := GetAdjacentPositions(row, n.start.y, n.end.y)
		found := false
		for _, p := range adjacentPositions {
			for _, validPosition := range symbolPositions {
				if p.equals(validPosition) {
					fmt.Printf("Number %v is valid because is adjacent at symbol in %v\n", n, validPosition)
					sum += n.number
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}
	return sum
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Printf("Error opening the file: %v", err)
		return
	}

	SymbolPositions := []Position{}

	fileScanner := bufio.NewScanner(readFile)
	i := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		SymbolPositions = append(SymbolPositions, GetSymbolPosition(i, line)...)
		i++
	}
	readFile.Close()

	readFile, err = os.Open("input.txt")

	if err != nil {
		fmt.Printf("Error opening the file: %v", err)
		return
	}
	defer readFile.Close()

	fmt.Println("Symbols in", SymbolPositions)
	fileScanner = bufio.NewScanner(readFile)
	i = 0
	sum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		sum += GetSumOfValidNumbers(i, line, SymbolPositions)
		i++
	}
	fmt.Println("The sum is", sum)
}
