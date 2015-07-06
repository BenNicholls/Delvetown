package data

type Entity struct {
	X, Y          int
	Name          string
	Enemy         bool
	Health        int
	ID            int
	LightStrength int
	NextTurn      int
	eType         int

	ActionQueue chan Action
}

type Action func(e *Entity)

func (e *Entity) Move(dx, dy int) {
	e.X += dx
	e.Y += dy
}

func (e *Entity) MoveTo(x, y int) {
	e.X = x
	e.Y = y
}

func (e Entity) GetVisuals() Visuals {
	return entitydata[e.eType].vis
}
