package day3

import (
	"strconv"
	"strings"

	"github.com/brainboxweb/advent-2018/day3/fabric"
)

func Part1(data []string) int {
	myFabric := populate(data)
	return myFabric.OverlapCount()
}

func populate(data []string) *fabric.Piece {
	myFabric := fabric.NewPiece()
	for _, line := range data {
		id, point, area := parse(line)
		myFabric.AddArea(id, point.x, point.y, area.x, area.y)
	}
	return myFabric
}

type point struct {
	x int
	y int
}

type area struct {
	x int
	y int
}

func parse(input string) (int, point, area) {
	parts := strings.Split(input, " ")

	// id
	idString := parts[0]
	idString = strings.Trim(idString, "#")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic("not expected")
	}

	// Position
	posnString := parts[2]
	posnString = strings.Trim(posnString, ":")

	ppParts := strings.Split(posnString, ",")
	x, err := strconv.Atoi(ppParts[0])
	if err != nil {
		panic("not expected")
	}
	y, err := strconv.Atoi(ppParts[1])
	if err != nil {
		panic("not expected")
	}

	// Area
	areaString := parts[3]

	aParts := strings.Split(areaString, "x")
	w, err := strconv.Atoi(aParts[0])
	if err != nil {
		panic("not expected")
	}
	h, err := strconv.Atoi(aParts[1])
	if err != nil {
		panic("not expected")
	}

	return id, point{x, y}, area{w, h}
}
