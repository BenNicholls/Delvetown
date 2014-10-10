package data

import "github.com/bennicholls/delvetown/entities"
import "math/rand"

type Level struct {
	Levelmap      *TileMap
	Width, Height int

	Player *entities.Entity

	//Map referencing all enemies in the level. indexed by ID  (found in Entity.ID)
	MobList map[int]*entities.Entity
}

//sets up a bare level object.
func NewLevel(w, h int) *Level {
	l := Level{Levelmap: NewMap(w, h), Width: w, Height: h}
	l.Player = &entities.Entity{2, w / 2, h / 2, "player", false, 50, 0}
	l.MobList = make(map[int]*entities.Entity)
	l.Levelmap.AddEntity(l.Player.X, l.Player.Y, l.Player)
	return &l
}

func (l *Level) MovePlayer(dx, dy int) {
	e := l.Levelmap.GetEntity(l.Player.X+dx, l.Player.Y+dy)
	if e != nil {
		e.Health -= 5
		l.Player.Health -= 5
		if e.Health > 0 {
			return
		}
	}
	t := l.Levelmap.GetTileType(l.Player.X+dx, l.Player.Y+dy)
	if IsPassable(t) {
		l.Levelmap.MoveEntity(l.Player.X, l.Player.Y, dx, dy)
		l.Player.Move(dx, dy)
	}
}

func (l *Level) GenerateRandom() {
	for i := 0; i < l.Width*l.Height; i++ {
		r := rand.Intn(MAX_TILETYPES-1) + 1
		if r != TILE_GRASS {
			r = r - rand.Intn(2)
		}
		l.Levelmap.ChangeTileType(i%l.Width, i/l.Width, r)
	}
}

func (l *Level) GenerateArena(w, h int) {
	for i := 0; i < l.Width*l.Height; i++ {
		x, y := i%l.Width, i/l.Width
		if x < l.Width/2-w/2 || x > l.Width/2+w/2 || y < l.Height/2-h/2 || y > l.Height/2+h/2 {
			l.Levelmap.ChangeTileType(x, y, 0)
		} else {
			l.Levelmap.ChangeTileType(x, y, 1)
			if rand.Intn(40) == 0 {
				l.AddMob(x, y)
			}
		}
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

	e := entities.Entity{15, x, y, "baddie", true, 10, id}
	l.MobList[id] = &e
	l.Levelmap.AddEntity(x, y, &e)
}
