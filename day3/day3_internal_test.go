package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input         string
		expectedID    int
		expectedPoint point
		expectedArea  area
	}{
		{
			"#1 @ 1,3: 4x4",
			1,
			point{1, 3},
			area{4, 4},
		},
		{
			"#2 @ 3,1: 4x4",
			2,
			point{3, 1},
			area{4, 4},
		},
		{
			"#3 @ 5,5: 2x2",
			3,
			point{5, 5},
			area{2, 2},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			id, p, a := parse(tt.input)
			assert.Equal(t, tt.expectedID, id)
			assert.Equal(t, tt.expectedPoint, p)
			assert.Equal(t, tt.expectedArea, a)
		})
	}
}
