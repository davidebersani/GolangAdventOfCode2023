func ConvertToint(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Error converting number ", n)
	}
	return n
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