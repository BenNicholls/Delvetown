package entities

type Entity struct {
	Glyph  int
	X, Y   int
	Name   string
	Enemy  bool
	Health int
	ID     int
}

func (e *Entity) Move(dx, dy int) {
	e.X += dx
	e.Y += dy
}
