package data

import "github.com/bennicholls/delvetown/util"
import "math/rand"

func (l *Level) GenerateRandom() {
	for i := 0; i < l.Width*l.Height; i++ {
		r := rand.Intn(MAX_TILETYPES-1) + 1
		if r != TILE_GRASS {
			r = r - rand.Intn(2)
		}
		l.LevelMap.ChangeTileType(i%l.Width, i/l.Width, r)
	}
}

func (l *Level) GenerateArena(w, h int) {
	for i := 0; i < l.Width*l.Height; i++ {
		x, y := i%l.Width, i/l.Width
		if x < l.Width/2-w/2 || x > l.Width/2+w/2 || y < l.Height/2-h/2 || y > l.Height/2+h/2 {
			l.LevelMap.ChangeTileType(x, y, 0)
		} else {
			l.LevelMap.ChangeTileType(x, y, 1)
			if rand.Intn(40) == 0 {
				l.AddMob(x, y)
			}
		}
	}
}

func (l *Level) GenerateCave() {

	l.LevelMap.RemoveEntity(l.Player.X, l.Player.Y)

	//fill with walls
	for i := 0; i < l.Width*l.Height; i++ {
		x, y := i%l.Width, i/l.Width
		l.LevelMap.ChangeTileType(x, y, TILE_WALL)
		l.LevelMap.SetVisible(x, y, 0)
	}

	l.seedBranch(l.Width/2, l.Height/2, 300, TILE_CAVEFLOOR)
	l.Player.MoveTo(l.Width/2, l.Height/2)
	l.LevelMap.AddEntity(l.Width/2, l.Height/2, l.Player)
	//place  more seeds, let 'em grow!
	for i := 0; i < 5; i++ {
		//keep seeds away from the edges (-10, +10)
		l.seedBranch(rand.Intn(l.Width-10)+10, rand.Intn(l.Height-10)+10, 200, TILE_CAVEFLOOR)
	}

	//generate some little pools of water
	for i := 0; i < 10; i++ {
		//keep seeds away from the edges (-10, +10)
		l.seedBranch(rand.Intn(l.Width-10)+10, rand.Intn(l.Height-10)+10, 40, TILE_WATER)
	}

	//populate with random enemies
	for i := 0; i < 20; i++ {
		x, y := rand.Intn(l.Width), rand.Intn(l.Height)
		if l.LevelMap.GetTile(x, y).Passable() {
			l.AddMob(x, y)
		} else {
			i -= 1
		}
	}
}

//tile is a data.TILETYPE indicating what we're putting down
func (l *Level) seedBranch(x, y, branch, tile int) {

	l.LevelMap.ChangeTileType(x, y, tile)
	if branch <= 0 {
		return
	}

	//decide num of branches, then branch that many times
	branches := 5
	for i := 0; i < branches; i++ {
		dx, dy := util.GenerateDirection()
		//ensure branch doesn't reach edge of map (ugly)
		if !util.CheckBounds(x+dx-1, y+dy-1, l.Width-2, l.Height-2) {
			continue
		} else if l.LevelMap.GetTileType(x+dx, y+dy) == TILE_WALL {
			l.seedBranch(x+dx, y+dy, branch-branches, tile)
		}
	}
}
