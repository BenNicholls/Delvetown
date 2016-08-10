package ui

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
	NONE int = iota
	ACTIVATE //used for buttons I guess?
	CHANGE   //used when a UIelem is changed
)

type Event struct {
	Caller UIElem
	ID int
}

var EventStream chan *Event

func init() {
	//TODO: 50 event limit... is that a problem??
	EventStream = make(chan *Event, 50)
}

func NewEvent(c UIElem, id int) *Event {
	return &Event{c, id}
}

func NextEvent() *Event {
	if len(EventStream) > 0 { 
		e := <- EventStream
		return e
	} else {
		return nil
	}
}