package data

import "github.com/bennicholls/delvetown/entities"

type TileMap struct {
	width, height int
	tiles         []Tile
}

func NewMap(w, h int) *TileMap {
	return &TileMap{width: w, height: h, tiles: make([]Tile, w*h)}
}

func (m *TileMap) ChangeTileType(x, y, tile int) {
	if x < m.width && y < m.height {
		m.tiles[y*m.width+x].tileType = tile
	}
}

func (m *TileMap) GetTileType(x, y int) int {
	if x < m.width && y < m.height && x >= 0 && y >= 0 {
		return m.tiles[y*m.width+x].tileType
	} else {
		return 0
	}
}

func (m *TileMap) AddEntity(x, y int, e *entities.Entity) {
	if x < m.width && y < m.height && x >= 0 && y >= 0 {
		m.tiles[x+y*m.width].Entity = e
	}
}

func (m *TileMap) RemoveEntity(x, y int) {
	if x < m.width && y < m.height && x >= 0 && y >= 0 {
		m.tiles[x+y*m.width].Entity = nil
	}
}

func (m *TileMap) MoveEntity(x, y, dx, dy int) {
	m.AddEntity(x+dx, y+dy, m.tiles[x+y*m.width].Entity)
	m.RemoveEntity(x, y)
}

//Basic unit for the world. Holds a type (grass, wall, etc), a list of contained items
//(dropped weapons, also furniture), and a pointer to an Entity if one is standing there
//Eventually will hold pathfinding information too.
type Tile struct {
	tileType, variant int
	passable          bool
	Entity            *entities.Entity
}
