package entities

type Entity struct {
	eType int
	x, y int
}

func (e *Entity) Move(dx, dy int) {
	e.x += dx
	e.y += dy
}