// package fabric represents fabric layout
package fabric

func NewPiece() *Piece {
	grid := make(map[point]int)
	return &Piece{
		grid:     grid,
		minPoint: point{x: 100000, y: 100000},
		maxPoint: point{x: 0, y: 0},
	}
}

type Piece struct {
	grid     map[point]int
	areas    []area
	minPoint point
	maxPoint point
}

type area struct {
	id int
	x  int
	y  int
	w  int
	h  int
}
type point struct {
	x int
	y int
}

func (p *Piece) AddArea(id, x, y, w, h int) {
	p.setMinMax(x, y, w, h)
	p.addToGrid(x, y, w, h) // For Part 2
	aa := area{id, x, y, w, h}
	p.areas = append(p.areas, aa)
}

func (p *Piece) addToGrid(x int, y int, w int, h int) {
	for yy := y; yy < (y + h); yy++ {
		for xx := x; xx < (x + w); xx++ {
			p.grid[point{x: xx, y: yy}]++
		}
	}
}

func (p *Piece) setMinMax(x int, y int, w int, h int) {
	if x < p.minPoint.x {
		p.minPoint.x = x
	}
	if y < p.minPoint.y {
		p.minPoint.y = y
	}
	if x+w > p.maxPoint.x {
		p.maxPoint.x = x + w
	}
	if y+h > p.maxPoint.y {
		p.maxPoint.y = y + h
	}
}

func (p *Piece) OverlapCount() int {
	type point = struct {
		X int
		Y int
	}
	grid := make(map[point]int)
	for _, area := range p.areas {
		for yy := area.y; yy < (area.y + area.h); yy++ {
			for xx := area.x; xx < (area.x + area.w); xx++ {
				grid[point{xx, yy}]++
			}
		}
	}
	total := 0
	for _, cnt := range grid {
		if cnt > 1 {
			total++
		}
	}
	return total
}

func (p *Piece) FindNotOverlappedArea() int {
	var candidateID int
	for _, myArea := range p.areas {
		if p.isOverlapped(myArea) {
			continue
		}
		candidateID = myArea.id
	}
	return candidateID
}

func (p *Piece) isOverlapped(myArea area) bool {
	for yy := myArea.y; yy < (myArea.y + myArea.h); yy++ {
		for xx := myArea.x; xx < (myArea.x + myArea.w); xx++ {
			targetPoint := point{x: xx, y: yy}
			if p.grid[targetPoint] > 1 {
				return true
			}
		}
	}
	return false
}

func (aa area) containsPoint(candidate point) bool {
	if candidate.x >= aa.x && candidate.x < aa.x+aa.w &&
		candidate.y >= aa.y && candidate.y < aa.y+aa.h {
		return true
	}
	return false
}
