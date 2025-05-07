package day1_test

import (
	"testing"

	"github.com/brainboxweb/advent-2018/day1"
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
			"day1.txt",
			466,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			dataset, err := helpers.GetDataInt(dataPath + "/" + tt.dataFile)
			if err != nil {
				panic("not expected")
			}
			result := day1.Part1(dataset)
			assert.Equal(t, tt.expected, result)
		})
	}
}
