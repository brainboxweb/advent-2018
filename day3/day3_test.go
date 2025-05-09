package day3_test

import (
	"testing"

	"github.com/brainboxweb/advent-2018/day3"
	"github.com/brainboxweb/advent-2018/helpers"
	"github.com/stretchr/testify/assert"
)

const dataPath = "../testdata"

func TestPart1(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			"day3_test.txt",
			4,
		},
		{
			"day3.txt",
			119551,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			dataset := helpers.GetDataString(dataPath + "/" + tt.dataFile)
			result := day3.Part1(dataset)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			"day3_test.txt",
			3,
		},
		{
			"day3.txt",
			1124,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			dataset := helpers.GetDataString(dataPath + "/" + tt.dataFile)
			result := day3.Part2(dataset)
			assert.Equal(t, tt.expected, result)
		})
	}
}
