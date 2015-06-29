package delvemode

import "github.com/bennicholls/delvetown/data"

//TODO: infer dur from entity's speed and whatnot
func (dm *DelveMode) MoveAction(dx, dy, dur int) data.Action {
	return func(e *data.Entity) {
		dm.level.MovePlayer(dx, dy)
		dm.step += 1
		e.NextTurn += dur
	}
}

func (dm *DelveMode) AttackAction(dx, dy, dur int) data.Action {
	return func(e *data.Entity) {
		t := dm.level.LevelMap.GetEntity(e.X+dx, e.Y+dy)
		if t != nil {
			t.Health -= 5
			e.Health -= 1
			dm.gamelog.AddMessage(e.Name + " attacks " + t.Name + "! TO VALHALLA!!!")
			if t.Health <= 0 {
				dm.level.RemoveEntity(t.ID)
				dm.gamelog.AddMessage("You are a MURDERER!")
			}
		} else {
			dm.gamelog.AddMessage("No one there to attack, stupid.")
		}
		dm.step += 1
		e.NextTurn += dur
	}
}
