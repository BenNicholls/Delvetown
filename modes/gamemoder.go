package modes

import "github.com/veandco/go-sdl2/sdl"

//eventually move this to a more general "modes.go" file
type GameModer interface {
	Update() GameModer
	Render()
	HandleKeypress(sdl.Keycode)
}