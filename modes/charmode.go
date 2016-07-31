package modes

import "github.com/veandco/go-sdl2/sdl"
import "github.com/bennicholls/delvetown/ui"

//TODO - this is a stupid name, change it
type CharMode struct {
	screen *ui.Container
	name   *ui.Inputbox
	class  *ui.List

	description *ui.Container
	flavourtext *ui.Textbox
	mainstat    *ui.Textbox

	stats  *ui.Container
	hp     *ui.Textbox
	att    *ui.Textbox
	weapon *ui.Textbox
	armour *ui.Textbox
	mind   *ui.Textbox
	body   *ui.Textbox
	spirit *ui.Textbox
}

func NewChar() *CharMode {

	cm := new(CharMode)

	cm.screen = ui.NewContainer(94, 52, 1, 1, 0, true)
	cm.screen.SetTitle("CHOOSE YOUR CHOICE, DELVEMAN.")

	cm.name = ui.NewInputbox(22, 1, 25, 16, 0, true)
	cm.name.SetTitle("Player Name")

	cm.class = ui.NewList(22, 3, 25, 20, 0, true, "")
	cm.class.SetTitle("Class")
	cm.class.Append("Fightman", "Book Nerd", "Bald Man")

	cm.description = ui.NewContainer(16, 26, 50, 16, 0, true)
	cm.flavourtext = ui.NewTextbox(16, 8, 0, 0, 0, false, false, "Fightman stuff")
	cm.mainstat = ui.NewTextbox(16, 1, 0, 9, 0, false, false, "MAIN STAT: Body")
	cm.description.Add(cm.flavourtext, cm.mainstat)

	cm.stats = ui.NewContainer(22, 16, 25, 26, 0, true)
	cm.hp = ui.NewTextbox(22, 1, 0, 1, 0, false, false, "HP: ")
	cm.att = ui.NewTextbox(22, 1, 0, 2, 0, false, false, "ATT: ")
	cm.weapon = ui.NewTextbox(22, 1, 0, 4, 0, false, false, "WEAPON: ")
	cm.armour = ui.NewTextbox(22, 1, 0, 5, 0, false, false, "ARMOUR: ")
	cm.mind = ui.NewTextbox(22, 1, 0, 8, 0, false, false, "MIND: ")
	cm.body = ui.NewTextbox(22, 1, 0, 9, 0, false, false, "BODY: ")
	cm.spirit = ui.NewTextbox(22, 1, 0, 10, 0, false, false, "SPIRIT: ")
	cm.stats.Add(cm.hp, cm.att, cm.weapon, cm.armour, cm.mind, cm.body, cm.spirit)

	cm.screen.Add(cm.description, cm.stats, cm.name, cm.class)

	return cm
}

func (mm *CharMode) Update() (error, GameModer) {

	return nil, nil
}

func (mm *CharMode) Render() {
	mm.screen.Render()
}

func (mm *CharMode) HandleKeypress(key sdl.Keycode) error {
	switch key {
	case sdl.K_DOWN, sdl.K_KP_2:
		mm.class.Next()
	case sdl.K_UP, sdl.K_KP_8:
		mm.class.Prev()
	}

	return nil
}
