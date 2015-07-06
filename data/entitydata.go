package data

var entitydata []entityTypeData

const (
	PLAYER = iota
	BUTTS
	MAX_ENTITYTYPES
)

type entityTypeData struct {
	name          string
	hp            int
	enemy         bool
	lightStrength int
	vis           Visuals
}

func init() {

	entitydata = make([]entityTypeData, MAX_ENTITYTYPES)

	entitydata[PLAYER] = entityTypeData{"Player", 100, false, 15, Visuals{2, 0xFFFFFFFF}}
	entitydata[BUTTS] = entityTypeData{"Butts", 10, true, 7, Visuals{15, 0xFFFF0000}}
}

func NewEntity(x, y, id, eType int) *Entity {

	if eType < MAX_ENTITYTYPES {
		e := entitydata[eType]
		return &Entity{x, y, e.name, e.enemy, e.hp, id, e.lightStrength, 1, eType, make(chan Action, 20)}
	} else {
		return nil
	}

}
