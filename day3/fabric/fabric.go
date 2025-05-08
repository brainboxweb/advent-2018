package fabric

func NewPiece() *Piece {
	return &Piece{}
}

type Piece struct {
	areas []area
}

type area struct {
	id int
	x  int
	y  int
	w  int
	h  int
}

func (p *Piece) AddArea(id, x, y, w, h int) {
	area := area{id, x, y, w, h}
	p.areas = append(p.areas, area)
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
