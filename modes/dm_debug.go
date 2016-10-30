package modes

import "github.com/bennicholls/delvetown/console"
import "github.com/bennicholls/delvetown/data"
import "github.com/veandco/go-sdl2/sdl"
import "github.com/bennicholls/delvetown/util"

func (dm *DelveMode) Execute(command string) {
	switch command {
	case "regen":
		dm.level.GenerateCave()
		dm.UpdatePlayerVision()
		dm.gamelog.AddMessage("Level Regenerated!")
	case "fps":
		console.ToggleFPS()
	case "items":
		dm.player.AddItem(data.NewItem(data.ITEM_SWORD))
		dm.player.AddItem(data.NewItem(data.ITEM_AXE))
		dm.player.AddItem(data.NewItem(data.ITEM_HEALTH))
		dm.player.AddItem(data.NewItem(data.ITEM_HEALTH))
		dm.BuildHUDInventory()
	case "hp":
		dm.player.Health += 100
		dm.UpdateUI() 
	}
}

func (dm *DelveMode) DebugKeypress(key sdl.Keycode) {
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
				dm.debug.ToggleFocus()
			case sdl.K_RETURN:
				dm.Execute(dm.debug.GetText())
				dm.debug.Reset()
				dm.activeElem = nil
				dm.debug.ToggleVisible()
				dm.debug.ToggleFocus()
			}
		}
}