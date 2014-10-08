package console

import "github.com/veandco/go-sdl2/sdl"
import "fmt"
import "os"
import "math/rand"

var window *sdl.Window
var renderer *sdl.Renderer
var sprites *sdl.Texture
var format *sdl.PixelFormat

var width, height, tileSize int

var grid []GridPoint
var masterDirty bool //is this necessary?

type GridPoint struct {
	Glyph int
	ForeColour uint32
	BackColour uint32
	Z int 
	Dirty bool
}

func (g *GridPoint) Set(gl int, fore, back uint32, z int) {
	if g.Glyph != gl || g.ForeColour != fore || g.BackColour != back {
		g.Glyph = gl
		g.ForeColour = fore
		g.BackColour = back
		g.Z = z
		g.Dirty = true
	}
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

	grid = make([]GridPoint, width*height)
	masterDirty = true
}

func Render() {
	if masterDirty {
		var src, dst sdl.Rect

		for i, s := range grid {
			if s.Dirty {
				dst = makeRect((i%width)*tileSize, (i/width)*tileSize, tileSize, tileSize)
				src = makeRect((s.Glyph%16)*tileSize, (s.Glyph/16)*tileSize, tileSize, tileSize)

				renderer.SetDrawColor(sdl.GetRGBA(s.BackColour, format))
				renderer.FillRect(&dst)

				sprites.SetColorMod(sdl.GetRGB(s.ForeColour, format))
				renderer.Copy(sprites, &src, &dst)

				grid[i].Dirty = false
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
	if x >= width || y >= height {
		return
	}
	if grid[y*width + x].Glyph != glyph {
		grid[y*width + x].Glyph = glyph	
		grid[y*width + x].Dirty = true
		masterDirty = true
	}
}

func ChangeForeColour(x, y int, fore uint32) {
	if x >= width || y >= height {
		return
	}
	if grid[y*width + x].ForeColour != fore {
		grid[y*width + x].ForeColour = fore
		grid[y*width + x].Dirty = true
		masterDirty = true
	}
}

func ChangeBackColour(x, y int, back uint32) {
	if x >= width || y >= height {
		return
	}
	if grid[y*width + x].BackColour != back {
		grid[y*width + x].BackColour = back
		grid[y*width + x].Dirty = true
		masterDirty = true
	}
}

func ChangeGridPoint(x, y, z, glyph int, fore, back uint32) {
	s := y*width + x
	if x >= width || y >= height || grid[s].Z > z {
		return
	}
	grid[s].Set(glyph, fore, back, z)
	masterDirty = true
}

//TODO: border glyph merging, custom colouring, multiple styles, title text
func DrawBorder(x, y, z, w, h int) {
	for i := 0; i < w; i++ {
		ChangeGridPoint(x + i, y - 1, z, 0xc4, 0xFFFFFF, 0x000000)
		ChangeGridPoint(x + i, y + h, z, 0xc4, 0xFFFFFF, 0x000000)
	}
	for i := 0; i < h; i++ {
		ChangeGridPoint(x - 1, y + i, z, 0xb3, 0xFFFFFF, 0x000000)
		ChangeGridPoint(x + w, y + i, z, 0xb3, 0xFFFFFF, 0x000000)
	}
	ChangeGridPoint(x - 1, y - 1, z, 0xda, 0xFFFFFF, 0x000000)
	ChangeGridPoint(x - 1, y + h, z, 0xc0, 0xFFFFFF, 0x000000)
	ChangeGridPoint(x + w, y + h, z, 0xd9, 0xFFFFFF, 0x000000)
	ChangeGridPoint(x + w, y - 1, z, 0xbf, 0xFFFFFF, 0x000000)

}

func GetDims() (w, h int) {
	return width, height
}

//Test function.
func SpamGlyphs() {
	for n := 0; n < 100; n++ {
		x := rand.Intn(width)
		y := rand.Intn(height)
		ChangeGridPoint(x, y, 0, rand.Intn(255), sdl.MapRGBA(format, 0, 255, 0, 50), sdl.MapRGBA(format, 255, 0, 0, 255))
	}
}