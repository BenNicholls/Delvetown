package data

import "github.com/bennicholls/delvetown/data/entities"
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
	l.Player = &entities.Entity{2, w / 2, h / 2, "player", false, 50, 0, 0xFFFFFFFF, 15}
	l.MobList = make(map[int]*entities.Entity)
	l.LevelMap.AddEntity(l.Player.X, l.Player.Y, l.Player)
	return &l
}

func (l *Level) MovePlayer(dx, dy int) {

	//move player if tile is passable
	t := l.LevelMap.GetTileType(l.Player.X+dx, l.Player.Y+dy)
	if IsPassable(t) {
		l.LevelMap.ShadowCast(l.Player.X, l.Player.Y, l.Player.LightStrength, Darken)
		l.LevelMap.MoveEntity(l.Player.X, l.Player.Y, dx, dy)
		l.Player.Move(dx, dy)
		l.LevelMap.ShadowCast(l.Player.X, l.Player.Y, l.Player.LightStrength, Light)
	}
}

func (l *Level) MoveMob(ID, dx, dy int) {

	//move player if tile is passable
	e := l.MobList[ID]
	if e != nil {
		t := l.LevelMap.GetTileType(e.X+dx, e.Y+dy)
		if IsPassable(t) {
			l.LevelMap.ShadowCast(e.X, e.Y, e.LightStrength, Darken)
			l.LevelMap.MoveEntity(e.X, e.Y, dx, dy)
			e.Move(dx, dy)
			l.LevelMap.ShadowCast(e.X, e.Y, e.LightStrength, Light)
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

	e := entities.Entity{15, x, y, "butts", true, 10, id, 0xFFFF0000, 7}
	l.MobList[id] = &e
	l.LevelMap.AddEntity(x, y, &e)
	l.LevelMap.ShadowCast(e.X, e.Y, e.LightStrength, Light)
}
