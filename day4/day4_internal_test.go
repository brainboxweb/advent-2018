package day4

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const dataTimeLayout = "2006-01-02 03:04"

func TestParse(t *testing.T) {
	tests := []struct {
		input          string
		expectedTime   string
		expectedAction string
		expectedID     int
	}{
		{
			"[1518-11-01 00:00] Guard #10 begins shift",
			"1518-11-01 00:00",
			"start",
			10,
		},
		{
			"[1518-11-01 00:05] falls asleep",
			"1518-11-01 00:05",
			"sleep",
			0,
		},
		{
			"[1518-11-01 00:25] wakes up",
			"1518-11-01 00:25",
			"wake",
			0,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			guardData := parse(tt.input)
			expectedTime, err := time.Parse(dataTimeLayout, tt.expectedTime)
			if err != nil {
				panic("not expected")
			}
			assert.Equal(t, expectedTime, guardData.theTime)
			assert.Equal(t, tt.expectedAction, guardData.action)
			assert.Equal(t, tt.expectedID, guardData.id)
		})
	}
}
