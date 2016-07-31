package ui

import "github.com/bennicholls/delvetown/console"

//Area for displaying text.
type Textbox struct {
	width, height int
	x, y, z       int
	bordered      bool
	centered      bool
	title         string
	text          string
	visible       bool
}

//TODO: sanity checks.
func NewTextbox(w, h, x, y, z int, bord, cent bool, txt string) *Textbox {
	return &Textbox{w, h, x, y, z, bord, cent, "", txt, true}
}

func (t *Textbox) SetTitle(s string) {
	t.title = s
}

//TODO: validate that 't' only includes ascii characters (rune < 255 i think)
func (t *Textbox) ChangeText(txt string) {
	if t.text != txt {
		t.text = txt
	}
}

//TODO: word wrap. scroll bar? (maybe a "MORE" prompt might be easier), separate dirty flag for the border?
//Render function optionally takes an offset (for containering), 2 or 3 ints.
func (t *Textbox) Render(offset ...int) {
	if t.visible {
		offX, offY, offZ := processOffset(offset)

		if t.bordered {
			console.DrawBorder(offX+t.x, offY+t.y, t.z+offZ, t.width, t.height, t.title)
		}

		//fill textbox with background colour
		for i := len(t.text); i < t.width*t.height; i++ {
			console.ChangeGridPoint(offX+t.x+i%t.width, offY+t.y+i/t.width, t.z+offZ, 0, 0xFFFFFFFF, 0xFF000000)
		}

		//offset if centerred
		if t.centered {
			offX += t.width/2 - len(t.text)/2
		}

		//print text
		for i, r := range t.text {
			if i >= t.width*t.height {
				break
			}
			console.ChangeGridPoint(offX+t.x+i%t.width, offY+t.y+i/t.width, t.z+offZ, int(r), 0xFFFFFFFF, 0xFF000000)
		}
	}
}

func (t Textbox) GetDims() (int, int) {
	return t.width, t.height
}

func (t *Textbox) ToggleVisible() {
	t.visible = !t.visible
	console.Clear()
}

func (t *Textbox) SetVisibility(v bool) {
	t.visible = v
	console.Clear()
}

func (t *Textbox) MoveTo(x, y, z int) {
	t.x = x
	t.y = y
	t.z = z
}
