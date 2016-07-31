package ui

import "github.com/bennicholls/delvetown/console"

type Animator interface {
	Tick()
	Render(offset ...int)
	Toggle()
}

type BlinkAnimation struct {
	tick  int
	speed int //number of frames between blinks
	x, y  int //position (possibly relative to element or container)
	enabled bool
}

func NewBlinkAnimation(x, y, speed int) *BlinkAnimation {
	return &BlinkAnimation{0, speed, x, y, true}
}

func (ba *BlinkAnimation) Toggle() {
	ba.enabled = !ba.enabled
	ba.tick = 0
}

func (ba *BlinkAnimation) Tick() {
	if ba.enabled {
		ba.tick++
	}
}

func (ba *BlinkAnimation) Render(offset ...int) {
	if ba.enabled {
		if ba.tick%(ba.speed*2) < ba.speed {
			offX, offY, offZ := processOffset(offset)
			console.Invert(ba.x+offX, ba.y+offY, offZ)
		}
	}
}
