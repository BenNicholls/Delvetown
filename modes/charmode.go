package modes

import "github.com/veandco/go-sdl2/sdl"
import "github.com/bennicholls/delvetown/ui"
import "github.com/bennicholls/delvetown/util"

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

	activeElem ui.UIElem
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
	cm.flavourtext = ui.NewTextbox(16, 8, 0, 0, 0, false, false, "The fightman is a muscley man who goes from town to town picking fights. He loves to battle, it gives him a big boner.")
	cm.mainstat = ui.NewTextbox(16, 1, 0, 9, 0, false, false, "MAIN STAT: Body")
	cm.description.Add(cm.flavourtext, cm.mainstat)

	cm.stats = ui.NewContainer(22, 16, 25, 26, 0, true)
	cm.stats.SetTitle("Stats")
	cm.hp = ui.NewTextbox(22, 1, 0, 1, 0, false, false, "HP: ")
	cm.att = ui.NewTextbox(22, 1, 0, 2, 0, false, false, "ATT: ")
	cm.weapon = ui.NewTextbox(22, 1, 0, 4, 0, false, false, "WEAPON: ")
	cm.armour = ui.NewTextbox(22, 1, 0, 5, 0, false, false, "ARMOUR: ")
	cm.mind = ui.NewTextbox(22, 1, 0, 8, 0, false, false, "MIND: ")
	cm.body = ui.NewTextbox(22, 1, 0, 9, 0, false, false, "BODY: ")
	cm.spirit = ui.NewTextbox(22, 1, 0, 10, 0, false, false, "SPIRIT: ")
	cm.stats.Add(cm.hp, cm.att, cm.weapon, cm.armour, cm.mind, cm.body, cm.spirit)

	cm.screen.Add(cm.description, cm.stats, cm.name, cm.class)

	cm.activeElem = cm.name

	return cm
}

func (cm *CharMode) Update() (error, GameModer) {

	return nil, nil
}

func (cm *CharMode) Render() {
	cm.screen.Render()
}

func (cm *CharMode) HandleKeypress(key sdl.Keycode) error {

	if key == sdl.K_TAB {
		cm.CycleUI()
		return nil
	}

	if cm.activeElem == cm.name {
		if util.ValidText(rune(key)) {
			cm.name.InsertText(rune(key))
		} else {
			switch key {
				case sdl.K_BACKSPACE:
					cm.name.Delete()
				case sdl.K_SPACE:
					cm.name.Insert(" ")
			}
		}
	} else if cm.activeElem == cm.class {
		switch key {
		case sdl.K_DOWN, sdl.K_KP_2:
			cm.class.Next()
		case sdl.K_UP, sdl.K_KP_8:
			cm.class.Prev()
		}
	}

	return nil
}

func (cm *CharMode) CycleUI() {
	switch cm.activeElem {
	case cm.name:
		cm.activeElem = cm.class
		cm.name.ToggleCursor()
	case cm.class:
		cm.activeElem = cm.name
		cm.name.ToggleCursor()
	}
}