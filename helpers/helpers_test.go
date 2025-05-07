package helpers_test

import (
	"testing"

	"github.com/brainboxweb/advent-2018/helpers"
	"github.com/stretchr/testify/assert"
)


func TestGetDataString(t *testing.T) {
	tests := []struct {
		datafile string
		expected []string
	}{
		{
			"test.txt",
			[]string{"hello", "world"},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			result := helpers.GetDataString(tt.datafile)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetDataInt(t *testing.T) {
	tests := []struct {
		name              string
		datafile          string
		expected          []int
		expectedErrorFunc func(assert.TestingT, any, ...any) bool
	}{
		{
			"happy path",
			"test_int.txt",
			[]int{1234, 4567},
			assert.Nil,
		},
		{
			"invalid data",
			"test_int_err.txt",
			nil,
			assert.NotNil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := helpers.GetDataInt(tt.datafile)
			assert.Equal(t, tt.expected, result)
			tt.expectedErrorFunc(t, err)
		})
	}
}
