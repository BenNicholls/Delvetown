package data

type Item struct {
	Name     string
	ItemType int
	Weight   int
	Flags    itemFlags
}

func NewItem(itemType int) *Item {
	if itemType < MAX_ITEMTYPES {
		i := itemdata[itemType]
		return &Item{i.name, itemType, i.weight, i.flags}
	} else {
		return nil
	}
}

func (i *Item) GetVisuals() Visuals {
	return itemdata[i.ItemType].vis
}
