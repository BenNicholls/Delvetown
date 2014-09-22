package console

import "github.com/veandco/go-sdl2/sdl"
import "fmt"
import "os"
import "math/rand"

var window *sdl.Window
var renderer *sdl.Renderer
var sprites *sdl.Texture

var width, height, tileSize int

var grid []square

//NOTE: rename this sometime. square? come on.
type square struct {
	glyph int
	foreColour uint32
	backColour uint32
	dirty bool
}

//Setup the game window, renderer, etc TODO: have this function emit errors
func Setup(w int, h int, size int) {

	width = w
	height = h
	tileSize = size

	window = sdl.CreateWindow("Delvetown", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, width*tileSize, height*tileSize, sdl.WINDOW_SHOWN)
	if window == nil {
		fmt.Println("Failed to create window: %s\n", sdl.GetError())
		os.Exit(1)
	}

	renderer = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if renderer == nil {
		fmt.Println("Failed to create renderer: %s\n", sdl.GetError())
		os.Exit(2)
	}

	image := sdl.LoadBMP("res/curses.bmp")
	if image == nil {
		fmt.Println("Failed to load image: %s\n", sdl.GetError())
		os.Exit(2)	
	}

	image.SetColorKey(1, 0xFF00FF)

	sprites = renderer.CreateTextureFromSurface(image)
	if sprites == nil {
		fmt.Println("Failed to create sprite texture: %s\n", sdl.GetError())
		os.Exit(2)	
	}

	image.Free()

	grid = make([]square, w*h)

}

func Render() {

	r := rand.Intn(width*height)
	grid[r].glyph = rand.Intn(255)
	grid[r].dirty = true

	for i, s := range grid {
		if s.dirty {

			dst := makeRect((i%width)*tileSize, (i/height)*tileSize, tileSize, tileSize)

			renderer.SetDrawColor(255, uint8(rand.Intn(255)),uint8(rand.Intn(255)), uint8(rand.Intn(255)))
			renderer.FillRect(&dst)

			src := makeRect((s.glyph%16)*tileSize, (s.glyph/16)*tileSize, tileSize, tileSize)
			renderer.Copy(sprites, &src, &dst)

			grid[i].dirty = false
		}
	}

	renderer.Present()
}

//int32 for rect arguments. what a world.
func makeRect(x, y, w, h int) sdl.Rect {
	return sdl.Rect{int32(x), int32(y), int32(w), int32(h)}
}

func Cleanup() {
	sprites.Destroy()
	renderer.Destroy()
	window.Destroy()
}