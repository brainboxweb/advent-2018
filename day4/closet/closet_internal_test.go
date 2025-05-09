package closet

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddAction(t *testing.T) {
	tests := []struct {
		guardID  int
		theTime  time.Time // don't need
		action   string
		expected Guard
	}{
		{
			guardID: 77,
			theTime: time.Date(2009, 11, 17, 20, 22, 0, 0, time.UTC),
			action:  actionSleep,
			expected: Guard{
				id: 77,
				actions: []action{
					{
						min: 22,
						act: "sleep",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			guard := Guard{id: tt.guardID}
			err := guard.AddAction(tt.action, tt.theTime)
			if err != nil {
				panic("not expected")
			}
			assert.Equal(t, tt.expected, guard)
		})
	}
}

func TestProcessTimes(t *testing.T) {
	tests := []struct {
		guard    Guard
		expected []int
	}{
		{
			guard: Guard{
				id: 77,
				actions: []action{
					{
						min: 5,
						act: "sleep",
					},
					{
						min: 25,
						act: "wake",
					},
				},
			},
			expected: []int{5, 6, 7, 8, 9, 10,
				11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			g := tt.guard
			g.processTimes()
			assert.Equal(t, tt.expected, g.sleeps)
		})
	}
}
