package main

import "github.com/veandco/go-sdl2/sdl"
import "github.com/bennicholls/delvetown/console"
import "github.com/bennicholls/delvetown/modes"
import "fmt"

func main() {

	var event sdl.Event

	console.Setup(100, 50, 16)

	m := modes.NewDelveMode()

	mode := modes.GameModer(m)

	running := true
	frames := 0

	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			// case *sdl.MouseMotionEvent:
			// 	fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
			// 		t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)
			// case *sdl.MouseButtonEvent:
			// 	fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
			// 		t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
			// case *sdl.MouseWheelEvent:
			// 	fmt.Printf("[%d ms] MouseWheel\ttype:%d\tid:%d\tx:%d\ty:%d\n",
			// 		t.Timestamp, t.Type, t.Which, t.X, t.Y)
			case *sdl.KeyUpEvent:
				//fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
				//	t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
				mode.HandleKeypress(t.Keysym.Sym)
			}
		}

		//Tick the game
		mode.Update()

		//Push changes to console
		mode.Render()

		//Render to screen NOTE: put this in a channel maybe? maybe with the mode.render?
		//reconsider when animations and UI effects go in.
		console.Render()

		frames += 1

		if frames%5000 == 0 {
			fmt.Printf("%d fps\n", frames*1000/int(sdl.GetTicks()))
		}
	}

	console.Cleanup()
}
