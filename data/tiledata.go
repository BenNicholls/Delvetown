package data

var tiledata []tileTypeData

//tiletypes NOTE: Currently capped to 50 tile types. see init()
const (
	TILE_NOTHING = iota
	TILE_GRASS
	TILE_WALL
	TILE_WATER
	MAX_TILETYPES
)

type tileTypeData struct {
	name        string
	passable    bool
	transparent bool
	vis         tileVisuals
}

type tileVisuals struct {
	Glyph      int
	ForeColour uint32
}

func init() {

	//tiledata[TILETYPE]
	tiledata = make([]tileTypeData, 50)

	tiledata[TILE_NOTHING] = tileTypeData{"Nothing", false, true, tileVisuals{0, 0x000000}}
	tiledata[TILE_GRASS] = tileTypeData{"Grass", true, true, tileVisuals{0x2e, 0xFF00FF00}}
	tiledata[TILE_WALL] = tileTypeData{"Wall", false, false, tileVisuals{0x23, 0xFF777777}}
	tiledata[TILE_WATER] = tileTypeData{"Water", false, true, tileVisuals{0xf7, 0xFF0000FF}}
}

//takes tiletype, returns glyph
func GetName(t int) string {
	return tiledata[t].name
}

func IsPassable(t int) bool {
	return tiledata[t].passable
}

func IsTransparent(t int) bool {
	return tiledata[t].transparent
}

func GetVisuals(t int) tileVisuals {
	return tiledata[t].vis
}
