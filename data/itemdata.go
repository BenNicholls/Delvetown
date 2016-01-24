package data

var itemdata []itemTypeData

//
const (
	ITEM_HEALTH int = iota
	ITEM_POWERUP
	MAX_ITEMTYPES
)

type itemTypeData struct {
	name   string
	weight int
	vis    Visuals
	flags  itemFlags
}

type itemFlags struct {
	USE_ON_PICKUP bool
	CONSUMABLE    bool
}

func init() {

	itemdata = make([]itemTypeData, MAX_ITEMTYPES)

	itemdata[ITEM_HEALTH] = itemTypeData{"Health", 5, Visuals{0x2b, 0xFF3EED41}, itemFlags{false, true}}
	itemdata[ITEM_POWERUP] = itemTypeData{"PowerUp", 5, Visuals{0x18, 0xFFFFFF00}, itemFlags{true, true}}
}
