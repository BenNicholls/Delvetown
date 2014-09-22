package modes

import "github.com/bennicholls/delvetown/tilemap"
import "github.com/bennicholls/delvetown/console"

type TownMode struct {
	width, height int
	townMap *tilemap.Map
}

func NewTown(w, h int) *TownMode {
	town := TownMode{townMap: tilemap.NewMap(w, h)}
	town.width = w
	town.height = h
	return &town
}

func (t *TownMode) Update() {
	t.townMap.ChangeTileType(10, 10, 2)
}

func (t *TownMode)  Render() {
	for i := 0; i < t.width*t.height; i++ {
		if t.townMap.GetTileType(i%t.width, i/t.width) == 2 {
			console.ChangeSquare(i%t.width, i/t.width, 50, 0xFFFFFFFF, 0)
		}
	}
	console.Render()
}