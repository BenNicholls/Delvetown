package ui

import "github.com/bennicholls/delvetown/console"
import "github.com/bennicholls/delvetown/data"

//View object for drawing tiles. (eg. maps). Effectively a buffer for drawing before the console grabs it.
type TileView struct {
	Width, Height   int
	x, y, z         int
	bordered, dirty bool
	title           string

	grid []console.GridCell
}

func NewTileView(w, h, x, y, z int, bord bool) *TileView {
	return &TileView{w, h, x, y, z, bord, true, "", make([]console.GridCell, w*h)}
}

func (tv *TileView) SetTitle(s string) {
	tv.title = s
}

//takes (x,y) and a tiletype
func (tv *TileView) DrawTile(x, y int, t data.Tile) {

	if x < tv.Width && y < tv.Height {
		v := data.GetVisuals(t.Type())
		tv.grid[y*tv.Width+x].Set(v.Glyph, v.ForeColour, console.MakeColour(t.TestColour, 0, 0), tv.z)
		tv.dirty = true
	}

}

func (tv *TileView) DrawEntity(x, y, g int, c uint32) {

	if x < tv.Width && y < tv.Height {
		tv.grid[y*tv.Width+x].Set(g, c, 0x000000, tv.z)
		tv.dirty = true
	}

}

func (tv TileView) Render(offset ...int) {
	offX, offY, offZ := 0, 0, 0
	if len(offset) >= 2 {
		offX, offY = offset[0], offset[1]
		if len(offset) == 3 {
			offZ = offset[2]
		}
	}
	for i, p := range tv.grid {
		if p.Dirty {
			console.ChangeGridPoint(tv.x+offX+i%tv.Width, tv.y+offY+i/tv.Width, tv.z+offZ, p.Glyph, p.ForeColour, p.BackColour)
			p.Dirty = false
		}
	}
	if tv.bordered {
		console.DrawBorder(tv.x+offX, tv.y+offY, tv.z+offZ, tv.Width, tv.Height, tv.title)
	}
}

func (tv TileView) GetDims() (int, int) {
	return tv.Width, tv.Height
}

func (tv *TileView) Clear() {
	for i, _ := range tv.grid {
		tv.grid[i].Set(0, 0, 0, 0)
	}
}
