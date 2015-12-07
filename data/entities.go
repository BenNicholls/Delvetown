package data

type Entity struct {
	X, Y                   int
	Name                   string
	Enemy                  bool
	Health                 int
	ID                     int
	LightStrength          int
	SightRange             int
	NextTurn               int
	EType                  int
	MoveSpeed, AttackSpeed int
	BaseAttack             int

	Inventory []*Item

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
	return entitydata[e.EType].vis
}

//This is going to do some heavy lifting someday.
func (e Entity) CalcAttack() int {
	return e.BaseAttack
}
