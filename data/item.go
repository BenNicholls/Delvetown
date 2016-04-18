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

//returns whether item can be equipped
func (i Item) Equippable() bool {
	return i.Flags.EQUIP != NOT_EQUIPPABLE
}
