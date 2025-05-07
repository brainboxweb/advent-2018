package day1

func Part1(data []int) int {
	res := 0
	for _, num := range data {
		res += num
	}
	return res
}
