package ui

import "strings"

type Inputbox struct {
	Text   Textbox
	cursor int
}

func NewInputbox(w, h, x, y, z int, bord bool) *Inputbox {
	return &Inputbox{Textbox{w, h, x, y, z, bord, "", "", true}, 0}
}

func (ib *Inputbox) SetTitle(t string) {
	ib.Text.title = t
}

func (ib *Inputbox) Render(offset ...int) {
	switch len(offset) {
	case 0:
		ib.Text.Render()
	case 2:
		ib.Text.Render(offset[0], offset[1])
	case 3:
		ib.Text.Render(offset[0], offset[1], offset[2])
	}
}

func (ib Inputbox) GetDims() (int, int) {
	return ib.Text.width, ib.Text.height
}

func (ib *Inputbox) ToggleVisible() {
	ib.Text.ToggleVisible()

}

func (ib *Inputbox) MoveCursor(dx, dy int) {
	ib.cursor += dx
	ib.cursor += dy * ib.Text.width
	if ib.cursor < 0 {
		ib.cursor = 0
	} else if ib.cursor > len(ib.Text.text)+1 {
		ib.cursor = ib.Text.width * ib.Text.height
	}
}

func (ib *Inputbox) Insert(s string) {
	if len(ib.Text.text) >= ib.Text.width*ib.Text.height {
		return
	}
	if ib.cursor == len(ib.Text.text) {
		ib.Text.ChangeText(ib.Text.text + s)
	} else {
		t := []string{ib.Text.text[0:ib.cursor], s, ib.Text.text[ib.cursor:]}
		ib.Text.ChangeText(strings.Join(t, ""))
	}
	ib.cursor += 1
}

func (ib *Inputbox) Delete() {

	switch len(ib.Text.text) {
	case 0:
		return
	case 1:
		ib.Text.ChangeText("")
	default:
		t := []string{ib.Text.text[0 : ib.cursor-1], ib.Text.text[ib.cursor:]}
		ib.Text.ChangeText(strings.Join(t, ""))
	}

	if ib.cursor > 0 {
		ib.cursor -= 1
	}
}
