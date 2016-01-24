package delvemode

import "github.com/bennicholls/delvetown/console"

func (dm *DelveMode) Execute(command string) {
	switch command {
	case "regen":
		dm.level.GenerateCave()
		dm.gamelog.AddMessage("Level Regenerated!")
	case "fps":
		console.ShowFPS = !console.ShowFPS
	}
}
