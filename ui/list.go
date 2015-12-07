package ui

import "github.com/bennicholls/delvetown/console"

type List struct {
	Container
	selected  int
	highlight bool
}

func NewList(w, h, x, y, z int, bord bool) *List {
	c := NewContainer(w, h, x, y, z, bord)
	return &List{*c, 0, true}
}

func (l *List) Select(s int) {
	if s < len(l.Elements) {
		l.selected = s
	}
}

func (l *List) Next() {
	//small list protection
	if len(l.Elements) <= 1 {
		l.selected = 0
		return
	}

	if l.selected >= len(l.Elements)-1 {
		l.selected = 0
	} else {
		l.selected++
	}
}

func (l *List) Prev() {
	//small list protection
	if len(l.Elements) <= 1 {
		l.selected = 0
		return
	}

	if l.selected == 0 {
		l.selected = len(l.Elements) - 1
	} else {
		l.selected--
	}
}

func (l List) GetSelection() int {
	return l.selected
}

func (l *List) ToggleHighlight() {
	l.highlight = !l.highlight
}

func (l *List) Render(offset ...int) {
	if l.visible {
		offX, offY, offZ := processOffset(offset)

		l.Container.Render(offX, offY, offZ)

		if len(l.Elements) > 0 && l.highlight {
			w, _ := l.Elements[l.selected].GetDims()
			for i := 0; i < w; i++ {
				console.Invert(offX+l.x+i, offY+l.y+l.selected, offZ+l.z)
			}
		}
	}
}
