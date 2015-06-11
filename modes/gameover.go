package modes

import "github.com/veandco/go-sdl2/sdl"
import "github.com/bennicholls/delvetown/ui"

type GameOverMode struct {
	toobad *ui.Textbox
}

func NewGameOver() *GameOverMode {
	toobad := ui.NewTextbox(20, 1, 15, 15, 0, true, "WAY TO GO, IDIOT")
	return &GameOverMode{toobad}
}

func (g *GameOverMode) Update() GameModer {
	sdl.Delay(10)
	return nil
}

func (g *GameOverMode) Render() {
	g.toobad.Render()
}

func (g *GameOverMode) HandleKeypress(sdl.Keycode) {

}
