package data

import "github.com/bennicholls/delvetown/util"
import "math/rand"

func GenerateCave(p *Entity, w, h int) *Level {

	l := NewLevel(w, h)
	l.SetPlayer(p)

	//fill with walls
	for i := 0; i < w*h; i++ {
		x, y := i%w, i/w
		l.LevelMap.ChangeTileType(x, y, TILE_WALL)
		l.LevelMap.SetVisible(x, y, 0)
	}

	l.seedBranch(w/2, h/2, 300, TILE_CAVEFLOOR)

	//generate some little pools of water
	for i := 0; i < 10; i++ {
		//keep seeds away from the edges (-10, +10)
		l.seedBranch(rand.Intn(w-10)+10, rand.Intn(h-10)+10, 40, TILE_WATER)
	}

	//place more seeds, let 'em grow!
	for i := 0; i < 5; i++ {
		//keep seeds away from the edges (-10, +10)
		l.seedBranch(rand.Intn(w-10)+10, rand.Intn(h-10)+10, 250, TILE_CAVEFLOOR)
	}

	//populate with random enemies
	l.RandomPlaceMobs(20, GNOLL)
	l.RandomPlaceMobs(1, SUPER_GNOLL)
	l.RandomPlaceMobs(2, RAT) //places 2 SWARMS of rats.

	l.PlaceItems(5, ITEM_HEALTH)
	l.PlaceItems(2, ITEM_POWERUP)
	l.PlaceItems(2, ITEM_SWORD)
	l.PlaceItems(2, ITEM_AXE)

	//place the stairs
	for {
		x, y := rand.Intn(w), rand.Intn(h)
		if l.LevelMap.GetTile(x, y).Passable() && l.LevelMap.GetItem(x, y) == nil {
			l.LevelMap.ChangeTileType(x, y, TILE_STAIRS)
			break
		}
	}

	l.SyncClock()

	return l
}

//Randomly places num copies of item in the level
func (l *Level) PlaceItems(num, item int) {
	for i := 0; i < num; {
		x, y := rand.Intn(l.Width), rand.Intn(l.Height)
		if l.LevelMap.GetTile(x, y).Passable() && l.LevelMap.GetItem(x, y) == nil {
			l.LevelMap.AddItem(x, y, NewItem(item))
			i++
		}
	}
}

func (l *Level) RandomPlaceMobs(num, eType int) {
	for i := 0; i < num; {
		x, y := rand.Intn(l.Width), rand.Intn(l.Height)
		if l.LevelMap.GetTile(x, y).Passable() {

			if eType == RAT {
				l.PlaceSwarm(x, y, 10)
			} else {
				l.AddMob(x, y, eType)
			}
			i++
		}
	}
}

func (l *Level) PlaceSwarm(x, y, num int) {
	l.AddMob(x, y, RAT)

	spaces := make([]coord, 0, 150)
	l.LevelMap.ShadowCast(x, y, 7, GetEmptySpacesCast(&spaces))

	for i := 0; i < num; i++ {
		if i == len(spaces) {
			break
		}
		l.AddMob(spaces[i].x, spaces[i].y, RAT)
	}

}

//tile is a data.TILETYPE indicating what we're putting down
func (l *Level) seedBranch(x, y, branch, tile int) {

	l.LevelMap.ChangeTileType(x, y, tile)
	if branch <= 0 {
		return
	}

	//decide num of branches, then branch that many times
	branches := 2 + rand.Intn(5)
	for i := 0; i < branches; i++ {
		dx, dy := util.GenerateDirection()
		//ensure branch doesn't reach edge of map (ugly)
		if !util.CheckBounds(x+dx-1, y+dy-1, l.Width-2, l.Height-2) {
			continue
		} else if t := l.LevelMap.GetTileType(x+dx, y+dy); t != tile && t != TILE_NOTHING {
			l.seedBranch(x+dx, y+dy, branch-branches, tile)
		}
	}
}

func (l *Level) GenerateArena(w, h int) {
	for i := 0; i < l.Width*l.Height; i++ {
		x, y := i%l.Width, i/l.Width
		if x < l.Width/2-w/2 || x > l.Width/2+w/2 || y < l.Height/2-h/2 || y > l.Height/2+h/2 {
			l.LevelMap.ChangeTileType(x, y, TILE_NOTHING)
		} else {
			l.LevelMap.ChangeTileType(x, y, TILE_GRASS)
			if rand.Intn(40) == 0 {
				l.AddMob(x, y, GNOLL)
			}
		}
	}
}
