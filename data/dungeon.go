package data

type Dungeon struct {
    name string
    desc string
    intro string
    depth int

    generators []func() //generation functions for the levels. each needs to return a *Level
    levels []*Level //once generated, levels go here.
}

