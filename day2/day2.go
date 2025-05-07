package day2

func Part1(data []string) int {
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

func Part2(data []string) string {
	for _, subject := range data {
		for _, target := range data {
			diffCount := compare(subject, target)
			if diffCount == 1 {
				newString := getCommon(subject, target)
				return newString
			}
		}
	}
	return "" // not expected
}

func getCommon(one, two string) string {
	for i := 0; i < len(one); i++ {
		if one[i] != two[i] {
			newstring := one[:i] + one[i+1:]
			return newstring
		}
	}
	return "" // not expected
}

func compare(subject, target string) int {
	sameCount := 0
	count := len(subject)
	for i := 0; i < count; i++ {
		if subject[i] == target[i] {
			sameCount++
		}
	}
	return count - sameCount
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
		register[rune]++
	}
	return register
}
