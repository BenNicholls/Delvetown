package modes

import "github.com/bennicholls/delvetown/data"
import "github.com/bennicholls/delvetown/util"

//If the entity e can see the player, step in their direction/attack. Otherwise,
//take a random step.
func (dm *DelveMode) HuntBehaviour(e *data.Entity) data.Action {
	seen := false
	dx, dy := 0, 0

	dm.level.LevelMap.ShadowCast(e.X, e.Y, e.CurStats.SightRange, dm.SeePlayerCast(&seen))

	if seen {
		dx, dy = dm.player.X-e.X, dm.player.Y-e.Y
		if dx < 0 {
			dx = -1
		} else if dx > 0 {
			dx = 1
		}

		if dy < 0 {
			dy = -1
		} else if dy > 0 {
			dy = 1
		}
	} else {
		dx, dy = util.GenerateDirection()
	}

	return dm.AttackMove(e, dx, dy)

}

func (dm *DelveMode) SeePlayerCast(seen *bool) data.Cast {
	return func(m *data.TileMap, x, y, d, r int) {
		if m.GetTile(x, y).Light.Bright > 0 && m.GetEntity(x, y) == dm.player {
			*seen = true
		}
	}
}

//Decide if a move is an atack or not
func (dm *DelveMode) AttackMove(e *data.Entity, dx, dy int) data.Action {
	t := dm.level.LevelMap.GetEntity(e.X+dx, e.Y+dy)

	if t == nil {
		return dm.MoveAction(dx, dy)
	} else if e.Enemy != t.Enemy {
		return dm.AttackAction(dx, dy)
	} else {
		return dm.RestAction()
	}
}
