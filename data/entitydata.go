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

	entitydata[PLAYER] = entityTypeData{"Player", 100, false, 15, 40, 3, 5, 5, Visuals{0x02, 0xFFFFFFFF}}
	entitydata[BUTTS] = entityTypeData{"Butts", 10, true, 7, 10, 5, 7, 3, Visuals{0x42, 0xFFFF0000}}
}
