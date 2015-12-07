package data

var entitydata []entityTypeData

const (
	PLAYER = iota
	BUTTS  //the main enemy is the infamous butts. Super mature.
	MAX_ENTITYTYPES
)

type entityTypeData struct {
	name          string
	hp            int
	enemy         bool
	lightStrength int
	sightRange    int
	mv, av        int //movespeed and attackspeed
	at            int
	vis           Visuals
}

func init() {

	entitydata = make([]entityTypeData, MAX_ENTITYTYPES)

	entitydata[PLAYER] = entityTypeData{"Player", 100, false, 15, 40, 3, 5, 5, Visuals{2, 0xFFFFFFFF}}
	entitydata[BUTTS] = entityTypeData{"Butts", 10, true, 7, 10, 5, 7, 3, Visuals{15, 0xFFFF0000}}
}

func NewEntity(x, y, id, eType int) *Entity {

	if eType < MAX_ENTITYTYPES {
		e := entitydata[eType]
		//Max Inventory space is 30 for now. POSSIBLE: dynamically sized inventory? (bags, stronger, whatever)
		return &Entity{x, y, e.name, e.enemy, e.hp, id, e.lightStrength, e.sightRange, 1, eType, e.mv, e.av, e.at, make([]*Item, 0, 30), make(chan Action, 20)}
	} else {
		return nil
	}
}
