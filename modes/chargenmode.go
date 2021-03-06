package modes

import "math/rand"
import "strconv"
import "errors"
import "github.com/veandco/go-sdl2/sdl"
import "github.com/bennicholls/delvetown/ui"
import "github.com/bennicholls/delvetown/util"
import "github.com/bennicholls/delvetown/data"

type CharGenMode struct {
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

	confirm *ui.Button

	activeElem ui.UIElem

	character *data.Entity
}

func NewCharGen() *CharGenMode {

	cm := new(CharGenMode)

	//BEGIN UI STUFF
	cm.screen = ui.NewContainer(94, 52, 1, 1, 0, true)
	cm.screen.SetTitle("CHOOSE YOUR CHOICE, DELVEMAN.")

	cm.name = ui.NewInputbox(22, 1, 25, 10, 0, true)
	cm.name.SetTitle("Player Name")

	cm.class = ui.NewList(22, 3, 25, 14, 0, true, "")
	cm.class.SetTitle("Class")
	cm.class.Append("Fightman", "Book Nerd", "Bald Man")

	cm.description = ui.NewContainer(18, 26, 50, 10, 0, true)
	cm.flavourtext = ui.NewTextbox(18, 10, 0, 0, 0, false, false, "FLAVOUR")
	cm.mainstat = ui.NewTextbox(18, 1, 0, 11, 0, false, false, "MAIN STAT: Body")
	cm.description.Add(cm.flavourtext, cm.mainstat)

	cm.stats = ui.NewContainer(22, 16, 25, 20, 0, true)
	cm.stats.SetTitle("Stats")
	cm.hp = ui.NewTextbox(22, 1, 0, 1, 0, false, false, "HP: ")
	cm.att = ui.NewTextbox(22, 1, 0, 2, 0, false, false, "ATT: ")
	cm.weapon = ui.NewTextbox(22, 1, 0, 4, 0, false, false, "WEAPON: ")
	cm.armour = ui.NewTextbox(22, 1, 0, 5, 0, false, false, "ARMOUR: ")
	cm.mind = ui.NewTextbox(22, 1, 0, 8, 0, false, false, "MIND: ")
	cm.body = ui.NewTextbox(22, 1, 0, 9, 0, false, false, "BODY: ")
	cm.spirit = ui.NewTextbox(22, 1, 0, 10, 0, false, false, "SPIRIT: ")
	cm.stats.Add(cm.hp, cm.att, cm.weapon, cm.armour, cm.mind, cm.body, cm.spirit)

	cm.confirm = ui.NewButton(30, 1, 32, 40, 2, true, true, "Press Enter to Get Delvin!")
	cm.confirm.Register(&ui.Event{cm.confirm, ui.ACTIVATE, "Create Character!"})

	cm.screen.Add(cm.description, cm.stats, cm.name, cm.class, cm.confirm)
	//END UI STUFF

	cm.activeElem = cm.name
	cm.name.ToggleFocus()

	cm.GenerateCharacter()

	return cm
}

func (cm *CharGenMode) Update() (error, GameModer) {

	for e := ui.PopEvent(); e != nil; e = ui.PopEvent() {
		switch e.ID {
		case ui.CHANGE:
			if e.Caller == cm.class {
				cm.GenerateCharacter()
			}
		case ui.ACTIVATE:
			if e.Caller == cm.confirm {
				dm := NewDelvemode(cm.character, data.NewCaveDungeon())
				return errors.New("Mode Change"), dm
			}
		}
	}

	return nil, nil
}

func (cm *CharGenMode) GenerateCharacter() {
	cm.character = data.NewEntity(0, 0, 0, data.PLAYER)
	cm.character.Name = cm.name.GetText()
	switch cm.class.GetSelection() {
	case 0:
		cm.character.BaseStats.Body = 10 + rand.Intn(5) - 2
		cm.character.BaseStats.Mind = 3 + rand.Intn(5) - 2
		cm.character.BaseStats.Spirit = 5 + rand.Intn(5) - 2
	case 1:
		cm.character.BaseStats.Body = 5 + rand.Intn(5) - 2
		cm.character.BaseStats.Mind = 10 + rand.Intn(5) - 2
		cm.character.BaseStats.Spirit = 3 + rand.Intn(5) - 2
	case 2:
		cm.character.BaseStats.Body = 3 + rand.Intn(5) - 2
		cm.character.BaseStats.Mind = 5 + rand.Intn(5) - 2
		cm.character.BaseStats.Spirit = 10 + rand.Intn(5) - 2
	}

	switch cm.class.GetSelection() {
	case 0:
		cm.flavourtext.ChangeText("The fightman is a muscley man who goes from town to town picking fights. He loves to battle, it gives him a big boner.")
		cm.mainstat.ChangeText("MAIN STAT: Body")
	case 1:
		cm.flavourtext.ChangeText("The Book Nerd has spent most of his Friday nights cuddled around a nice tome, learning how to vaporize his friends who went to the club.")
		cm.mainstat.ChangeText("MAIN STAT: Mind")
	case 2:
		cm.flavourtext.ChangeText("Not to be underestimated, the bald man is a formidable foe. His fervour is fueled by a deep longing for his old hair.")
		cm.mainstat.ChangeText("MAIN STAT: Spirit")
	}

	//update ui elements
	cm.hp.ChangeText("HP: " + strconv.Itoa(cm.character.BaseStats.HP))
	cm.att.ChangeText("ATT: " + strconv.Itoa(cm.character.BaseStats.Attack))
	cm.mind.ChangeText("MIND: " + strconv.Itoa(cm.character.BaseStats.Mind))
	cm.body.ChangeText("BODY: " + strconv.Itoa(cm.character.BaseStats.Body))
	cm.spirit.ChangeText("SPIRIT: " + strconv.Itoa(cm.character.BaseStats.Spirit))
}

func (cm *CharGenMode) Render() {
	cm.screen.Render()
}

func (cm *CharGenMode) HandleKeypress(key sdl.Keycode) {

	if key == sdl.K_TAB {
		cm.CycleUI()
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
		cm.character.Name = cm.name.GetText()
	} else if cm.activeElem == cm.class {
		switch key {
		case sdl.K_DOWN, sdl.K_KP_2:
			cm.class.Next()
		case sdl.K_UP, sdl.K_KP_8:
			cm.class.Prev()
		}
	} else if cm.activeElem == cm.confirm {
		if key == sdl.K_RETURN {
			cm.confirm.Press()
		}
	}
}

func (cm *CharGenMode) CycleUI() {
	switch cm.activeElem {
	case cm.name:
		cm.name.ToggleFocus()
		cm.class.ToggleFocus()
		cm.activeElem = cm.class
	case cm.class:
		cm.class.ToggleFocus()
		cm.confirm.ToggleFocus()
		cm.activeElem = cm.confirm
	case cm.confirm:
		cm.confirm.ToggleFocus()
		cm.activeElem = cm.name
		cm.name.ToggleFocus()
	}
}
