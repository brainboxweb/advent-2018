// Package closet represents closet
package closet

import (
	"time"

	"errors"
)

const (
	actionSleep = "sleep"
	actionWake  = "wake"
)

func New() *Closet {
	gg := make(map[int]*Guard)
	return &Closet{guards: gg}
}

type Closet struct {
	guards map[int]*Guard
}

type Guard struct {
	id      int
	actions []action
	sleeps  []int
}

type action struct {
	min int
	act string
}

func (c *Closet) AddGuard(id int) *Guard {
	val, ok := c.guards[id]
	if ok {
		return val
	}
	guard := &Guard{id: id}
	c.guards[id] = guard
	return guard
}

func (c *Closet) ProcessTimes() {
	for _, guard := range c.guards {
		guard.processTimes()
	}
}

func (c *Closet) GetLongestSleep() int {
	var maxSleep int
	var maxSleepGuard int
	for id, guard := range c.guards {
		sleepLength := len(guard.sleeps)
		if sleepLength > maxSleep {
			maxSleep = sleepLength
			maxSleepGuard = id
		}
	}
	return maxSleepGuard
}

func (c *Closet) GetMostCommonSleep(guardID int) int {
	g := c.guards[guardID]
	sleepLog := make(map[int]int) // minute, count
	for _, m := range g.sleeps {
		sleepLog[m]++
	}
	maxMn := 0
	maxCount := 0
	for mn, count := range sleepLog {
		if count > maxCount {
			maxMn = mn
			maxCount = count
		}
	}
	return maxMn
}

func (c *Closet) MostFrequent() (maxSleepGuard, maxSleepMinute int) {
	allMinutes := c.buildMinuteMap()
	var maxSleepCount int
	for mn, thing := range allMinutes {
		for guideID, count := range thing {
			if count > maxSleepCount {
				maxSleepCount = count
				maxSleepGuard = guideID
				maxSleepMinute = mn
			}
		}
	}
	return maxSleepGuard, maxSleepMinute
}

func (c *Closet) buildMinuteMap() map[int]map[int]int {
	allMinutes := make(map[int]map[int]int) // minute, guard, frequency
	for mn := 0; mn < 60; mn++ {
		guardMn := make(map[int]int) // guard, frequency
		for _, guard := range c.guards {
			for _, sleep := range guard.sleeps {
				if sleep == mn {
					guardMn[guard.id]++
				}
			}
		}
		allMinutes[mn] = guardMn
	}
	return allMinutes
}

// --- Guard

func (g *Guard) AddAction(act string, theTime time.Time) error {
	if !(act == actionSleep || act == actionWake) {
		return errors.New("unknown action")
	}
	g.actions = append(g.actions, action{min: theTime.Minute(), act: act})
	return nil
}

func (g *Guard) processTimes() {
	var startSleep int
	for _, aa := range g.actions {
		if aa.act == actionSleep {
			startSleep = aa.min // starting value
			continue
		}
		if aa.act == actionWake {
			for i := startSleep; i < aa.min; i++ {
				g.sleeps = append(g.sleeps, i)
			}
			continue
		}
	}
}
