package ui

import "github.com/bennicholls/delvetown/console"
import "github.com/bennicholls/delvetown/data"

type UIElem interface {
	Render(offset ...int)
	GetDims() (int, int)
}

type Container struct {
	width, height int
	x, y, z int
	bordered bool

	Elements []UIElem
}

func NewContainer(w,h,x,y,z int, bord bool) *Container {
	return &Container{w,h,x,y,z,bord, make([]UIElem, 0, 20)}
}

func (c *Container) Add(elem UIElem) {
	c.Elements = append(c.Elements, elem)
}

func (c *Container) Render() {
	for i := 0; i < len(c.Elements); i++ {
		c.Elements[i].Render(c.x, c.y)
	}

	if c.bordered {
			console.DrawBorder(c.x, c.y, c.z, c.width, c.height)
		}
}

//TODO: Split this into separate files for each UI element
type Textbox struct {
	width, height int
	x, y, z int
	bordered bool

	text string
	dirty bool
}

//TODO: sanity checks.
func NewTextbox(w, h, x, y, z int, bord bool, txt string) *Textbox {
	return &Textbox{w, h, x, y, z, bord, txt, true}
}

//TODO: validate that 't' only includes ascii characters (rune < 255 i think)
func (t *Textbox) ChangeText(txt string) {
	if t.text != txt {
		t.text = txt
		t.dirty = true
	}
}

//TODO: word wrap. scroll bar? (maybe a "MORE" prompt might be easier), separate dirty flag for the border?
//Render function optionally takes an offset (for containering), strictly 2 ints.
func (t *Textbox) Render(offset ...int) {
	if t.dirty{
		offX, offY := 0, 0
		if len(offset) == 2 {
			offX, offY = offset[0], offset[1]
		}

		i, r := 0, rune(0)
		for i, r = range t.text {
			if i >= t.width*t.height {
				break
			}
			console.ChangeGridPoint(offX + t.x + i%t.width, offY + t.y + i/t.width, t.z, int(r), 0xFFFFFF, 0x000000)
		}
		for i++ ; i < t.width*t.height; i++ {
			console.ChangeGridPoint(offX + t.x + i%t.width, offY + t.y + i/t.width, t.z, 0, 0x000000, 0x000000)
		}

		if t.bordered {
			console.DrawBorder(offX + t.x, offY + t.y, t.z, t.width, t.height)
		}
		t.dirty = false
	}	
}

func (t Textbox) GetDims() (int, int) {
	return t.width, t.height
}

//View object for drawing tiles. (eg. maps). Effectively a buffer for drawing before the console grabs it.
type TileView struct {
	Width, Height int
	x, y, z int
	bordered, dirty bool

	grid []console.GridCell
}

func NewTileView(w, h, x, y, z int, bord bool) *TileView {
	return &TileView{w, h, x, y, z, bord, true, make([]console.GridCell, w*h)}
}

//takes (x,y) and a tiletype 
func (tv *TileView) DrawTile(x, y, t int) {

	if x < tv.Width && y < tv.Height {
			v := data.GetVisuals(t)
			tv.grid[y * tv.Width + x].Set(v.Glyph, v.ForeColour, 0x000000, tv.z)
			tv.dirty = true
	}
		
}

func (tv *TileView) DrawEntity(x, y, g int, c uint32) {

	if x < tv.Width && y < tv.Height {
			tv.grid[y * tv.Width + x].Set(g, c, 0x000000, tv.z)
			tv.dirty = true
	}
		
}

func (tv TileView) Render() {
	for i, p := range tv.grid {
		if p.Dirty {
			console.ChangeGridPoint(tv.x + i%tv.Width, tv.y + i/tv.Width, tv.z, p.Glyph, p.ForeColour, 0x000000)
			p.Dirty = false
		}
	}
	if tv.bordered {
		console.DrawBorder(tv.x, tv.y, tv.z, tv.Width, tv.Height)
	}
}

func (tv TileView) GetDims() (int, int) {
	return tv.Width, tv.Height
}