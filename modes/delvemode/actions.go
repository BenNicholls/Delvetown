package delvemode

import "strconv"
import "github.com/bennicholls/delvetown/data"

func (dm *DelveMode) MoveAction(dx, dy int) data.Action {
	return func(e *data.Entity) {
		dm.level.MoveEntity(dx, dy, e)

		if dm.level.LevelMap.GetTileType(e.X, e.Y) == data.TILE_STAIRS {
			dm.level.GenerateCave()
			return
		}

		//item pickup check
		if item := dm.level.LevelMap.GetItem(e.X, e.Y); item != nil {
			dm.level.LevelMap.RemoveItem(e.X, e.Y)

			if item.Flags.USE_ON_PICKUP {
				action := dm.UseItem(e, item.ItemType)
				action(e)
			} else {
				e.Inventory = append(e.Inventory, item)
				if e == dm.player {
					dm.UpdateHUDInventory()
				}
				dm.gamelog.AddMessage(e.Name + " picks up the " + item.Name + "!")
			}
		}
		e.NextTurn += e.MoveSpeed
	}
}

func (dm *DelveMode) AttackAction(dx, dy int) data.Action {
	return func(e *data.Entity) {
		t := dm.level.LevelMap.GetEntity(e.X+dx, e.Y+dy)
		if t != nil {
			damage := e.CalcAttack()
			t.Health -= damage
			attackMessage := e.Name + " attacks " + t.Name + " for " + strconv.Itoa(damage) + " damage!"
			if t.Health <= 0 {
				dm.level.RemoveEntity(t.ID)
				attackMessage += " " + t.Name + " dies horribly!"
			}
			dm.gamelog.AddMessage(attackMessage)

		} else {
			dm.gamelog.AddMessage("No one there to attack, stupid.")
		}
		e.NextTurn += e.AttackSpeed
	}
}

//Do nothing, try again next turn?
func (dm *DelveMode) RestAction() data.Action {
	return func(e *data.Entity) {
		e.NextTurn += 1
	}
}

func (dm *DelveMode) UseInventory(e *data.Entity, selection int) data.Action {
	item := e.Inventory[selection]

	//consume the item
	if item.Flags.CONSUMABLE {
		e.RemoveItem(selection)
		if e == dm.player {
			dm.UpdateHUDInventory()
		}
	}

	if item.Flags.EQUIP == data.NOT_EQUIPPABLE {
		return dm.UseItem(e, item.ItemType)
	} else {
		return dm.UseEquip(e, selection)
	}
}

func (dm *DelveMode) UseItem(e *data.Entity, itemType int) data.Action {
	switch itemType {
	case data.ITEM_HEALTH:
		return func(e *data.Entity) {
			dm.gamelog.AddMessage(e.Name + " uses the Health!")
			e.Health += 10
			e.NextTurn += 5 //hardcoded speed for consuming this item. TODO: not hardcode
		}
	case data.ITEM_POWERUP:
		return func(e *data.Entity) {
			dm.gamelog.AddMessage(e.Name + " powers up!")
			e.BaseAttack += 5
		}
	}

	//This only gets returned if you try to use an item that's not defined in itemdata.go
	//This should never happen, but if it does I guess you get penalized a turn for
	//screwing up so damn bad.
	return dm.RestAction()
}

func (dm *DelveMode) UseEquip(e *data.Entity, selection int) data.Action {
	return func(e *data.Entity) {
		dm.gamelog.AddMessage(e.Name + " equips the " + e.Inventory[selection].Name + "!")
		e.EquipItem(selection)
		if e == dm.player {
			dm.UpdateHUDInventory()
		}
	}
}

func (dm *DelveMode) DropInventoryItem(e *data.Entity, selection int) data.Action {
	return func(e *data.Entity) {
		item := e.Inventory[selection]

		if item != nil {
			if dm.level.DropItem(e.X, e.Y, item) {
				e.RemoveItem(selection)
				dm.UpdateHUDInventory()
			} else {
				dm.gamelog.AddMessage(item.Name + " could not be dropped!")
			}
		}
	}
}
