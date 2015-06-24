package ui

import "github.com/bennicholls/delvetown/util"
import "strings"
import "strconv"

//TODO: String longer than size of Textbox. How hard could that be????

//Inputbox is based on textbox.
type Inputbox struct {
	Textbox
	cursor int
	anims  []Animator
}

func NewInputbox(w, h, x, y, z int, bord bool) *Inputbox {

	ib := &Inputbox{Textbox{w, h, x, y, z, bord, "", "", true}, 0, make([]Animator, 0, 20)}
	ib.anims = append(ib.anims, NewBlinkAnimation(0, 0, 30))

	return ib
}

func (ib *Inputbox) MoveCursor(dx, dy int) {
	ib.cursor += dx
	ib.cursor += dy * ib.width
	if ib.cursor < 0 {
		ib.cursor = 0
	} else if ib.cursor > len(ib.text)+1 {
		ib.cursor = ib.width * ib.height
	}
}

func (ib *Inputbox) Insert(s string) {
	if len(ib.text) >= ib.width*ib.height {
		return
	}
	if ib.cursor == len(ib.text) {
		ib.ChangeText(ib.text + s)
	} else {
		t := []string{ib.text[0:ib.cursor], s, ib.text[ib.cursor:]}
		ib.ChangeText(strings.Join(t, ""))
	}
	ib.cursor += 1
}

func (ib *Inputbox) Delete() {

	switch len(ib.text) {
	case 0:
		return
	case 1:
		ib.ChangeText("")
	default:
		t := []string{ib.text[0 : ib.cursor-1], ib.text[ib.cursor:]}
		ib.ChangeText(strings.Join(t, ""))
	}

	if ib.cursor > 0 {
		ib.MoveCursor(-1, 0)
	}
}

func (ib *Inputbox) Reset() {
	ib.ChangeText("")
	ib.cursor = 0
}

//takes a key representing a letter and inserts. TODO: capital support
func (ib *Inputbox) InsertText(key rune) {
	if !util.ValidText(key) {
		return
	}
	s := strconv.QuoteRuneToASCII(key)
	s, _ = strconv.Unquote(s)
	ib.Insert(s)
}

func (ib Inputbox) GetText() string {
	return ib.text
}

func (ib *Inputbox) Render(offset ...int) {
	if ib.visible {
		offX, offY, offZ := processOffset(offset)

		ib.Textbox.Render(offX, offY, offZ)

		for i, _ := range ib.anims {
			ib.anims[i].Tick()
			ib.anims[i].Render(ib.x+offX+ib.cursor, ib.y+offY, ib.z+offZ)
		}
	}
}