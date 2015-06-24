package delvemode

func (dm *DelveMode) Execute(command string) {
	switch command {
	case "regen":
		dm.level.GenerateCave()
		dm.gamelog.AddMessage("Level Regenerated!")
	}
}
