// package day4 is Day 4 of the Advent of Code
package day4

import (
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/brainboxweb/advent-2018/day4/closet"
)

const (
	actionStart = "start"
	actionSleep = "sleep"
	actionWake  = "wake"
)

func Part1(data []string) int {
	guardDataset := prepData(data)
	myCloset := closet.New()
	var guard *closet.Guard
	for _, data := range guardDataset {
		if data.id != 0 {
			guard = myCloset.AddGuard(data.id)
			continue
		}
		err := guard.AddAction(data.action, data.theTime)
		if err != nil {
			panic("not expected")
		}
	}
	myCloset.ProcessTimes()
	guardID := myCloset.GetLongestSleep()
	mostCommon := myCloset.GetMostCommonSleep(guardID)

	return guardID * mostCommon
}

func prepData(data []string) []guardData {
	guardDataset := []guardData{}
	for _, line := range data {
		gData := parse(line)
		guardDataset = append(guardDataset, gData)
	}
	sort.Slice(guardDataset, func(i, j int) bool {
		return guardDataset[i].theTime.Before(guardDataset[j].theTime)
	})
	return guardDataset
}

type guardData struct {
	id      int
	theTime time.Time
	action  string
}

func parse(input string) guardData {
	parts := strings.Split(input, "]")
	timePart := strings.Trim(parts[0], "[]")
	layout := "2006-01-02 15:04" // to const
	theTime, err := time.Parse(layout, timePart)
	if err != nil {
		panic("not expected")
	}
	// instruction
	parts2 := strings.Split(input, " ")
	var action string
	switch parts2[2] {
	case "Guard":
		action = actionStart
	case "falls":
		action = actionSleep
	case "wakes":
		action = actionWake
	}
	// Guard id
	id := 0
	if action == actionStart {
		idPart := parts2[3]
		idPart = strings.Trim(idPart, "#")
		id, err = strconv.Atoi(idPart)
		if err != nil {
			panic("not expected")
		}
	}
	return guardData{id, theTime, action}
}
