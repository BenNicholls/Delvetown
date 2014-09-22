package map

import "github.com/bennicholls/delvetown/entities"

type Map struct {
	width, height int
	tileMap []Tile
}

type Tile struct {
	tileType, variant int
	passable bool
	entity Entity
}