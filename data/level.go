package data

import "math/rand"
import "github.com/bennicholls/delvetown/util"

type Level struct {
	LevelMap      *TileMap
	MemoryMap     *TileMap
	Width, Height int

	Player *Entity

	Floor int

	//Map referencing all enemies in the level. indexed by ID  (found in Entity.ID)
	MobList map[int]*Entity
}

//sets up a bare level object.
func NewLevel(w, h int) *Level {
	l := Level{LevelMap: NewMap(w, h), MemoryMap: NewMap(w, h), Width: w, Height: h}
	l.MobList = make(map[int]*Entity)
	return &l
}

func (l *Level) ResetLevel() {
	l.LevelMap = NewMap(l.Width, l.Height)
	l.MemoryMap = NewMap(l.Width, l.Height)
	l.MobList = make(map[int]*Entity)
}

func (l *Level) SetPlayer(p *Entity) {

	//remove current player and reset memory if necessary
	if l.Player != nil {
		l.LevelMap.RemoveEntity(l.Player.X, l.Player.Y)
		l.MemoryMap = NewMap(l.Width, l.Height)
	}
	
	l.Player = p
	l.Player.X, l.Player.Y = l.Width/2, l.Height/2 //until we get starting location code in
	l.LevelMap.AddEntity(l.Player.X, l.Player.Y, l.Player)

}

//ensures the enemies' turn counters are synchronized with the player's
func (l *Level) SyncClock() {
	for id := range l.MobList {
		l.MobList[id].NextTurn = l.Player.NextTurn
	}
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
func (l *Level) AddMob(x, y, eType int) {

	//generate new unique id. loop checks for previous entity with that id.
	id := rand.Int()
	for _, ok := l.MobList[id]; ok; {
		id = rand.Int()
		_, ok = l.MobList[id]
	}

	e := NewEntity(x, y, id, eType)
	l.MobList[id] = e
	l.LevelMap.AddEntity(x, y, e)
}

type coord struct {
	x, y int
}

//finds the closest empty, visible spot within 5 squares and drops the item.
//returns false if there is no empty space
func (l *Level) DropItem(x, y int, i *Item) bool {

	spaces := make([]coord, 0, 68) //dropradius is 5, 68 possible locations
	l.LevelMap.ShadowCast(x, y, 5, GetEmptySpacesCast(&spaces))

	if len(spaces) == 0 {
		return false
	} else {

		d := 25 //util.distance returns d^2, so 25 is max
		best := 0
		//find closest open space
		for i, c := range spaces {
			if util.Distance(c.x, x, c.y, y) < d {
				d = util.Distance(c.x, x, c.y, y)
				best = i
			}
		}

		l.LevelMap.AddItem(spaces[best].x, spaces[best].y, i)
		return true
	}
}
