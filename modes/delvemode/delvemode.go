package delvemode

import "github.com/bennicholls/delvetown/ui"
import "github.com/bennicholls/delvetown/data"
import "github.com/bennicholls/delvetown/modes"
import "github.com/bennicholls/delvetown/util"
import "github.com/veandco/go-sdl2/sdl"
import "strconv"

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

	player *data.Entity

	tick int

	memBrightness int //brightness to show memory tiles
}

func New() *DelveMode {
	dm := new(DelveMode)

	//UI stuff
	dm.view = ui.NewTileView(83, 38, 1, 1, 0, true)

	dm.logbox = ui.NewTextbox(83, 8, 1, 41, 1, true, "TEST TEXT whatever blah blah blahxxx")
	dm.logbox.SetTitle("LOG")

	dm.gamelog = NewLog(dm.logbox)

	dm.sidebar = ui.NewContainer(13, 48, 86, 1, 0, true)
	dm.sidebar.SetTitle("GOOD STATS")
	dm.hp = ui.NewTextbox(10, 1, 0, 0, 0, false, "hello")
	dm.stepCounter = ui.NewTextbox(10, 1, 0, 1, 0, false, "")

	dm.sidebar.Add(dm.hp)
	dm.sidebar.Add(dm.stepCounter)

	dm.debug = ui.NewInputbox(83, 1, 1, 48, 2, true)
	dm.debug.SetTitle("Debugger")
	dm.debug.ToggleVisible()
	dm.activeElem = nil

	//Level Up!
	dm.level = data.NewLevel(100, 100)
	dm.level.GenerateCave()
	dm.player = dm.level.Player

	dm.tick = 1

	dm.memBrightness = 80

	return dm
}

func (dm *DelveMode) HandleKeypress(key sdl.Keycode) {

	if dm.activeElem == dm.debug {
		if util.ValidText(rune(key)) {
			dm.debug.InsertText(rune(key))
		} else {
			switch key {
			case sdl.K_BACKSPACE:
				dm.debug.Delete()
			case sdl.K_SPACE:
				dm.debug.Insert(" ")
			case sdl.K_ESCAPE:
				dm.activeElem = nil
				dm.debug.ToggleVisible()
			case sdl.K_RETURN:
				dm.Execute(dm.debug.GetText())
				dm.debug.Reset()
				dm.activeElem = nil
				dm.debug.ToggleVisible()
			}
		}
	} else {

		switch key {
		case sdl.K_DOWN, sdl.K_KP_2:
			dm.player.ActionQueue <- dm.AttackMove(dm.player, 0, 1)
		case sdl.K_UP, sdl.K_KP_8:
			dm.player.ActionQueue <- dm.AttackMove(dm.player, 0, -1)
		case sdl.K_LEFT, sdl.K_KP_4:
			dm.player.ActionQueue <- dm.AttackMove(dm.player, -1, 0)
		case sdl.K_RIGHT, sdl.K_KP_6:
			dm.player.ActionQueue <- dm.AttackMove(dm.player, 1, 0)
		case sdl.K_KP_7:
			dm.player.ActionQueue <- dm.AttackMove(dm.player, -1, -1)
		case sdl.K_KP_9:
			dm.player.ActionQueue <- dm.AttackMove(dm.player, 1, -1)
		case sdl.K_KP_1:
			dm.player.ActionQueue <- dm.AttackMove(dm.player, -1, 1)
		case sdl.K_KP_3:
			dm.player.ActionQueue <- dm.AttackMove(dm.player, 1, 1)
		case sdl.K_SPACE:
			dm.player.NextTurn += 1
		case sdl.K_ESCAPE:
			dm.activeElem = dm.debug
			dm.debug.ToggleVisible()
		}
	}
}

//Decide if a move is an atack or not
func (dm *DelveMode) AttackMove(e *data.Entity, dx, dy int) data.Action {
	t := dm.level.LevelMap.GetEntity(e.X+dx, e.Y+dy)

	if t != nil {
		return dm.AttackAction(dx, dy, 5)
	} else {
		return dm.MoveAction(dx, dy, 3)
	}
}

func (dm *DelveMode) Update() modes.GameModer {

	if dm.player.NextTurn <= dm.tick {
		if len(dm.player.ActionQueue) == 0 {
			//x, y := util.GenerateDirection()
			//dm.player.ActionQueue <- dm.MoveAction(x, y, 5)
			return nil // block update, waiting on player move
		} else {
			action := <-dm.player.ActionQueue
			action(dm.player)
		}
	}

	for k, e := range dm.level.MobList {
		if e.NextTurn <= dm.tick {
			if len(e.ActionQueue) > 0 {
				action := <-e.ActionQueue
				action(dm.level.MobList[k])
			} else {
				dx, dy := util.GenerateDirection()
				action := dm.MoveAction(dx, dy, 10)
				action(dm.level.MobList[k])
			}
		}
	}

	//update UI elements
	dm.hp.ChangeText("HP: " + strconv.Itoa(dm.player.Health))
	dm.stepCounter.ChangeText("Ticks: " + strconv.Itoa(dm.tick))
	dm.tick++

	//check for gamestate changes
	if dm.player.Health <= 0 {
		return modes.NewGameOver()
	} else if len(dm.level.MobList) == 0 {
		return modes.NewWinner()
	}

	return nil
}

func (dm *DelveMode) Render() {

	//focus camera on player
	w, h := dm.view.Width, dm.view.Height
	dm.xCamera, dm.yCamera = dm.player.X-w/2, dm.player.Y-h/2

	dm.view.Clear()

	var e *data.Entity

	//update player memory
	dm.level.LevelMap.ShadowCast(dm.player.X, dm.player.Y, 50, dm.MemoryCast())

	//Draw the world.
	for i := 0; i < w*h; i++ {

		//Map coordinates
		x, y := i%w+dm.xCamera, i/w+dm.yCamera

		//check if tile has ever been seen.
		if dm.level.MemoryMap.LastVisible(x, y) != 0 {

			//try to see if an entity is occupying the space. if so, draw it. otherwise, draw the tile.
			e = dm.level.MemoryMap.GetEntity(x, y)
			if e != nil {
				dm.view.DrawEntity(i%w, i/w, e.Glyph, e.Fore)
			} else {
				t := dm.level.MemoryMap.GetTile(x, y)
				dm.view.DrawTile(i%w, i/w, t)
			}

			if dm.level.MemoryMap.LastVisible(x, y) != dm.tick {
				dm.view.ApplyLight(i%w, i/w, dm.memBrightness)
			} else {
				dm.view.ApplyLight(i%w, i/w, dm.level.LevelMap.GetTile(x, y).Light.Bright)
			}

		}
	}

	//render ui elements
	dm.logbox.Render()
	dm.view.Render()
	dm.sidebar.Render()
	dm.debug.Render()
}

//Baby's first closure. Holy crap I am so proud of this.
//Pass this into the shadowcaster to copy the visible portion
//of the map into the player's memory.
func (dm *DelveMode) MemoryCast() data.Cast {
	return func(m *data.TileMap, x, y, d, r int) {
		if m.GetTile(x, y).Light.Bright > 0 {
			m.SetVisible(x, y, dm.tick)
			dm.level.MemoryMap.SetTile(x, y, m.GetTile(x, y))
		}
	}
}
