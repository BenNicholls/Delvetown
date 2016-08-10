package ui

import "github.com/bennicholls/delvetown/util"

type UIElem interface {
	Render(offset ...int)
	GetDims() (w int, h int)
	SetTitle(title string)
	ToggleVisible()
	SetVisibility(v bool)
	MoveTo(x, y, z int)
}

func processOffset(offset []int) (x, y, z int) {
	x, y, z = 0, 0, 0
	if len(offset) >= 2 {
		x, y = offset[0], offset[1]
		if len(offset) == 3 {
			z = offset[2]
		}
	}
	return
}

//event IDs
const (
	NONE     int = iota
	ACTIVATE     //used for buttons I guess?
	CHANGE       //used when a UIelem is changed
)

type Event struct {
	Caller  UIElem
	ID      int
	Message string
}

var EventStream chan *Event

func init() {
	EventStream = make(chan *Event, 100)
}

func PushEvent(c UIElem, id int, m string) {
	if len(EventStream) == cap(EventStream) {
		ClearEvents()
		util.LogError("UI Eventstream limit reached! FLUSHY FLUSHY.")
	}

	EventStream <- &Event{c, id, m}
}

func ClearEvents() {
	EventStream = make(chan *Event, 100)
}

func PopEvent() *Event {
	if len(EventStream) > 0 {
		e := <-EventStream
		return e
	} else {
		return nil
	}
}
