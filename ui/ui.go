package ui

import "github.com/bennicholls/delvetown/console"

type UIElem interface {
	Render(offset ...int)
	GetDims() (int, int)
}

type Container struct {
	width, height int
	x, y, z       int
	bordered      bool
	title         string

	Elements []UIElem
}

func NewContainer(w, h, x, y, z int, bord bool) *Container {
	return &Container{w, h, x, y, z, bord, "", make([]UIElem, 0, 20)}
}

func (c *Container) Add(elem UIElem) {
	c.Elements = append(c.Elements, elem)
}

func (c *Container) SetTitle(s string) {
	c.title = s
}

//UIElem that acts as a container for others. Offets (x,y,z, all optional) are passed through
//to the nested elements.
func (c *Container) Render(offset ...int) {
	offX, offY, offZ := 0, 0, 0
	if len(offset) >= 2 {
		offX, offY = offset[0], offset[1]
		if len(offset) == 3 {
			offZ = offset[2]
		}
	}
	for i := 0; i < len(c.Elements); i++ {
		c.Elements[i].Render(c.x+offX, c.y+offY, c.z+offZ)
	}

	if c.bordered {
		console.DrawBorder(c.x+offX, c.y+offY, c.z+offZ, c.width, c.height, c.title)
	}
}

func (c Container) GetDims() (int, int) {
	return c.width, c.height
}

//TODO: Split this into separate files for each UI element
