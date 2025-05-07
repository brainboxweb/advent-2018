package helpers

import (
	"bufio"
	"os"
	"strconv"
)

func GetDataString(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var ret []string
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	return ret
}

func GetDataInt(filename string) ([]int, error) {
	data := GetDataString(filename)
	var ret []int
	for _, line := range data {
		val, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		ret = append(ret, val)
	}

	return ret, nil
}
