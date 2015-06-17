package delvemode

import "github.com/bennicholls/delvetown/ui"
import "github.com/bennicholls/delvetown/data"
import "github.com/bennicholls/delvetown/entities"
import "github.com/bennicholls/delvetown/modes"
import "github.com/veandco/go-sdl2/sdl"
import "strconv"
import "math/rand"
import "math"

type DelveMode struct {

	//UI stuff!
	logbox  *ui.Textbox   //combat and action log
	view    *ui.TileView  //main map area
	sidebar *ui.Container //holds the various ui elements for the player's info display.

	//sidebar ui elements
	hp          *ui.Textbox
	stepCounter *ui.Textbox

	debug *ui.Inputbox

	activeElem ui.UIElem

	level   *data.Level
	gamelog Log

	xCamera, yCamera int

	player *entities.Entity

	tick, step int

	//player offset
	pDX, pDY int
}

func New() *DelveMode {
	dm := new(DelveMode)

	//UI stuff
	dm.view = ui.NewTileView(83, 39, 1, 1, 0, true)

	dm.logbox = ui.NewTextbox(83, 8, 1, 41, 1, true, "TEST TEXT whatever blah blah blahxxx")
	dm.logbox.SetTitle("LOG")

	dm.gamelog = NewLog(dm.logbox)

	dm.sidebar = ui.NewContainer(14, 48, 85, 1, 0, true)
	dm.sidebar.SetTitle("GOOD STATS")
	dm.hp = ui.NewTextbox(10, 1, 0, 0, 0, false, "hello")
	dm.stepCounter = ui.NewTextbox(10, 1, 0, 1, 0, false, "")

	dm.sidebar.Add(dm.hp)
	dm.sidebar.Add(dm.stepCounter)

	dm.debug = ui.NewInputbox(20, 1, 10, 20, 2, true)
	dm.debug.SetTitle("Debugger")
	dm.debug.ToggleVisible()
	dm.activeElem = nil

	//Level Up!
	dm.level = data.NewLevel(100, 100)
	dm.level.GenerateCave()
	dm.player = dm.level.Player

	dm.pDX, dm.pDY = 0, 0
	dm.tick, dm.step = 0, 0

	return dm
}

func (dm *DelveMode) HandleKeypress(key sdl.Keycode) {

	if dm.activeElem == dm.debug {
		switch key {
		case sdl.K_a:
			dm.debug.Insert("a")
		case sdl.K_BACKSPACE:
			dm.debug.Delete()
		case sdl.K_ESCAPE:
			dm.activeElem = nil
			dm.debug.ToggleVisible()
		}
	} else {

		switch key {
		case sdl.K_DOWN:
			dm.pDY = 1
		case sdl.K_UP:
			dm.pDY = -1
		case sdl.K_LEFT:
			dm.pDX = -1
		case sdl.K_RIGHT:
			dm.pDX = 1
		case sdl.K_SPACE:
			dm.level.GenerateCave()
		case sdl.K_ESCAPE:
			dm.activeElem = dm.debug
			dm.debug.ToggleVisible()
		}
	}
}

func (dm *DelveMode) Update() modes.GameModer {

	//Brownian motion for player, for testing engine speed.
	//dm.pDX, dm.pDY = rand.Intn(3)-1, rand.Intn(3)-1

	//player movement
	if dm.pDX != 0 || dm.pDY != 0 {

		//check if this is an attack, if so, attack!
		e := dm.level.LevelMap.GetEntity(dm.player.X+dm.pDX, dm.player.Y+dm.pDY)
		if e != nil {
			e.Health -= 5
			dm.player.Health -= 1
			dm.gamelog.AddMessage("You attack! TO VALHALLA!!!")
			if e.Health <= 0 {
				dm.level.RemoveEntity(e.ID)
				dm.gamelog.AddMessage("You are a MURDERER!")
			}
		} else {
			dm.level.MovePlayer(dm.pDX, dm.pDY)
			dm.step += 1
		}

		dm.pDX, dm.pDY = 0, 0

		//enemy movement
		for ID, mob := range dm.level.MobList {
			eDX, eDY := rand.Intn(3)-1, rand.Intn(3)-1

			//check if attacking the player
			if mob.X+eDX == dm.player.X && mob.Y+eDY == dm.player.Y {
				dm.player.Health -= 5
				dm.level.MobList[ID].Health -= 1
				dm.gamelog.AddMessage("It HIT YOU. OUCH!")
				if mob.Health <= 0 {
					dm.level.RemoveEntity(ID)
				}
			} else {
				e = dm.level.LevelMap.GetEntity(mob.X+eDX, mob.Y+eDY)
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
		return modes.NewGameOver()
	}

	return nil
}

func (dm *DelveMode) Render() {

	//focus camera on player
	w, h := dm.view.Width, dm.view.Height
	dm.xCamera, dm.yCamera = dm.player.X-w/2, dm.player.Y-h/2

	dm.view.Clear()

	var e *entities.Entity

	//naive raycasting FOV. NOTE: Replace this with something fancy
	for ray := 0.; ray < 2*math.Pi; ray += 0.025 {
		for rx, ry, d := 0, 0, 0.; d < 40; d += 1 {
			rx, ry = dm.player.X+int(d*math.Cos(ray)), dm.player.Y+int(d*math.Sin(ray))
			dm.level.LevelMap.SetVisible(rx, ry, dm.tick)
			dm.level.MemoryMap.SetTile(rx, ry, dm.level.LevelMap.GetTile(rx, ry))
			if !data.IsPassable(dm.level.LevelMap.GetTileType(rx, ry)) {
				break
			}
		}
	}

	//Draw the world.
	for i := 0; i < w*h; i++ {

		//Map coordinates
		x, y := i%w+dm.xCamera, i/w+dm.yCamera
		//try to see if an entity is occupying the space. if so, draw it. otherwise, draw the tile.
		if dm.level.MemoryMap.LastVisible(x, y) != 0 {
			if dm.level.MemoryMap.LastVisible(x, y) != dm.tick {
				dm.level.MemoryMap.ChangeTileColour(x, y, 70)
			}
			e = dm.level.MemoryMap.GetEntity(x, y)
			if e != nil {
				dm.view.DrawEntity(i%w, i/w, e.Glyph, e.Fore)
			} else {
				t := dm.level.MemoryMap.GetTile(x, y)
				dm.view.DrawTile(i%w, i/w, t)
			}
		}
	}

	//render ui elements
	dm.logbox.Render()
	dm.view.Render()
	dm.sidebar.Render()
	dm.debug.Render()
}
