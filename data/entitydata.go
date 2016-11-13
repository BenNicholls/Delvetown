package data

var entitydata []entityTypeData

const (
	PLAYER = iota
	GNOLL  
	SUPER_GNOLL
	RAT
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
	entitydata[GNOLL] = entityTypeData{"Gnoll", true, Stats{15, 10, 3, 5, 7, 5, 10, 7, 1, 1, 1}, Visuals{0x67, 0xFFFF0000}}
	entitydata[SUPER_GNOLL] = entityTypeData{"Super Gnoll", true, Stats{40, 10, 10, 10, 8, 10, 10, 10, 1, 1, 1}, Visuals{0x47, 0xFFFF0000}}
	entitydata[RAT] = entityTypeData{"Rat", true, Stats{5, 10, 1, 1, 1, 1, 7, 4, 1, 1, 1}, Visuals{0x72, 0xFFFF0000}}
}
