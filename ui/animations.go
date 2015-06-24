package ui

import "github.com/bennicholls/delvetown/console"

type Animator interface {
	Tick()
	Render(offset ...int)
}

type BlinkAnimation struct {
	tick  int
	speed int //number of frames between blinks
	x, y  int //position (possibly relative to element or container)
}

func NewBlinkAnimation(x, y, speed int) *BlinkAnimation {
	return &BlinkAnimation{0, speed, x, y}
}

func (ba *BlinkAnimation) Tick() {
	ba.tick++
}

func (ba *BlinkAnimation) Render(offset ...int) {

	if ba.tick%(ba.speed*2) < ba.speed {
		offX, offY, offZ := processOffset(offset)
		console.Invert(ba.x+offX, ba.y+offY, offZ)
	}
}
