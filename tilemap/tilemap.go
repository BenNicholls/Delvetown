package tilemap

import "github.com/bennicholls/delvetown/entities"

type Map struct {
	width, height int
	tileMap []Tile
}

func NewMap(w, h int) *Map {
	return &Map{width: w, height: h, tileMap: make([]Tile, w*h)}
}

func (m *Map) ChangeTileType(x, y, tile int) {
	m.tileMap[y*m.width + x].tileType = tile
}

func (m *Map) GetTileType(x, y int) int {
	return m.tileMap[y*m.width + x].tileType
}

type Tile struct {
	tileType, variant int
	passable bool
	entity entities.Entity
}