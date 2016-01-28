package data

var tiledata []tileTypeData

//default tiletypes POSSIBLE TODO: dynamic changing tile properties? Think about this.
const (
	TILE_NOTHING = iota
	TILE_GRASS
	TILE_WALL
	TILE_WATER
	TILE_CAVEFLOOR
	TILE_STAIRS
	MAX_TILETYPES
)

type tileTypeData struct {
	name        string
	passable    bool
	transparent bool
	vis         Visuals
}

type Visuals struct {
	Glyph      int
	ForeColour uint32
}

func init() {

	//tiledata[TILETYPE]
	tiledata = make([]tileTypeData, MAX_TILETYPES)

	tiledata[TILE_NOTHING] = tileTypeData{"Nothing", false, true, Visuals{0, 0x000000}}
	tiledata[TILE_GRASS] = tileTypeData{"Grass", true, true, Visuals{0x2e, 0xFF00FF00}}
	tiledata[TILE_WALL] = tileTypeData{"Wall", false, false, Visuals{0x23, 0xFF777777}}
	tiledata[TILE_WATER] = tileTypeData{"Water", false, true, Visuals{0xf7, 0xFF0000FF}}
	tiledata[TILE_CAVEFLOOR] = tileTypeData{"Cave Floor", true, true, Visuals{0x2e, 0xFF746253}}
	tiledata[TILE_STAIRS] = tileTypeData{"Stairs", true, false, Visuals{0x1e, 0xFFFFFFFF}}
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

func GetTileVisuals(t int) Visuals {
	return tiledata[t].vis
}
