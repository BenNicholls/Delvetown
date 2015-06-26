package data

import "github.com/bennicholls/delvetown/util"

//TODO: replace ugly 8 calls for a nice loop over a nice array. Nice.
// fn is a function for the shadowcaster to apply to open spaces it finds.
func (m *TileMap) ShadowCast(x, y, strength int, fn Cast) {
	fn(m, x, y, 0, strength*strength)
	m.Scan(x, y, 1, 1.0, 0.0, strength, []int{1, 0, 0, 1}, fn)
	m.Scan(x, y, 1, 1.0, 0.0, strength, []int{0, 1, 1, 0}, fn)
	m.Scan(x, y, 1, 1.0, 0.0, strength, []int{0, -1, 1, 0}, fn)
	m.Scan(x, y, 1, 1.0, 0.0, strength, []int{-1, 0, 0, 1}, fn)
	m.Scan(x, y, 1, 1.0, 0.0, strength, []int{-1, 0, 0, -1}, fn)
	m.Scan(x, y, 1, 1.0, 0.0, strength, []int{0, -1, -1, 0}, fn)
	m.Scan(x, y, 1, 1.0, 0.0, strength, []int{0, 1, -1, 0}, fn)
	m.Scan(x, y, 1, 1.0, 0.0, strength, []int{1, 0, 0, -1}, fn)
}

//TODO: General cleanup. Direct port from python, not exactly golangish.
//TODO: FInd some way to make it not do the diagonals twice, it's causing artifacts
func (m *TileMap) Scan(x, y, row int, slope1, slope2 float32, radius int, t []int, fn Cast) {
	if slope1 < slope2 {
		return
	}
	//scan #radius rows
	for j := row; j < radius+1; j++ {
		dx, dy := -j, -j
		blocked := false
		newStart := slope1

		//scan row
		for ; dx <= 0; dx++ {
			mx, my := x+dx*t[0]+dy*t[1], y+dx*t[2]+dy*t[3] //map coordinates
			if !util.CheckBounds(mx, my, m.width, m.height) {
				continue
			}
			lSlope, rSlope := (float32(dx)-0.5)/(float32(dy)+0.5), (float32(dx)+0.5)/(float32(dy)-0.5)

			if newStart < rSlope {
				continue
			} else if slope2 > lSlope {
				break
			} else {
				if d := util.Distance(0, dx, 0, dy); d < radius*radius {
					fn(m, mx, my, d, radius*radius) //BRIGHTNESS CALC)
				}
				//scanning a block
				if blocked {
					if m.tiles[mx+my*m.width].Transparent() {
						blocked = false
						slope1 = newStart
					} else {
						newStart = rSlope
					}
				} else {
					//blocked square, commence child scan
					if !m.tiles[mx+my*m.width].Transparent() && j < radius {
						blocked = true
						m.Scan(x, y, j+1, newStart, lSlope, radius, t, fn)
						newStart = rSlope
					}
				}
			}
		}
		if blocked {
			break
		}
	}

}

//type specifying precisely what you can pass to the shadowcaster.
//parameters here are the info that the shadowcaster will deliver,
//do the rest via a function closure (see dm.MemoryCast())
type Cast func(m *TileMap, x, y, d, r int)

//Run this over the levelmap to light squares. Linearly interpolates
//from max (255) at center to 0 at radius.
func Light(m *TileMap, x, y, d, r int) {
	m.tiles[x+y*m.width].Light.Bright += (255 - int(255*float32(d)/float32(r)))
}

//Same as above, but opposite.
func Darken(m *TileMap, x, y, d, r int) {
	m.tiles[x+y*m.width].Light.Bright -= (255 - int(255*float32(d)/float32(r)))

	if m.tiles[x+y*m.width].Light.Bright < 0 {
		m.tiles[x+y*m.width].Light.Bright = 0
	}
}
