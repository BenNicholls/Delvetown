package modes

import "github.com/bennicholls/delvetown/ui"
import "github.com/bennicholls/delvetown/data"
import "github.com/veandco/go-sdl2/sdl"
import "strconv"
import "errors"

type DelveMode struct {

	//UI stuff!
	logbox  *ui.Textbox   //combat and action log
	view    *ui.TileView  //main map area
	sidebar *ui.Container //holds the various ui elements for the player's info display.

	//sidebar ui elements
	HUDname      *ui.Textbox
	HUDhp        *ui.Textbox
	HUDattack    *ui.Textbox
	HUDtestvalue *ui.Textbox

	HUDinventory *ui.List

	HUDenemylist *ui.List

	HUDweapon *ui.Textbox
	HUDarmour *ui.Textbox

	debug *ui.Inputbox

	activeElem ui.UIElem

	dungeon *data.Dungeon
	level   *data.Level
	player  *data.Entity
	gamelog Log

	xCamera, yCamera int

	tick          int
	memBrightness int //brightness to show memory tiles
	frame         int //framecounter, for testing

}

func NewDelvemode(p *data.Entity, d *data.Dungeon) *DelveMode {
	dm := new(DelveMode)
	dm.tick = 1

	//Level Up!
	dm.dungeon = d
	dm.level = d.Generators[0](p, 100, 100)
	dm.player = p

	//UI stuff
	dm.view = ui.NewTileView(78, 38, 0, 0, 0, false)

	dm.logbox = ui.NewTextbox(76, 8, 1, 45, 1, true, false, "")
	dm.logbox.SetTitle("LOG")

	dm.gamelog = NewLog(dm.logbox)
	dm.gamelog.AddMessage(dm.dungeon.Intro)

	dm.sidebar = ui.NewContainer(16, 52, 79, 1, 0, true)
	dm.HUDname = ui.NewTextbox(16, 1, 0, 0, 0, false, true, dm.player.Name)
	dm.HUDhp = ui.NewTextbox(16, 1, 0, 2, 0, false, false, "HP: "+strconv.Itoa(dm.player.CalcHP())+"/"+strconv.Itoa(dm.player.MaxStats.HP))
	dm.HUDattack = ui.NewTextbox(16, 1, 0, 3, 0, false, false, "Attack: "+strconv.Itoa(dm.player.CalcAttack()))
	dm.HUDweapon = ui.NewTextbox(16, 1, 0, 5, 0, false, false, "W: "+dm.player.GetEquipmentName(data.SLOT_WEAPON))
	dm.HUDarmour = ui.NewTextbox(16, 1, 0, 6, 0, false, false, "A: "+dm.player.GetEquipmentName(data.SLOT_ARMOUR))
	dm.HUDtestvalue = ui.NewTextbox(16, 1, 0, 8, 0, false, false, "MonNum :"+strconv.Itoa(len(dm.level.MobList)))

	dm.sidebar.Add(dm.HUDname, dm.HUDhp, dm.HUDattack, dm.HUDweapon, dm.HUDarmour, dm.HUDtestvalue)

	dm.HUDinventory = ui.NewList(16, 15, 79, 38, 0, true, "No Items")
	dm.HUDinventory.SetTitle("Inventory")

	dm.HUDenemylist = ui.NewList(16, 15, 79, 21, 0, true, "No Enemies")
	dm.HUDenemylist.SetTitle("Enemies")
	dm.HUDenemylist.Highlight = false

	dm.debug = ui.NewInputbox(76, 1, 1, 1, 2, true)
	dm.debug.SetTitle("Debugger")
	dm.debug.ToggleVisible()
	dm.activeElem = nil

	dm.memBrightness = 80
	dm.UpdatePlayerVision()

	dm.frame = 0

	return dm
}

func (dm *DelveMode) BuildHUDInventory() {
	dm.HUDinventory.ClearElements()

	for _, item := range dm.player.Inventory {
		dm.HUDinventory.Append(item.Name)
	}
}

func (dm *DelveMode) BuildHUDenemylist() {
	dm.HUDenemylist.ClearElements()
	w, _ := dm.HUDenemylist.GetDims()

	for i, enemy := range dm.player.VisibleEntities {
		dm.HUDenemylist.Add(ui.NewTextbox(w, 1, 0, i, 0, false, false, enemy.Name))
	}
}

