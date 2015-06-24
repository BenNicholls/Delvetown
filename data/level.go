package data

import "github.com/bennicholls/delvetown/entities"
import "github.com/bennicholls/delvetown/util"
import "math/rand"

type Level struct {
	LevelMap      *TileMap
	MemoryMap     *TileMap
	Width, Height int

	Player *entities.Entity

	//Map referencing all enemies in the level. indexed by ID  (found in Entity.ID)
	MobList map[int]*entities.Entity
}

//sets up a bare level object.
func NewLevel(w, h int) *Level {
	l := Level{LevelMap: NewMap(w, h), MemoryMap: NewMap(w, h), Width: w, Height: h}
	l.Player = &entities.Entity{2, w / 2, h / 2, "player", false, 50, 0, 0xFFFFFFFF}
	l.MobList = make(map[int]*entities.Entity)
	l.LevelMap.AddEntity(l.Player.X, l.Player.Y, l.Player)
	return &l
}

func (l *Level) MovePlayer(dx, dy int) {

	//move player if tile is passable
	t := l.LevelMap.GetTileType(l.Player.X+dx, l.Player.Y+dy)
	if IsPassable(t) {
		l.LevelMap.MoveEntity(l.Player.X, l.Player.Y, dx, dy)
		l.Player.Move(dx, dy)
	}
}

func (l *Level) MoveMob(ID, dx, dy int) {

	//move player if tile is passable
	e := l.MobList[ID]
	if e != nil {
		t := l.LevelMap.GetTileType(e.X+dx, e.Y+dy)
		if IsPassable(t) {
			l.LevelMap.MoveEntity(e.X, e.Y, dx, dy)
			e.Move(dx, dy)
		}
	}
}

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

	//fill with walls
	for i := 0; i < l.Width*l.Height; i++ {
		x, y := i%l.Width, i/l.Width
		l.LevelMap.ChangeTileType(x, y, 2)
		l.LevelMap.ChangeTileColour(x, y, 0)
		l.LevelMap.SetVisible(x, y, 0)
	}

	l.seedBranch(l.Width/2, l.Height/2, 300)
	l.LevelMap.RemoveEntity(l.Player.X, l.Player.Y)
	l.Player.MoveTo(l.Width/2, l.Height/2)
	l.LevelMap.AddEntity(l.Width/2, l.Height/2, l.Player)

	for i := 0; i < 5; i++ {
		l.seedBranch(rand.Intn(l.Width), rand.Intn(l.Height), 200)
	}

}

func (l *Level) seedBranch(x, y, branch int) {

	l.LevelMap.ChangeTileType(x, y, 1)
	if branch <= 0 {
		return
	}

	//decide num of branches, then branch that many times
	branches := 4
	for i := 0; i < branches; i++ {
		dx, dy := util.GenerateDirection()
		if x < 0 || x >= l.Width || y < 0 || y >= l.Height {
			continue
		} else if l.LevelMap.GetTileType(x+dx, y+dy) > 1 {
			l.seedBranch(x+dx, y+dy, branch-branches)
		}
	}
}

func (l *Level) RemoveEntity(id int) {
	e := l.MobList[id]
	if e != nil {
		l.LevelMap.RemoveEntity(e.X, e.Y)
		delete(l.MobList, id)
	}
}

//Creates a new entity and add it to the list. Generates id as well.
func (l *Level) AddMob(x, y int) {

	//generate new unique id. loop checks for previous entity with that id.
	id := rand.Int()
	for _, ok := l.MobList[id]; ok; {
		id = rand.Int()
		_, ok = l.MobList[id]
	}

	e := entities.Entity{15, x, y, "butts", true, 10, id, 0xFFFF0000}
	l.MobList[id] = &e
	l.LevelMap.AddEntity(x, y, &e)
}
