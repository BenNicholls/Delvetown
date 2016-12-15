package data

type Dungeon struct {
    Name string
    Desc string
    Intro string
    Depth int

    Generators []Generator //generation functions for the levels. each needs to return a *Level
    Levels []*Level //once generated, levels go here.
}

type Generator func(p *Entity, w, h int) *Level

func NewCaveDungeon() *Dungeon {
    d := new(Dungeon)
    d.Name = "Cave"
    d.Desc = "A big scary cave, how exciting!"
    d.Intro = "You find the mouth of the big scary cave. You can hear dripping water and the sound of bons being gnawed on by things that are probably gross."
    d.Depth = 5

    d.Generators = make([]Generator, d.Depth)
    d.Levels = make([]*Level, d.Depth)
    
    d.Generators[0] = GenerateCave
    
    return d
}