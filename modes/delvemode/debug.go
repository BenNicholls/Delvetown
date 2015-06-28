package delvemode

import "github.com/bennicholls/delvetown/data"

func (dm *DelveMode) Execute(command string) {
	switch command {
	case "regen":
		dm.level.GenerateCave()
		dm.gamelog.AddMessage("Level Regenerated!")

	case "nolights":
		dm.level.LevelMap.ShadowCast(dm.player.X, dm.player.Y, dm.player.LightStrength, data.Darken)

		for i, _ := range dm.level.MobList {
			e := dm.level.MobList[i]
			dm.level.LevelMap.ShadowCast(e.X, e.Y, e.LightStrength, data.Darken)
		}
	}
}
