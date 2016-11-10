package modes

import "github.com/veandco/go-sdl2/sdl"
import "github.com/bennicholls/delvetown/ui"
import "errors"

type MainMenuMode struct {
	menu *ui.List
}

func NewMainMenu() *MainMenuMode {
	menu := ui.NewList(20, 3, 25, 20, 0, true, "WAY TO GO, IDIOT")
	menu.SetTitle("Menu")
	menu.Append("New Game", "High Scores", "Quit")
	return &MainMenuMode{menu}
}

func (mm *MainMenuMode) Update() (error, GameModer) {

	for e := ui.PopEvent(); e != nil; e = ui.PopEvent() {
		switch e.ID {
		case ui.ACTIVATE:
			switch mm.menu.GetSelection() {
			case 0:
				return errors.New("Mode Change"), NewCharGen()
			case 2:
				return errors.New("Quit"), nil
			}
		}
	}

	return nil, nil
}

func (mm *MainMenuMode) Render() {
	mm.menu.Render()
}

func (mm *MainMenuMode) HandleKeypress(key sdl.Keycode) {
	switch key {
	case sdl.K_DOWN, sdl.K_KP_2:
		mm.menu.Next()
	case sdl.K_UP, sdl.K_KP_8:
		mm.menu.Prev()
	case sdl.K_RETURN, sdl.K_KP_ENTER:
		ui.PushEvent(mm.menu, ui.ACTIVATE, "go")
	}
}
