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

type Interval struct {
	sourceStart int
	destStart   int
	rangeLen    int
}

func ConvertToint(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Error converting number ", n)
	}
	return n
}

func FindMinLocationForRangeOfSeeds(start int, end int, maps map[string]map[string][]Interval, ch chan int) {
	minLocation := math.MaxInt64
	for s := start; s < end; s++ {
		soil := s
		for _, interval := range maps["seed"]["soil"] {
			if s >= interval.sourceStart && s < interval.sourceStart+interval.rangeLen {
				soil = interval.destStart + (s - interval.sourceStart)
				break
			}
		}
		fertilizer := soil
		for _, interval := range maps["soil"]["fertilizer"] {
			if soil >= interval.sourceStart && soil < interval.sourceStart+interval.rangeLen {
				fertilizer = interval.destStart + (soil - interval.sourceStart)
				break
			}
		}
		water := fertilizer
		for _, interval := range maps["fertilizer"]["water"] {
			if fertilizer >= interval.sourceStart && fertilizer < interval.sourceStart+interval.rangeLen {
				water = interval.destStart + (fertilizer - interval.sourceStart)
				break
			}
		}
		light := water
		for _, interval := range maps["water"]["light"] {
			if water >= interval.sourceStart && water < interval.sourceStart+interval.rangeLen {
				light = interval.destStart + (water - interval.sourceStart)
				break
			}
		}
		temperature := light
		for _, interval := range maps["light"]["temperature"] {
			if light >= interval.sourceStart && light < interval.sourceStart+interval.rangeLen {
				temperature = interval.destStart + (light - interval.sourceStart)
				break
			}
		}
		humidity := temperature
		for _, interval := range maps["temperature"]["humidity"] {
			if temperature >= interval.sourceStart && temperature < interval.sourceStart+interval.rangeLen {
				humidity = interval.destStart + (temperature - interval.sourceStart)
				break
			}
		}
		location := humidity
		for _, interval := range maps["humidity"]["location"] {
			if humidity >= interval.sourceStart && humidity < interval.sourceStart+interval.rangeLen {
				location = interval.destStart + (humidity - interval.sourceStart)
				break
			}
		}

		if location < minLocation {
			minLocation = location
		}
	}

	ch <- minLocation
}

func main() {
	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Printf("Error opening the file: %v", err)
		return
	}

	// Init all maps
	maps := map[string]map[string][]Interval{}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Scan()
	line := fileScanner.Text()
	seeds := []int{}
	for _, s := range strings.Split(strings.Split(line, ": ")[1], " ") {
		seeds = append(seeds, ConvertToint(s))
	}
	source := ""
	dest := ""
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, "map:") {
			mapName := strings.Split(line, " ")[0]
			source = strings.Split(mapName, "-to-")[0]
			dest = strings.Split(mapName, "-to-")[1]
			if _, ok := maps[source]; !ok {
				maps[source] = map[string][]Interval{}
			}
			if _, ok := maps[source][dest]; !ok {
				maps[source][dest] = []Interval{}
			}
		} else {
			parts := strings.Split(line, " ")
			destRangeStart := ConvertToint(parts[0])
			sourceRangeStart := ConvertToint(parts[1])
			rangeLength := ConvertToint(parts[2])
			maps[source][dest] = append(maps[source][dest], Interval{sourceStart: sourceRangeStart, destStart: destRangeStart, rangeLen: rangeLength})
		}
	}

	minLocation := math.MaxInt64
	ch := make(chan int)
	for i := 0; i < len(seeds); i = i + 2 {
		go FindMinLocationForRangeOfSeeds(seeds[i], seeds[i]+seeds[i+1], maps, ch)
	}

	for i := 0; i < len(seeds); i = i + 2 {
		loc := <-ch
		fmt.Printf("New location found: %v\n", loc)
		if loc < minLocation {
			minLocation = loc
		}
	}

	fmt.Println("Lowest location is:", minLocation)
}
