package modes

import "github.com/bennicholls/delvetown/ui"
import "github.com/bennicholls/delvetown/data"
import "github.com/bennicholls/delvetown/entities"
import "github.com/veandco/go-sdl2/sdl"
import "fmt"
import "strconv"

type DelveMode struct {

	//UI stuff!
	logbox  *ui.Textbox   //combat and action log
	view    *ui.TileView  //main map area
	sidebar *ui.Container //holds the various ui elements for the player's info display.

	//sidebar ui elements
	hp *ui.Textbox

	level *data.Level

	xCamera, yCamera int

	player *entities.Entity
	log    string

	tick int

	//player offset
	pDX, pDY int
}

func NewDelveMode() *DelveMode {
	dm := new(DelveMode)

	//UI stuff
	dm.logbox = ui.NewTextbox(98, 8, 1, 1, 0, true, "TEST TEXT whatever blah blah blahxxx")
	dm.view = ui.NewTileView(83, 39, 1, 10, 0, true)
	dm.sidebar = ui.NewContainer(14, 39, 85, 10, 0, true)
	dm.hp = ui.NewTextbox(10, 1, 0, 0, 0, false, "hello")

	dm.sidebar.Add(dm.hp)

	//Level Up!
	dm.level = data.NewLevel(50, 50)
	dm.level.GenerateArena(30, 30)
	dm.player = dm.level.Player

	dm.pDX, dm.pDY = 0, 0
	dm.tick = 0

	return dm
}

func (dm *DelveMode) HandleKeypress(key sdl.Keycode) {
	switch key {
	case sdl.K_DOWN:
		dm.pDY = 1
	case sdl.K_UP:
		dm.pDY = -1
	case sdl.K_LEFT:
		dm.pDX = -1
	case sdl.K_RIGHT:
		dm.pDX = 1
	}
}

func (dm *DelveMode) Update() {

	//player movement
	if dm.pDX != 0 || dm.pDY != 0 {
		dm.level.MovePlayer(dm.pDX, dm.pDY)
		dm.logbox.ChangeText(fmt.Sprint("(", dm.pDX, ", ", dm.pDY, ")"))
		dm.pDX, dm.pDY = 0, 0
	}

	dm.hp.ChangeText("HP: " + strconv.Itoa(dm.player.Health))
	dm.tick++
}

func (dm *DelveMode) Render() {

	w, h := dm.view.Width, dm.view.Height
	dm.xCamera, dm.yCamera = dm.player.X-dm.view.Width/2, dm.player.Y-dm.view.Height/2

	var e *entities.Entity
	for i := 0; i < w*h; i++ {
		x, y := i%w+dm.xCamera, i/w+dm.yCamera
		e = dm.level.Levelmap.GetEntity(x, y)
		if e != nil {
			dm.view.DrawEntity(i%w, i/w, e.Glyph)
		} else {
			t := dm.level.Levelmap.GetTileType(x, y)
			dm.view.DrawTile(i%w, i/w, t)
		}
	}
	dm.logbox.Render()
	dm.view.Render()
	dm.sidebar.Render()
}
