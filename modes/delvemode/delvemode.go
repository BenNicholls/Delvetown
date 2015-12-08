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
	HUDname   *ui.Textbox
	HUDhp     *ui.Textbox
	HUDattack *ui.Textbox

	HUDinventory *ui.List

	debug *ui.Inputbox

	activeElem ui.UIElem

	level   *data.Level
	player  *data.Entity
	gamelog Log

	xCamera, yCamera int

	tick          int
	memBrightness int //brightness to show memory tiles
}

func New() *DelveMode {
	dm := new(DelveMode)
	dm.tick = 1

	//Level Up!
	dm.level = data.NewLevel(100, 100)
	dm.level.GenerateCave()
	dm.player = dm.level.Player
	dm.player.Name = "The Ben"

	//UI stuff
	dm.view = ui.NewTileView(78, 38, 0, 0, 0, false)

	dm.logbox = ui.NewTextbox(76, 8, 1, 45, 1, true, false, "The Cave Feels Like It's Full of Mansters!!")
	dm.logbox.SetTitle("LOG")

	dm.gamelog = NewLog(dm.logbox)

	dm.sidebar = ui.NewContainer(16, 52, 79, 1, 0, true)
	dm.HUDname = ui.NewTextbox(16, 1, 0, 0, 0, false, true, dm.player.Name)
	dm.HUDhp = ui.NewTextbox(16, 1, 0, 2, 0, false, false, "HP: "+strconv.Itoa(dm.player.Health))
	dm.HUDattack = ui.NewTextbox(16, 1, 0, 3, 0, false, false, "Attack: "+strconv.Itoa(dm.player.BaseAttack))

	dm.HUDinventory = ui.NewList(16, 15, 79, 38, 0, true)
	dm.HUDinventory.SetTitle("Inventory")
	dm.UpdateHUDInventory()

	dm.sidebar.Add(dm.HUDname)
	dm.sidebar.Add(dm.HUDhp)
	dm.sidebar.Add(dm.HUDattack)

	dm.debug = ui.NewInputbox(76, 1, 1, 1, 2, true)
	dm.debug.SetTitle("Debugger")
	dm.debug.ToggleVisible()
	dm.activeElem = nil

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
		case sdl.K_KP_ENTER:
			if len(dm.player.Inventory) == 0 {
				dm.gamelog.AddMessage("No item to use!")
			} else {
				dm.player.ActionQueue <- dm.UseItem(dm.player, dm.HUDinventory.GetSelection())
			}
		case sdl.K_KP_PLUS:
			dm.HUDinventory.Next()
		case sdl.K_KP_MINUS:
			dm.HUDinventory.Prev()
		case sdl.K_ESCAPE:
			dm.activeElem = dm.debug
			dm.debug.ToggleVisible()
		}
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
				action := dm.HuntBehaviour(e) //TODO: replace with context-specific AI
				action(dm.level.MobList[k])
			}
		}
	}

	//update UI elements
	dm.HUDhp.ChangeText("HP: " + strconv.Itoa(dm.player.Health))
	dm.tick++

	//check for gamestate changes
	if dm.player.Health <= 0 {
		return modes.NewGameOver()
	} else if len(dm.level.MobList) == 0 {
		return modes.NewWinner()
	}

	return nil
}

func (dm *DelveMode) UpdateHUDInventory() {
	dm.HUDinventory.ClearElements()

	if len(dm.player.Inventory) == 0 {
		invElem := ui.NewTextbox(16, 1, 0, 6, 0, false, true, "No Items")
		dm.HUDinventory.Add(invElem)
		dm.HUDinventory.Highlight = false
	} else {
		for i, item := range dm.player.Inventory {
			invElem := ui.NewTextbox(16, 1, 0, i, 0, false, false, item.Name)
			dm.HUDinventory.Add(invElem)
			dm.HUDinventory.Highlight = true
		}
	}
	dm.HUDinventory.CheckSelection()

}

func (dm *DelveMode) Render() {

	//focus camera on player
	w, h := dm.view.Width, dm.view.Height
	dm.xCamera, dm.yCamera = dm.player.X-w/2, dm.player.Y-h/2

	dm.view.Clear()

	//update player memory
	dm.level.LevelMap.ShadowCast(dm.player.X, dm.player.Y, dm.player.SightRange, dm.MemoryCast())

	//Draw the world.
	for i := 0; i < w*h; i++ {

		//Map coordinates
		x, y := i%w+dm.xCamera, i/w+dm.yCamera

		//check if tile has ever been seen.
		if dm.level.MemoryMap.LastVisible(x, y) != 0 {

			//try to see if an entity is occupying the space. if so, draw it. otherwise, draw the tile.
			e := dm.level.MemoryMap.GetEntity(x, y)
			item := dm.level.MemoryMap.GetItem(x, y)
			if e != nil {
				dm.view.DrawEntity(i%w, i/w, e.GetVisuals())
			} else if item != nil {
				dm.view.DrawEntity(i%w, i/w, item.GetVisuals())
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
	dm.HUDinventory.Render()
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
