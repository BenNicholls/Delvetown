package delvemode

import "github.com/bennicholls/delvetown/console"
import "github.com/bennicholls/delvetown/data"

func (dm *DelveMode) Execute(command string) {
	switch command {
	case "regen":
		dm.level.GenerateCave()
		dm.gamelog.AddMessage("Level Regenerated!")
	case "fps":
		console.ToggleFPS()
	case "items":
		dm.player.AddItem(data.NewItem(data.ITEM_SWORD))
		dm.player.AddItem(data.NewItem(data.ITEM_AXE))
		dm.player.AddItem(data.NewItem(data.ITEM_HEALTH))
		dm.player.AddItem(data.NewItem(data.ITEM_HEALTH))
		dm.BuildHUDInventory()
	}
}
