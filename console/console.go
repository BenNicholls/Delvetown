package console

import "github.com/veandco/go-sdl2/sdl"
import "fmt"
import "os"
import "math/rand"
import "github.com/bennicholls/delvetown/data"

var window *sdl.Window
var renderer *sdl.Renderer
var sprites *sdl.Texture
var format *sdl.PixelFormat

var width, height, tileSize int

var grid []square
var masterDirty bool

//NOTE: rename this sometime. square? come on.
type square struct {
	glyph int
	foreColour uint32
	backColour uint32
	dirty bool
}

//Setup the game window, renderer, etc TODO: have this function emit errors instead of just borking the program
func Setup(w, h, size int) {

	width = w
	height = h
	tileSize = size

	window = sdl.CreateWindow("Delvetown", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, width*tileSize, height*tileSize, sdl.WINDOW_SHOWN)
	if window == nil {
		fmt.Println("Failed to create window: %s\n", sdl.GetError())
		os.Exit(1)
	}

	var err error
	format, err = sdl.AllocFormat(uint(window.GetPixelFormat()))
	if err != nil {
		fmt.Println("No pixelformat: %s\n", sdl.GetError())
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

	grid = make([]square, width*height)
	masterDirty = true

}

func Render() {
	if masterDirty {
		var src, dst sdl.Rect

		for i, s := range grid {
			if s.dirty {
				dst = makeRect((i%width)*tileSize, (i/width)*tileSize, tileSize, tileSize)
				src = makeRect((s.glyph%16)*tileSize, (s.glyph/16)*tileSize, tileSize, tileSize)

				renderer.SetDrawColor(sdl.GetRGBA(s.backColour, format))
				renderer.FillRect(&dst)

				sprites.SetColorMod(sdl.GetRGB(s.foreColour, format))
				renderer.Copy(sprites, &src, &dst)

				grid[i].dirty = false
			}
		}

		renderer.Present()
		masterDirty = false
	}
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

func ChangeGlyph(x, y, glyph int) {
	if (x >= width || y >= height) {
		return
	}
	if grid[y*width + x].glyph != glyph {
		grid[y*width + x].glyph = glyph	
		grid[y*width + x].dirty = true
		masterDirty = true
	}
}

func ChangeForeColour(x, y int, fore uint32) {
	if (x >= width || y >= height) {
		return
	}
	if grid[y*width + x].foreColour != fore {
		grid[y*width + x].foreColour = fore
		grid[y*width + x].dirty = true
		masterDirty = true
	}
}

func ChangeBackColour(x, y int, back uint32) {
	if (x >= width || y >= height) {
		return
	}
	if grid[y*width + x].backColour != back {
		grid[y*width + x].backColour = back
		grid[y*width + x].dirty = true
		masterDirty = true
	}
}

func ChangeSquare(x, y, glyph int, fore, back uint32) {
	if (x >= width || y >= height) {
		return
	}
	s := y*width + x
	if grid[s].glyph != glyph || grid[s].foreColour != fore || grid[s].backColour != back {
		grid[s].glyph = glyph
		grid[s].foreColour = fore
		grid[s].backColour = back
		grid[s].dirty = true
		masterDirty = true
	}
}

//takes (x,y) and a tiletype 
func DrawTile(x, y, t int) {
	v := data.GetVisuals(t)
	ChangeSquare(x, y, v.Glyph, v.ForeColour, v.BackColour)
}

func GetDims() (w, h int) {
	return width, height
}

//Test function.
func SpamGlyphs() {
	for n := 0; n < 100; n++ {
		x := rand.Intn(width)
		y := rand.Intn(height)
		ChangeSquare(x, y, rand.Intn(255), sdl.MapRGBA(format, 0, 255, 0, 50), sdl.MapRGBA(format, 255, 0, 0, 255))
	}
}