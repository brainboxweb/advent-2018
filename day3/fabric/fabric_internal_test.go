package fabric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsPoint(t *testing.T) {
	tests := []struct {
		input    area
		target   point
		expected bool
	}{
		{
			area{x: 1, y: 1, w: 2, h: 2},
			point{0, 0},
			false,
		},
		{
			area{x: 1, y: 1, w: 2, h: 2},
			point{1, 1},
			true,
		},
		{
			area{x: 1, y: 1, w: 2, h: 2},
			point{2, 2},
			true,
		},
		{
			area{x: 1, y: 1, w: 2, h: 2},
			point{1, 3},
			false,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			aa := tt.input
			result := aa.containsPoint(tt.target)

			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsOverlapped(t *testing.T) {
	type inputData struct {
		id int
		x  int
		y  int
		w  int
		h  int
	}

	tests := []struct {
		input    []inputData
		area     area
		expected bool
	}{
		{
			input: []inputData{
				{1, 1, 3, 4, 4},
				{2, 3, 1, 4, 4},
				{3, 5, 5, 2, 2},
			},
			area:     area{1, 1, 3, 4, 4},
			expected: true,
		},
		{
			input: []inputData{
				{1, 1, 3, 4, 4},
				{2, 3, 1, 4, 4},
				{3, 5, 5, 2, 2},
			},
			area:     area{2, 3, 1, 4, 4},
			expected: true,
		},
		{
			input: []inputData{
				{1, 1, 3, 4, 4},
				{2, 3, 1, 4, 4},
				{3, 5, 5, 2, 2},
			},
			area:     area{3, 5, 5, 2, 2},
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			fab := NewPiece()
			for _, data := range tt.input {
				fab.addToGrid(data.x, data.y, data.w, data.h)
			}

			// spew.Dump(fab)
			result := fab.isOverlapped(tt.area)

			// result := 7
			assert.Equal(t, tt.expected, result)
		})
	}
}
