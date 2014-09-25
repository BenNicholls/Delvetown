package data

import "github.com/bennicholls/delvetown/entities"

type TileMap struct {
	width, height int
	tiles []Tile
}

func NewMap(w, h int) *TileMap {
	return &TileMap{width: w, height: h, tiles: make([]Tile, w*h)}
}

func (m *TileMap) ChangeTileType(x, y, tile int) {
	if x < m.width && y < m.height {
		m.tiles[y*m.width + x].tileType = tile
	}
}

func (m *TileMap) GetTileType(x, y int) int {
	if x < m.width && y < m.height {
		return m.tiles[y*m.width + x].tileType
	} else {
		return 0
	}
}

//Basic unit for the world. Holds a type (grass, wall, etc), a list of contained items
//(dropped weapons, also furniture), and a pointer to an Entity if one is standing there
//Eventually will hold pathfinding information too.
type Tile struct {
	tileType, variant int
	passable bool
	entity *entities.Entity
}