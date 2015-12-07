package data

type Item struct {
	Name   string
	weight int
	vis    Visuals
}

func (i *Item) GetVisuals() Visuals {
	return i.vis
}
