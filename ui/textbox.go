package ui

import "github.com/bennicholls/delvetown/console"

//Area for displyaing text.
type Textbox struct {
	width, height int
	x, y, z       int
	bordered      bool
	title         string
	text          string
	dirty         bool
}

//TODO: sanity checks.
func NewTextbox(w, h, x, y, z int, bord bool, txt string) *Textbox {
	return &Textbox{w, h, x, y, z, bord, "", txt, true}
}

func (t *Textbox) SetTitle(s string) {
	t.title = s
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
	if t.dirty {
		offX, offY, offZ := 0, 0, 0
		if len(offset) >= 2 {
			offX, offY = offset[0], offset[1]
			if len(offset) == 3 {
				offZ = offset[2]
			}
		}

		i, r := 0, rune(0)
		for i, r = range t.text {
			if i >= t.width*t.height {
				break
			}
			console.ChangeGridPoint(offX+t.x+i%t.width, offY+t.y+i/t.width, t.z+offZ, int(r), 0xFFFFFF, 0x000000)
		}
		for i++; i < t.width*t.height; i++ {
			console.ChangeGridPoint(offX+t.x+i%t.width, offY+t.y+i/t.width, t.z+offZ, 0, 0x000000, 0x000000)
		}

		if t.bordered {
			console.DrawBorder(offX+t.x, offY+t.y, t.z+offZ, t.width, t.height, t.title)
		}
		t.dirty = false
	}
}

func (t Textbox) GetDims() (int, int) {
	return t.width, t.height
}
