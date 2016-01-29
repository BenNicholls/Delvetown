package data

var itemdata []itemTypeData

//
const (
	ITEM_HEALTH int = iota
	ITEM_POWERUP
	ITEM_SWORD
	ITEM_AXE
	MAX_ITEMTYPES
)

type EquipType int

//equip types. TODO: different types for boots, body, gloves, amulets, whatever
const (
	NOT_EQUIPPABLE EquipType = iota
	EQUIP_WEAPON
	EQUIP_ARMOUR
	MAX_EQUIPTYPES
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
	EQUIP         EquipType
}

func init() {

	itemdata = make([]itemTypeData, MAX_ITEMTYPES)

	itemdata[ITEM_HEALTH] = itemTypeData{"Health", 5, Visuals{0x2b, 0xFF3EED41}, itemFlags{false, true, NOT_EQUIPPABLE}}
	itemdata[ITEM_POWERUP] = itemTypeData{"PowerUp", 5, Visuals{0x18, 0xFFFFFF00}, itemFlags{true, true, NOT_EQUIPPABLE}}
	itemdata[ITEM_SWORD] = itemTypeData{"Sword", 5, Visuals{0x18, 0xFF00FF00}, itemFlags{false, false, EQUIP_WEAPON}}
	itemdata[ITEM_AXE] = itemTypeData{"Axe", 5, Visuals{0x19, 0xFF00FF00}, itemFlags{false, false, EQUIP_WEAPON}}
}
