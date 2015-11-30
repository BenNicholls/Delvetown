package delvemode

import "github.com/bennicholls/delvetown/data"

func (dm *DelveMode) MoveAction(dx, dy int) data.Action {
	return func(e *data.Entity) {
		dm.level.MoveEntity(dx, dy, e)
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
			t.Health -= 5
			e.Health -= 1
			dm.gamelog.AddMessage(e.Name + " attacks " + t.Name + "!")
			if t.Health <= 0 {
				dm.level.RemoveEntity(t.ID)
				dm.gamelog.AddMessage(t.Name + " dies horribly!")
			}
		} else {
			dm.gamelog.AddMessage("No one there to attack, stupid.")
		}
		e.NextTurn += e.AttackSpeed
	}
}
