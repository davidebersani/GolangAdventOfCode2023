func main() {
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
		gameStr := strings.Split(line, ": ")[1]
		sum += CalculatePowerOfMinSet(gameStr)
	}
	fmt.Println("Sum of is", sum)
}
