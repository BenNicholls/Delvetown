package data

var entitydata []entityTypeData

const (
	PLAYER = iota
	BUTTS  //the main enemy is the infamous butts. Super mature.
	SUPER_BUTTS
	BUTT_SWARM
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
	entitydata[BUTTS] = entityTypeData{"Butts", 15, true, 7, 10, 5, 7, 3, Visuals{0x42, 0xFFFF0000}}
	entitydata[SUPER_BUTTS] = entityTypeData{"Super Butts", 40, true, 10, 10, 10, 8, 10, Visuals{0xE1, 0xFFFF0000}}
	entitydata[BUTT_SWARM] = entityTypeData{"Butt Swarmer", 5, true, 4, 7, 1, 1, 1, Visuals{0x62, 0xFFFF0000}}
}
