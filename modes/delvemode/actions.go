package delvemode

import "strconv"
import "github.com/bennicholls/delvetown/data"

func (dm *DelveMode) MoveAction(dx, dy int) data.Action {
	return func(e *data.Entity) {
		dm.level.MoveEntity(dx, dy, e)

		//item pickup check
		if dm.level.LevelMap.GetItem(e.X, e.Y) != nil {
			dm.level.LevelMap.RemoveItem(e.X, e.Y)
			e.Health += 10
			dm.gamelog.AddMessage(e.Name + " picks up the health!")
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
