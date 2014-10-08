package data

import "github.com/bennicholls/delvetown/entities"
import "math/rand"

type Level struct {
	
	Levelmap *TileMap
	Width, Height int

	Player *entities.Entity

	MobList []entities.Entity
}

//sets up a bare level object.
func NewLevel(w, h int) *Level {
	l := Level{Levelmap: NewMap(w, h), Width: w, Height: h}
	l.Player = &entities.Entity{2, w/2, h/2, "player", false, 50}
	l.MobList = make([]entities.Entity, 50)
	l.Levelmap.AddEntity(l.Player.X, l.Player.Y, l.Player)
	return &l
}

func (l *Level) MovePlayer(dx, dy int) {
	e := l.Levelmap.GetEntity(l.Player.X + dx, l.Player.Y + dy)
	if e != nil {
		e.Health -= 5
		if e.Health > 0 {
			return
		}
	}
	t := l.Levelmap.GetTileType(l.Player.X + dx, l.Player.Y + dy)
	if IsPassable(t) {
		l.Levelmap.MoveEntity(l.Player.X, l.Player.Y, dx, dy)
		l.Player.Move(dx, dy)
	}
}

func (l *Level) GenerateRandom() {
	for i := 0; i < l.Width*l.Height; i++ {
		r := rand.Intn(MAX_TILETYPES - 1) + 1
		if r != TILE_GRASS {
			r = r - rand.Intn(2)
		}
		l.Levelmap.ChangeTileType(i%l.Width, i/l.Width, r)
	}
}

func (l *Level) GenerateArena(w, h int) {
	for i:=0; i < l.Width*l.Height; i++ {
		x, y := i%l.Width, i/l.Width
		if x < l.Width/2 - w/2 || x > l.Width/2 + w/2 || y < l.Height/2 - h/2 || y > l.Height/2 + h/2 {
			l.Levelmap.ChangeTileType(x, y, 0)
		} else {
			l.Levelmap.ChangeTileType(x, y, 1)
			if rand.Intn(40) == 0 {
				l.AddMob(x, y)
			}
		}
	}
}

func (l *Level) AddMob(x, y int) {
	e := entities.Entity{15, x, y, "baddie", true, 10}
	l.MobList = append(l.MobList, e)
	l.Levelmap.AddEntity(x, y, &e)
}