package main

import "github.com/veandco/go-sdl2/sdl"
import "github.com/bennicholls/delvetown/console"
import "github.com/bennicholls/delvetown/modes"
import "github.com/bennicholls/delvetown/modes/delvemode"
import "math/rand"
import "time"

func main() {

	//Set the seed for the RNG. TODO: be able to manually set seed
	rand.Seed(time.Now().UTC().UnixNano())

	var event sdl.Event
	var mode modes.GameModer

	console.Setup(100, 50, 16)
	defer console.Cleanup()

	mode = delvemode.New()

	//mode := modes.GameModer(m)

	running := true
	frames := 0
	frameTime, ticks := uint32(0), uint32(0)
	fps := uint32(17) //17ms = 60 FPS approxs

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
		m := mode.Update()
		if m != nil {
			console.Clear()
			mode = m
		}
		//Push changes to console
		mode.Render()

		//Render to screen NOTE: put this in a channel maybe? maybe with the mode.render?
		//reconsider when animations and UI effects go in.
		console.Render()

		frames += 1

		//framerate printer thing. super tacky.
		// if frames%50 == 0 {
		// 	fmt.Printf("%d fps\n", frames*1000/int(sdl.GetTicks()))
		// }

		//framerate limiter. keeps the cpu usage down, you know?
		ticks = sdl.GetTicks() - frameTime
		if ticks < fps {
			sdl.Delay(fps - ticks)
		}
		frameTime = sdl.GetTicks()
	}

}
