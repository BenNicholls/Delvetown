package main

import "github.com/veandco/go-sdl2/sdl"
import "github.com/bennicholls/delvetown/console"
import "github.com/bennicholls/delvetown/modes"
import "math/rand"
import "time"
import "fmt"
import "runtime"

func main() {

	//go-sdl2 requires this to not crash the program for unknown reasons. "fix" incoming eventually apparently.
	runtime.LockOSThread()

	//Set the seed for the RNG. TODO: be able to manually set seed
	rand.Seed(time.Now().UTC().UnixNano())

	var event sdl.Event
	var mode modes.GameModer

	err := console.Setup(96, 54, "res/curses.bmp", "Delvetown")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer console.Cleanup()

	mode = modes.NewMainMenu()

	running := true

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
		err, m := mode.Update()
		if err != nil {
			switch err.Error() {
				case "Mode Change":
					if m != nil {
						console.Clear()
						mode = m
					}

				case "Quit":
					return
			}
		}

		//Push changes to console
		mode.Render()

		//Render to screen NOTE: put this in a channel maybe? maybe with the mode.render?
		//reconsider when animations and UI effects go in.
		console.Render()
	}
}
