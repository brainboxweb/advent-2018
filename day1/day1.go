package day1

func Part1(data []int) int {
	res := 0
	for _, num := range data {
		res += num
	}
	return res
}

func Part2(data []int) int {
	frequencies := make(map[int]int) // frequency, count
	freq := 0
	for { // may need many loops
		for _, num := range data {
			freq += num
			if frequencies[freq] > 0 { // It repeats!
				return freq
			}
			frequencies[freq]++
		}
	}
}
