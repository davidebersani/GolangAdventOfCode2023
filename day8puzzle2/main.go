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

func GetNumStep(node string, node2AdjacentNodes map[string]Node, instructions string, finalNodech chan int) {
	i := 0
	for node[len(node)-1] != 'Z' {
		for _, instr := range strings.Split(instructions, "") {
			if instr == "L" {
				node = node2AdjacentNodes[node].left
			} else {
				node = node2AdjacentNodes[node].right
			}
		}
		i++
	}
	finalNodech <- i
}

func calculateLCM(numbers []int) int {
	lcm := numbers[0]
	for i := 1; i < len(numbers); i++ {
		lcm = lcm * numbers[i] / calculateGCD(lcm, numbers[i])
	}
	return lcm
}

func calculateGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
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
	startNodes := []string{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Split(line, " = ")
		node2AdjacentNodes[parts[0]] = Node{
			left:  parts[1][1:4],
			right: parts[1][6:9],
		}

		if parts[0][len(parts[0])-1] == 'A' {
			startNodes = append(startNodes, parts[0])
		}
	}

	fmt.Println("Start nodes:", startNodes)
	ch := make(chan int)
	for _, sn := range startNodes {
		go GetNumStep(sn, node2AdjacentNodes, instructions, ch)
	}

	numSteps := []int{}
	for i := 0; i < len(startNodes); i++ {
		numSteps = append(numSteps, <-ch)
	}

	fmt.Println("Num steps:", numSteps)

	lcm := calculateLCM(numSteps)
	fmt.Println("LCM of numSteps:", lcm*len(instructions))

}
