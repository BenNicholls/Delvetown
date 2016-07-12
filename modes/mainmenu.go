package modes

import "github.com/veandco/go-sdl2/sdl"
import "github.com/bennicholls/delvetown/ui"

type MainMenuMode struct {
	menu *ui.List
	enter bool
}

func NewMainMenu() *MainMenuMode {
	menu := ui.NewList(20, 3, 25, 20, 0, true, "WAY TO GO, IDIOT")
	menu.SetTitle("Menu")
	menu.Append("New Game")
	menu.Append("High Scores")
	menu.Append("Quit")
	return &MainMenuMode{menu, false}
}

func (mm *MainMenuMode) Update() GameModer {
	if mm.enter {
		mm.enter = false
		switch mm.menu.GetSelection() {
			case 0:
				return NewDelvemode()
		}
	}

	return nil
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
		mm.enter = true
    }
}
