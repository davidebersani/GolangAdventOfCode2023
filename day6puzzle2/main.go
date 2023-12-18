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

func ConvertToint(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Error converting number ", n)
	}
	return n
}

func GetNumOfWayToWin(t int, d int) (n int) {
	t1 := float64(t)
	d1 := float64(d)
	println("T1 = ", t1)
	println("D1 = ", d1)
	x1 := (t1 - math.Sqrt((t1*t1)+(-4*d1))) / 2
	x2 := (t1 + math.Sqrt((t1*t1)-4*d1)) / 2
	fmt.Printf("Ottengo x1=%v e x2=%v\n", x1, x2)
	fmt.Printf("Ottengo x1=%v e x2=%v\n", int(x1), int(x2))
	n = int(x2) - int(x1)
	if float64(int(x1)) == x1 || float64(int(x2)) == x2 {
		n -= 1
	}
	return
}

func main() {
	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Printf("Error opening the file: %v", err)
		return
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Scan()
	line := fileScanner.Text()
	line = strings.Split(strings.ReplaceAll(line, " ", ""), ":")[1]
	t := ConvertToint(line)

	fileScanner.Scan()
	line = fileScanner.Text()
	line = strings.Split(strings.ReplaceAll(line, " ", ""), ":")[1]
	d := ConvertToint(line)

	result := GetNumOfWayToWin(t, d)
	fmt.Println("Result:", result)

}
