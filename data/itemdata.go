package data

var itemdata []itemTypeData

//
const (
	ITEM_HEALTH int = iota
	MAX_ITEMTYPES
)

type itemTypeData struct {
	name   string
	weight int
	vis    Visuals
}

func init() {

	itemdata = make([]itemTypeData, MAX_ITEMTYPES)

	itemdata[ITEM_HEALTH] = itemTypeData{"Health", 5, Visuals{0x2b, 0xFF3EED41}}
}
