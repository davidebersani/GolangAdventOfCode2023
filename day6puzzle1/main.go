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

func removeEmptyItems(list []string) (result []string) {
	for _, item := range list {
		if strings.Trim(item, " ") == "" {
			continue
		}
		result = append(result, item)
	}
	return
}

func ConvertToInt(s []string) []int {
	result := []int{}
	for _, item := range s {
		n, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal("Error converting number ", n)
		}
		result = append(result, n)
	}
	return result
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
	line = strings.Split(line, ":")[1]
	times := ConvertToInt(removeEmptyItems(strings.Split(line, " ")))

	fileScanner.Scan()
	line = fileScanner.Text()
	line = strings.Split(line, ":")[1]
	distances := ConvertToInt(removeEmptyItems(strings.Split(line, " ")))

	result := 1
	for i := 0; i < len(distances); i++ {
		result = result * GetNumOfWayToWin(times[i], distances[i])
		fmt.Printf("Per t=%v e d=%v ottengo n=%v\n", times[i], distances[i], result)
	}
	fmt.Println("Result:", result)

}
