package day2_test

import (
	"testing"

	"github.com/brainboxweb/advent-2018/day2"
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
			"day2_test.txt",
			12,
		},
		{
			"day2.txt",
			6696,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			dataset := helpers.GetDataString(dataPath + "/" + tt.dataFile)
			result := day2.Part1(dataset)
			assert.Equal(t, tt.expected, result)
		})
	}
}
