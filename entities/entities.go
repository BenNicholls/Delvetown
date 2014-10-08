package entities

type Entity struct {
	Glyph int
	X, Y int
	Name string
	Enemy bool
	Health int
}

func (e *Entity) Move(dx, dy int) {
	e.X += dx
	e.Y += dy
}