package data

type Item struct {
	Name     string
	ItemType int
	weight   int
	vis      Visuals
}

func NewItem(itemType int) *Item {
	if itemType < MAX_ITEMTYPES {
		i := itemdata[itemType]
		return &Item{i.name, itemType, i.weight, i.vis}
	} else {
		return nil
	}
}

func (i *Item) GetVisuals() Visuals {
	return i.vis
}
