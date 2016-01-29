package ui

import "github.com/bennicholls/delvetown/console"

type List struct {
	Container
	selected     int
	Highlight    bool
	scrollOffset int
}

func NewList(w, h, x, y, z int, bord bool) *List {
	c := NewContainer(w, h, x, y, z, bord)
	return &List{*c, 0, true, 0}
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
	l.Highlight = !l.Highlight
}

//Ensures Selected item is not out of bounds.
func (l *List) CheckSelection() {
	if l.selected < 0 {
		l.selected = 0
	} else if l.selected >= len(l.Elements) {
		l.selected = len(l.Elements) - 1
	}
}

func (l *List) Render(offset ...int) {
	if l.visible {
		offX, offY, offZ := processOffset(offset)

		//calc scrollOffset
		if l.selected < l.scrollOffset {
			l.scrollOffset = l.selected
		} else if l.scrollOffset < l.selected-l.height {
			l.scrollOffset = l.selected - l.height + 1
		}

		if l.redraw {
			console.Clear(l.x+offX, l.y+offY, l.width, l.height)
			l.redraw = false
		}
		for i := l.scrollOffset; i < l.scrollOffset+l.height; i++ {
			if i >= len(l.Elements) {
				break
			}
			l.Elements[i].Render(l.x+offX, l.y+offY-l.scrollOffset, l.z+offZ)
		}

		if l.bordered {
			console.DrawBorder(l.x+offX, l.y+offY, l.z+offZ, l.width, l.height, l.title)
		}

		if len(l.Elements) > 0 && l.Highlight {
			w, _ := l.Elements[l.selected].GetDims()
			for i := 0; i < w; i++ {
				console.Invert(offX+l.x+i, offY+l.y+l.selected-l.scrollOffset, offZ+l.z)
			}
		}
	}
}
