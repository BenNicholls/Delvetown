package data

import "math/rand"

type Level struct {
	LevelMap      *TileMap
	MemoryMap     *TileMap
	Width, Height int

	Player *Entity

	//Map referencing all enemies in the level. indexed by ID  (found in Entity.ID)
	MobList map[int]*Entity
}

//sets up a bare level object.
func NewLevel(w, h int) *Level {
	l := Level{LevelMap: NewMap(w, h), MemoryMap: NewMap(w, h), Width: w, Height: h}
	l.Player = NewEntity(w/2, h/2, 0, PLAYER)
	l.MobList = make(map[int]*Entity)
	l.LevelMap.AddEntity(l.Player.X, l.Player.Y, l.Player)
	return &l
}

func (l *Level) MoveEntity(dx, dy int, e *Entity) {

	if dx == 0 && dy == 0 {
		return
	}
	//move entity if tile is passable
	if l.LevelMap.GetTile(e.X+dx, e.Y+dy).Passable() {
		l.LevelMap.MoveEntity(e.X, e.Y, dx, dy)
		e.Move(dx, dy)
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

	e := NewEntity(x, y, id, BUTTS)
	l.MobList[id] = e
	l.LevelMap.AddEntity(x, y, e)
}