func (dm *DelveMode) HandleKeypress(key sdl.Keycode) {

	if dm.activeElem == dm.debug {
		dm.DebugKeypress(key)
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
				dm.player.ActionQueue <- dm.UseInventory(dm.player, dm.HUDinventory.GetSelection())
			}
		case sdl.K_KP_PLUS:
			dm.HUDinventory.Next()
		case sdl.K_KP_MINUS:
			dm.HUDinventory.Prev()
		case sdl.K_ESCAPE:
			dm.activeElem = dm.debug
			dm.debug.ToggleVisible()
			dm.debug.ToggleFocus()
		case sdl.K_d:
			if len(dm.player.Inventory) == 0 {
				dm.gamelog.AddMessage("No item to drop!")
			} else {
				dm.player.ActionQueue <- dm.DropInventoryItem(dm.player, dm.HUDinventory.GetSelection())
			}
		}
	}
}

func (dm *DelveMode) Update() (error, GameModer) {

	if dm.player.NextTurn <= dm.tick {
		if len(dm.player.ActionQueue) == 0 {
			return nil, nil // block update, waiting on player move
		} else {
			action := <-dm.player.ActionQueue
			action(dm.player)
		}
	}

	for _, e := range dm.level.MobList {
		if e.NextTurn <= dm.tick {
			if len(e.ActionQueue) > 0 {
				action := <-e.ActionQueue
				action(e)
			} else {
				action := dm.HuntBehaviour(e) //TODO: replace with context-specific AI
				action(e)
			}
		}
	}

	dm.tick++

	//update player memory (has to be after tick++, uses tick to mark when space was last seen)
	dm.UpdatePlayerVision()

	//update UI elements
	dm.UpdateUI()

	//check for gamestate changes
	if dm.player.CalcHP() <= 0 {
		return errors.New("Mode Change"), NewGameOver()
	} else if len(dm.level.MobList) == 0 {
		return errors.New("Mode Change"), NewWinner()
	}

	return nil, nil
}

func (dm *DelveMode) UpdateUI() {
	dm.HUDhp.ChangeText("HP: " + strconv.Itoa(dm.player.CalcHP()) + "/" + strconv.Itoa(dm.player.MaxStats.HP))
	dm.HUDattack.ChangeText("Attack: " + strconv.Itoa(dm.player.CalcAttack()))
	dm.HUDweapon.ChangeText("W: " + dm.player.GetEquipmentName(data.SLOT_WEAPON))
	dm.HUDarmour.ChangeText("A: " + dm.player.GetEquipmentName(data.SLOT_ARMOUR))
	dm.HUDtestvalue.ChangeText("MonNum :" + strconv.Itoa(len(dm.level.MobList)))
}

func (dm *DelveMode) Render() {

	//focus camera on player
	w, h := dm.view.Width, dm.view.Height
	dm.xCamera, dm.yCamera = dm.player.X-w/2, dm.player.Y-h/2

	dm.view.Clear()

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
				if dm.player.CanSee(e.ID) && !(e.X == x && e.Y == y) {
					//dm.level.MemoryMap.RemoveEntity(x, y)
				} else {
					dm.view.DrawVisuals(i%w, i/w, e.GetVisuals())
				}
			} else if item != nil {
				dm.view.DrawVisuals(i%w, i/w, item.GetVisuals())
			} else {
				t := dm.level.MemoryMap.GetTile(x, y)
				dm.view.DrawVisuals(i%w, i/w, t.GetVisuals())
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
	dm.HUDenemylist.Render()

	if dm.activeElem != nil {
		dm.activeElem.Render()
	}

	//flickering torch test... not ready for primetime.
	// if dm.frame % 5 == 0 {
	// 	dm.level.LevelMap.ShadowCast(dm.player.X, dm.player.Y, dm.player.LightStrength, data.Darken)
	// 	dm.level.LevelMap.ShadowCast(dm.player.X, dm.player.Y, dm.player.LightStrength + rand.Intn(3) -1, data.Lighten)
	// 	dm.UpdatePlayerVision()
	// }

	dm.frame++
}

func (dm *DelveMode) UpdatePlayerVision() {
	dm.level.LevelMap.ShadowCast(dm.player.X, dm.player.Y, dm.player.MaxStats.SightRange, dm.MemoryCast())
	dm.BuildHUDenemylist()
}

//Baby's first closure. Holy crap I am so proud of this.
//Pass this into the shadowcaster to copy the visible portion
//of the map into the player's memory.
func (dm *DelveMode) MemoryCast() data.Cast {
	dm.player.ClearVisible()
	return func(m *data.TileMap, x, y, d, r int) {
		if m.GetTile(x, y).Light.Bright > 0 {
			m.SetVisible(x, y, dm.tick)
			dm.level.MemoryMap.SetTile(x, y, m.GetTile(x, y))
			if m.GetTile(x, y).Entity != nil {
				dm.player.AddVisibleEntity(m.GetTile(x, y).Entity)
			}
		}
	}
}
