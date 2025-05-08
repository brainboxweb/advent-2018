package fabric_test

import (
	"testing"

	"github.com/brainboxweb/advent-2018/day3/fabric"
	"github.com/stretchr/testify/assert"
)

func TestOverlapCount(t *testing.T) {
	type inputData struct {
		id int
		x  int
		y  int
		w  int
		h  int
	}

	tests := []struct {
		input    []inputData
		expected int
	}{
		{
			input: []inputData{
				{1, 1, 3, 4, 4},
				{2, 3, 1, 4, 4},
				{3, 5, 5, 2, 2},
			},
			expected: 4,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			fab := fabric.NewPiece()
			for _, data := range tt.input {
				fab.AddArea(data.id, data.x, data.y, data.w, data.h)
			}
			result := fab.OverlapCount()
			assert.Equal(t, tt.expected, result)
		})
	}
}
