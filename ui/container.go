package ui

import "github.com/bennicholls/delvetown/console"

type Container struct {
	width, height int
	x, y, z       int
	bordered      bool
	title         string
	visible       bool
	redraw        bool

	Elements []UIElem
}

func NewContainer(w, h, x, y, z int, bord bool) *Container {
	return &Container{w, h, x, y, z, bord, "", true, true, make([]UIElem, 0, 20)}
}

func (c *Container) Add(elems ...UIElem) {
	for _, e := range elems {
		c.Elements = append(c.Elements, e)
	}
}

func (c *Container) ClearElements() {
	c.Elements = make([]UIElem, 0, 20)
	c.redraw = true
}

func (c *Container) SetTitle(s string) {
	c.title = s
}

//UIElem that acts as a container for others. Offets (x,y,z, all optional) are passed through
//to the nested elements.
func (c *Container) Render(offset ...int) {
	if c.visible {
		offX, offY, offZ := processOffset(offset)

		if c.redraw {
			console.Clear(c.x+offX, c.y+offY, c.width, c.height)
			c.redraw = false
		}
		for i := 0; i < len(c.Elements); i++ {
			c.Elements[i].Render(c.x+offX, c.y+offY, c.z+offZ)
		}

		if c.bordered {
			console.DrawBorder(c.x+offX, c.y+offY, c.z+offZ, c.width, c.height, c.title)
		}
	}
}

func (c Container) GetDims() (int, int) {
	return c.width, c.height
}

func (c *Container) ToggleVisible() {
	c.visible = !c.visible
}

func (c *Container) SetVisibility(v bool) {
	c.visible = v
}

func (c *Container) MoveTo(x, y, z int) {
	c.x = x
	c.y = y
	c.z = z
}
