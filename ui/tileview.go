package ui

import "github.com/bennicholls/delvetown/console"
import "github.com/bennicholls/delvetown/data"
import "github.com/bennicholls/delvetown/util"

//View object for drawing tiles. (eg. maps). Effectively a buffer for drawing before the console grabs it.
type TileView struct {
	Width, Height int
	x, y, z       int
	bordered      bool
	title         string
	visible       bool

	grid []console.GridCell
}

func NewTileView(w, h, x, y, z int, bord bool) *TileView {
	return &TileView{w, h, x, y, z, bord, "", true, make([]console.GridCell, w*h)}
}

func (tv *TileView) SetTitle(s string) {
	tv.title = s
}

func (tv *TileView) DrawVisuals(x, y int, v data.Visuals) {

	if util.CheckBounds(x, y, tv.Width, tv.Height) {
		tv.grid[y*tv.Width+x].Set(v.Glyph, v.ForeColour, 0x000000, tv.z)
	}
}

//Apply light level. 0-255. TODO: add colour mask (soft orange glow??)
func (tv *TileView) ApplyLight(x, y, b int) {
	if util.CheckBounds(x, y, tv.Width, tv.Height) {
		s := y*tv.Width + x
		if b > 255 {
			b = 255
		} else if b < 80 && b > 0 {
			b = 80 //Baseline brightness for memory... TODO: implement this less magically.
		}
		tv.grid[s].ForeColour = console.ChangeColourAlpha(tv.grid[s].ForeColour, uint8(b))
	}
}

func (tv TileView) Render(offset ...int) {
	if tv.visible {
		offX, offY, offZ := processOffset(offset)
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
}

func (tv TileView) GetDims() (int, int) {
	return tv.Width, tv.Height
}

func (tv *TileView) Clear() {
	for i, _ := range tv.grid {
		tv.grid[i].Set(0, 0, 0, 0)
	}
}

func (tv *TileView) ToggleVisible() {
	tv.visible = !tv.visible
	console.Clear()
}

func (tv *TileView) SetVisibility(v bool) {
	tv.visible = v
	console.Clear()
}
