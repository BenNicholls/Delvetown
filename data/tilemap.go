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

func (m TileMap) GetTileType(x, y int) int {
	if x < m.width && y < m.height && x >= 0 && y >= 0 {
		return m.tiles[y*m.width+x].tileType
	} else {
		return 0
	}
}

func (m TileMap) GetTile(x, y int) Tile {
	if x < m.width && y < m.height && x >= 0 && y >= 0 {
		return m.tiles[y*m.width+x]
	} else {
		return Tile{}
	}
}

func (m *TileMap) SetTile(x, y int, t Tile) {
	if x < m.width && y < m.height && x >= 0 && y >= 0 {
		m.tiles[x+y*m.width] = t
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

func (m TileMap) GetEntity(x, y int) *entities.Entity {
	if x < m.width && y < m.height && x >= 0 && y >= 0 {
		return m.tiles[x+y*m.width].Entity
	} else {
		return nil
	}
}

//For testing purposes.
func (m *TileMap) ChangeTileColour(x, y, c int) {
	if x < m.width && y < m.height && x >= 0 && y >= 0 {
		m.tiles[x+y*m.width].TestColour = c
	}
}

func (m TileMap) LastVisible(x, y int) int {
	if x < m.width && y < m.height && x >= 0 && y >= 0 {
		return m.tiles[x+y*m.width].lastSeen
	} else {
		return 0
	}
}

//NOTE: Consider renaming this.
func (m *TileMap) SetVisible(x, y, tick int) {
	if x < m.width && y < m.height && x >= 0 && y >= 0 {
		m.tiles[x+y*m.width].lastSeen = tick
	}
}

//Basic unit for the world. Holds a type (grass, wall, etc), a list of contained items
//(dropped weapons, also furniture), and a pointer to an Entity if one is standing there
//Eventually will hold pathfinding information too.
type Tile struct {
	tileType, variant int
	passable          bool
	Entity            *entities.Entity
	TestColour        int // NOTE: DELETE THIS SOMEDAY.
	mask              tileVisuals
	lastSeen          int // Records the last tick that this tile was seen
}

func (t Tile) Type() int {
	return t.tileType
}
