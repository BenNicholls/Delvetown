package ui

import "github.com/bennicholls/delvetown/console"

type List struct {
	Container
	selected     int
	Highlight    bool
	scrollOffset int
	empty bool
	emptyElem UIElem	
}

func NewList(w, h, x, y, z int, bord bool, empty string) *List {
	c := NewContainer(w, h, x, y, z, bord)
	return &List{*c, 0, true, 0, true, NewTextbox(w, 1, 0, h/2, z, false, true, empty)}
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

//appends an item to the internal list of items
func (l *List) Append(item string) {
	l.Add(NewTextbox(l.width, 1, 0, len(l.Elements), 0, false, false, item))
	l.CheckSelection()
}

//removes the ith item from the internal list of items
func (l *List) Remove(i int) {
	if i < len(l.Elements) && len(l.Elements) != 0 {
		l.Elements = append(l.Elements[:i], l.Elements[i+1:]...)
		l.redraw = true
		l.Calibrate()
		l.CheckSelection()
	}
}

//Ensures list element y values are correct after the list has been tampered with
func (l *List) Calibrate() {
	for i := range l.Elements {
		l.Elements[i].MoveTo(0, i, 0)
	}
}

//Changes the text of the ith item in the internal list of items
func (l *List) Change(i int, item string) {
	l.Elements[i] = NewTextbox(l.width, 1, 0, i, l.z, false, false, item)
}

func (l *List) Render(offset ...int) {
	if l.visible {
		offX, offY, offZ := processOffset(offset)
		
		if len(l.Elements) <= 0 {
			l.emptyElem.Render(l.x+offX, l.y+offY, l.z+offZ)
		} else {
			//calc scrollOffset
			if l.selected < l.scrollOffset {
				l.scrollOffset = l.selected
			} else if l.scrollOffset < l.selected-l.height+1 {
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

			if l.Highlight {
				w, _ := l.Elements[l.selected].GetDims()
				for i := 0; i < w; i++ {
					console.Invert(offX+l.x+i, offY+l.y+l.selected-l.scrollOffset, offZ+l.z)
				}
			}
		}
		
		if l.bordered {
			console.DrawBorder(l.x+offX, l.y+offY, l.z+offZ, l.width, l.height, l.title)
		}
		
		//draw scrollbar
		if len(l.Elements) > l.height {
			console.ChangeGridPoint(offX + l.x + l.width, offY + l.y, offZ + l.z, 0x1e, 0xFFFFFFFF, 0xFF000000)
			console.ChangeGridPoint(offX + l.x + l.width, offY + l.y + l.height - 1, offZ + l.z, 0x1f, 0xFFFFFFFF, 0xFF000000)
				
			sliderHeight := l.height - 2 - (len(l.Elements) - l.height)
			
			for i:= 0; i < sliderHeight; i ++ {
				console.ChangeGridPoint(offX + l.x + l.width, offY + l.y + i + 1 + l.scrollOffset, offZ + l.z, 0xb1, 0xFFFFFFFF, 0xFF000000)
			}				
		}
	}
}
