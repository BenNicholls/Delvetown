package modes

import "github.com/bennicholls/delvetown/data"
import "github.com/bennicholls/delvetown/console"
import "errors"
import "math/rand"

//eventually move this to a more general "modes.go" file
type GameModer interface {
	Update()
	Render()
}

type TownMode struct {
	town *data.Town
}

func NewTownMode() TownMode {

	tm := TownMode{town: nil}
	_ = tm.LoadTown()

	return tm
}

type Camera struct {
	width, height int
	x, y int
}

//Load a town. For now though, this will create our test town
func (tm *TownMode) LoadTown() error {
	if tm.town != nil {
		return errors.New("Town already loaded.")
	}

	tm.town = data.NewTown(200, 100, "Test Town")
	for i := 0; i < tm.town.Width*tm.town.Height; i++ {
		tm.town.Townmap.ChangeTileType(i%tm.town.Width, i/tm.town.Width, rand.Intn(2))
	}

	return nil
}

//eventually put town-saving code here
func (tm *TownMode) UnloadTown() {
	tm.town = nil
}

func (tm TownMode) Update() {

}

func (tm TownMode)  Render() {
	cw, cy := console.GetDims()
	for i := 0; i < cw*cy; i++ {
		t := tm.town.Townmap.GetTileType(i%cw, i/cw)
		console.DrawTile(i%cw, i/cw, t)
	}

	console.Render()
}

