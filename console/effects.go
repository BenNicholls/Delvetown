package console

//Inverts the foreground and background colours
func Invert(x, y, z int) {
	if x < width && x >= 0 && y < height && y >= 0 {
		s := y*width + x
		if grid[s].Z > z {
			return
		}
		f, b := grid[s].ForeColour, grid[s].BackColour
		ChangeBackColour(x, y, f)
		ChangeForeColour(x, y, b)
	}
}
