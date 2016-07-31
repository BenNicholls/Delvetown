package modes

import "github.com/veandco/go-sdl2/sdl"
import "github.com/bennicholls/delvetown/ui"
import "errors"

type GameOverMode struct {
	toobad *ui.Textbox
}

func NewGameOver() *GameOverMode {
	toobad := ui.NewTextbox(20, 1, 15, 15, 0, true, true, "WAY TO GO, IDIOT")
	return &GameOverMode{toobad}
}

func (g *GameOverMode) Update() (error, GameModer) {
	sdl.Delay(10)
	return nil, nil
}

func (g *GameOverMode) Render() {
	g.toobad.Render()
}

func (g *GameOverMode) HandleKeypress(key sdl.Keycode) error {
	return errors.New("quit")
}
