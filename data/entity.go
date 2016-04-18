package data

//equip slots. each slot is like a special inventory slot. TODO: more slots.
const (
	SLOT_WEAPON int = iota
	SLOT_ARMOUR
	MAX_SLOTS
)

type Entity struct {
	X, Y                   int
	Name                   string
	Enemy                  bool
	Health                 int
	ID                     int
	LightStrength          int
	SightRange             int
	NextTurn               int
	EType                  int
	MoveSpeed, AttackSpeed int
	BaseAttack             int

	Inventory []*Item
	Equipment []*Item //indexed by the SLOT enum above

	ActionQueue chan Action
}

type Action func(e *Entity)

func NewEntity(x, y, id, eType int) *Entity {

	if eType < MAX_ENTITYTYPES {
		e := entitydata[eType]
		//Max Inventory space is 30 for now. POSSIBLE: dynamically sized inventory? (bags, stronger, whatever)
		return &Entity{x, y, e.name, e.enemy, e.hp, id, e.lightStrength, e.sightRange, 1, eType, e.mv, e.av, e.at, make([]*Item, 0, 30), make([]*Item, MAX_SLOTS), make(chan Action, 20)}
	} else {
		return nil
	}
}

func (e *Entity) Move(dx, dy int) {
	e.X += dx
	e.Y += dy
}

func (e *Entity) MoveTo(x, y int) {
	e.X = x
	e.Y = y
}

func (e Entity) GetVisuals() Visuals {
	return entitydata[e.EType].vis
}

//This is going to do some heavy lifting someday.
func (e Entity) CalcAttack() int {
	return e.BaseAttack
}

//Removes item from inventory at index i
func (e *Entity) RemoveItem(i int) {
	//from Slicetricks... ensures the removed item can be garbage collected
	//if consumed instead of dropped
	if len(e.Inventory) > 1 {
		e.Inventory, e.Inventory[len(e.Inventory)-1] = append(e.Inventory[:i], e.Inventory[i+1:]...), nil
	} else {
		e.Inventory = make([]*Item, 0, 30)
	}
}

//TODO: throw error or drop or something if inventory limit reached
func (e *Entity) AddItem(i *Item) {
	if len(e.Inventory) < cap(e.Inventory) {
		e.Inventory = append(e.Inventory, i)
	}
}

//equip item at index i. TODO: throw errors or something if equip fails
//For now, only call this if you've checked if the item is equippable first
func (e *Entity) EquipItem(i int) {
	equipItem := e.Inventory[i]
	var targetSlot int

	switch equipItem.Flags.EQUIP {
	case EQUIP_WEAPON:
		targetSlot = SLOT_WEAPON
	case EQUIP_ARMOUR:
		targetSlot = SLOT_ARMOUR
	}

	if e.Equipment[targetSlot] != nil {
		e.Inventory[i] = e.Equipment[targetSlot]
	} else {
		return
	}

	e.Equipment[targetSlot] = equipItem
}

//Returns name of item in equipped slot, or "empty"
func (e Entity) GetEquipmentName(slot int) string {
	if e.Equipment[slot] != nil {
		return e.Equipment[slot].Name
	} else {
		return "empty"
	}
}
