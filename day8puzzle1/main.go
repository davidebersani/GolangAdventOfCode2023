package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	left  string
	right string
}

func main() {
	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Printf("Error opening the file: %v", err)
		return
	}

	fileScanner := bufio.NewScanner(readFile)
	// First row: instructions
	fileScanner.Scan()
	instructions := fileScanner.Text()

	// Skip line
	fileScanner.Scan()

	// Init map of nodes
	// node "AAA" -> [ "BBB", "CCC" ] BBB is the left adjacent, CCC is the right one
	node2AdjacentNodes := map[string]Node{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Split(line, " = ")
		node2AdjacentNodes[parts[0]] = Node{
			left:  parts[1][1:4],
			right: parts[1][6:9],
		}
	}

	i := 0
	node := "AAA"
	for node != "ZZZ" {
		for _, instr := range strings.Split(instructions, "") {
			if instr == "L" {
				node = node2AdjacentNodes[node].left
			} else {
				node = node2AdjacentNodes[node].right
			}
		}
		i++
	}

	fmt.Println("Number of step:", i*len(instructions))
}
