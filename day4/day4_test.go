package day4_test

import (
	"testing"

	"github.com/brainboxweb/advent-2018/day4"
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
			"day4_test.txt",
			240,
		},
		{
			"day4.txt",
			12169,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			dataset := helpers.GetDataString(dataPath + "/" + tt.dataFile)
			result := day4.Part1(dataset)
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
			"day4_test.txt",
			4455,
		},
		{
			"day4.txt",
			16164,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			dataset := helpers.GetDataString(dataPath + "/" + tt.dataFile)
			result := day4.Part2(dataset)
			assert.Equal(t, tt.expected, result)
		})
	}
}
