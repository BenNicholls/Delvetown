package data

var tiledata []tileTypeData

//tiletypes NOTE: Currently capped to 50 tile types. see init()
const (
	TILE_GRASS = iota
	TILE_WALL
	MAX_TILETYPES
)

type tileTypeData struct {
	name string
	passable bool
	vis tileVisuals
}

type tileVisuals struct {
	Glyph int
	ForeColour uint32
	BackColour uint32
}

func init() {
	
	//tiledata[TILETYPE]
	tiledata = make([]tileTypeData, 50)

	tiledata[TILE_GRASS] = tileTypeData{"Grass", true, tileVisuals{176, 0x00FF00, 0x000000}}
	tiledata[TILE_WALL] = tileTypeData{"Wall", false, tileVisuals{178, 0x333333, 0x000000}}
}

//takes tiletype, returns glyph
func GetName(t int) string {
	return tiledata[t].name
}

func IsPassable(t int) bool {
	return tiledata[t].passable
}

func GetVisuals(t int) tileVisuals {
	return tiledata[t].vis
}

