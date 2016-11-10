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
	name      string
	enemy     bool
	baseStats Stats
	vis       Visuals
}

func init() {

	entitydata = make([]entityTypeData, MAX_ENTITYTYPES)

	entitydata[PLAYER] = entityTypeData{"Player", false, Stats{100, 10, 5, 5, 5, 3, 40, 15, 1, 1, 1}, Visuals{0x02, 0xFFFFFFFF}}
	entitydata[BUTTS] = entityTypeData{"Butts", true, Stats{15, 10, 3, 5, 7, 5, 10, 7, 1, 1, 1}, Visuals{0x42, 0xFFFF0000}}
	entitydata[SUPER_BUTTS] = entityTypeData{"Super Butts", true, Stats{40, 10, 10, 10, 8, 10, 10, 10, 1, 1, 1}, Visuals{0xE1, 0xFFFF0000}}
	entitydata[BUTT_SWARM] = entityTypeData{"Butt Swarmer", true, Stats{5, 10, 1, 1, 1, 1, 7, 4, 1, 1, 1}, Visuals{0x62, 0xFFFF0000}}
}
