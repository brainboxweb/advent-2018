package day2

func Part1(data []string) int {
	// things := strings.Split(input, "\n")
	twiceCount, thriceCount := 0, 0
	for _, item := range data {
		tokens := tokenise(item)
		if haveThisCount(tokens, 2) {
			twiceCount++
		}
		if haveThisCount(tokens, 3) {
			thriceCount++
		}
	}
	return twiceCount * thriceCount
}

func haveThisCount(theSet map[rune]int, count int) bool {
	for _, val := range theSet {
		if val == count {
			return true
		}
	}
	return false
}

func tokenise(theString string) map[rune]int {
	register := make(map[rune]int)
	for _, rune := range theString {
		_, ok := register[rune] // Already stored?
		if !ok {
			register[rune] = 1
		} else {
			register[rune]++
		}
	}
	return register
}
