package modes

import "github.com/veandco/go-sdl2/sdl"

type GameModer interface {
	Update() (error, GameModer)
	Render()
	HandleKeypress(sdl.Keycode) error
}
