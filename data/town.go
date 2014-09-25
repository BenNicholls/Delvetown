package data

import "github.com/bennicholls/delvetown/entities"

type Town struct {
	
	Name string

	Townmap *TileMap
	Width, Height int

	entityList []entities.Entity

}

//sets up a bare town object.
func NewTown(w, h int, n string) *Town {
	return &Town{Name: n, Townmap: NewMap(w, h), Width: w, Height: h}
}