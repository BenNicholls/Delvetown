package modes

import "github.com/veandco/go-sdl2/sdl"
import "github.com/bennicholls/delvetown/ui"

type WinnerMode struct {
	winner *ui.Textbox
}

func NewWinner() *WinnerMode {
	win := ui.NewTextbox(40, 1, 15, 15, 0, true, true, "YOU DEFEATED THE RAVAGING HORDE!")
	return &WinnerMode{win}
}

func (g *WinnerMode) Update() GameModer {
	sdl.Delay(50)
	return nil
}

func (g *WinnerMode) Render() {
	g.winner.Render()
}

func (g *WinnerMode) HandleKeypress(key sdl.Keycode) {

}
