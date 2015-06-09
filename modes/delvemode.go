package modes

import "github.com/bennicholls/delvetown/ui"
import "github.com/bennicholls/delvetown/data"
import "github.com/bennicholls/delvetown/entities"
import "github.com/veandco/go-sdl2/sdl"
import "strconv"
import "math/rand"

type DelveMode struct {

	//UI stuff!
	logbox  *ui.Textbox   //combat and action log
	view    *ui.TileView  //main map area
	sidebar *ui.Container //holds the various ui elements for the player's info display.

	//sidebar ui elements
	hp          *ui.Textbox
	stepCounter *ui.Textbox

	level *data.Level

	xCamera, yCamera int

	player *entities.Entity
	log    string

	tick int
	step int

	//player offset
	pDX, pDY int
}

func NewDelveMode() *DelveMode {
	dm := new(DelveMode)

	//UI stuff
	dm.logbox = ui.NewTextbox(98, 8, 1, 1, 0, true, "TEST TEXT whatever blah blah blahxxx")
	dm.view = ui.NewTileView(83, 39, 1, 10, 0, true)
	dm.sidebar = ui.NewContainer(14, 39, 85, 10, 0, true)
	dm.hp = ui.NewTextbox(10, 1, 0, 0, 0, false, "hello")
	dm.stepCounter = ui.NewTextbox(10, 1, 0, 1, 0, false, "")
	
	dm.sidebar.Add(dm.hp)
	dm.sidebar.Add(dm.stepCounter)

	//Level Up!
	dm.level = data.NewLevel(100, 100)
	dm.level.GenerateRandom()
	dm.player = dm.level.Player

	dm.pDX, dm.pDY = 0, 0
	dm.tick, dm.step = 0, 0

	return dm
}

func (dm *DelveMode) HandleKeypress(key sdl.Keycode) {
	switch key {
	case sdl.K_DOWN:
		dm.pDY = 1
	case sdl.K_UP:
		dm.pDY = -1
	case sdl.K_LEFT:
		dm.pDX = -1
	case sdl.K_RIGHT:
		dm.pDX = 1
	}
}


func (dm *DelveMode) Update() GameModer {

	//player movement
	if dm.pDX != 0 || dm.pDY != 0 {
		
		//check if this is an attack, if so, attack!
		e := dm.level.GetEntity(dm.player.X + dm.pDX, dm.player.Y + dm.pDY)
		if e != nil {
			e.Health -= 5
			dm.player.Health -= 1
			dm.logbox.ChangeText("You attack! TO VALHALLA!!!")
			if e.Health <= 0 {
				dm.level.RemoveEntity(e.ID)
				dm.logbox.ChangeText("You are a MURDERER!")
			}
		} else {
			dm.level.MovePlayer(dm.pDX, dm.pDY)
			dm.step += 1
		}
		
		dm.pDX, dm.pDY = 0, 0
		
		//enemy movement
		for ID, mob := range dm.level.MobList {
			eDX, eDY := rand.Intn(3) - 1, rand.Intn(3) - 1
		
			//check if attacking the player
			if mob.X + eDX == dm.player.X && mob.Y + eDY == dm.player.Y {
				dm.player.Health -= 5
				dm.level.MobList[ID].Health -= 1
				if mob.Health <= 0 {
					dm.level.RemoveEntity(ID)
					dm.logbox.ChangeText("It HIT YOU. OUCH!")
				}
			} else {
				e = dm.level.GetEntity(mob.X + eDX, mob.Y + eDY)
				if e == nil {
					dm.level.MoveMob(ID, eDX, eDY)
				}
			}
		}	
	}
	
	
	//update UI elements
	dm.hp.ChangeText("HP: " + strconv.Itoa(dm.player.Health))
	dm.stepCounter.ChangeText("Steps: " + strconv.Itoa(dm.step))
	dm.tick++

	//check for gamestate changes
	if dm.player.Health <= 0 {
		return NewGameOver()
	}

	return nil
}

func (dm *DelveMode) Render() {

	w, h := dm.view.Width, dm.view.Height
	dm.xCamera, dm.yCamera = dm.player.X-dm.view.Width/2, dm.player.Y-dm.view.Height/2

	//Draw the world.
	var e *entities.Entity
	for i := 0; i < w*h; i++ {
		x, y := i%w+dm.xCamera, i/w+dm.yCamera

		//try to see if an entity is occupying the space. if so, draw it. otherwise, draw the tile.
		e = dm.level.GetEntity(x, y)
		if e != nil {
			dm.view.DrawEntity(i%w, i/w, e.Glyph, e.Fore)
		} else {
			t := dm.level.Levelmap.GetTileType(x, y)
			dm.view.DrawTile(i%w, i/w, t)
		}
	}

	//render ui elements
	dm.logbox.Render()
	dm.view.Render()
	dm.sidebar.Render()
}
