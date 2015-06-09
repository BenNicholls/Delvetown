package modes

import "github.com/bennicholls/delvetown/data"
import "github.com/bennicholls/delvetown/ui"
import "errors"

type TownMode struct {
	town *data.Town
	text *ui.Textbox
}

func NewTownMode() TownMode {

	tm := TownMode{town: nil}
	_ = tm.LoadTown()

	tm.text = ui.NewTextbox(20, 2, 10, 10, 0, true, "TEST TEXT whatever blah blah blah")

	return tm
}

func (tm *TownMode) LoadTown() error {
	if tm.town != nil {
		return errors.New("Town already loaded.")
	}

	tm.town = data.NewTown(200, 100, "Test Town")

	return nil
}

//eventually put town-saving code here
func (tm *TownMode) UnloadTown() {
	tm.town = nil
}

func (tm TownMode) Update() GameModer {
	return nil
}

func (tm TownMode) Render() {

	tm.text.Render()

}
