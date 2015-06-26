package entities

type Entity struct {
	Glyph         int
	X, Y          int
	Name          string
	Enemy         bool
	Health        int
	ID            int
	Fore          uint32
	LightStrength int
}

func (e *Entity) Move(dx, dy int) {
	e.X += dx
	e.Y += dy
}

func (e *Entity) MoveTo(x, y int) {
	e.X = x
	e.Y = y
}
