package data

type Item struct {
	name   string
	weight int
	vis    Visuals
}

func (i *Item) GetVisuals() Visuals {
	return i.vis
}
