package util

// Point4 is a four-dimensional point.
type Point4 [4]int

// Origin4 returns the point (0, 0, 0, 0).
func Origin4() Point4 {
	return Point4{0, 0, 0, 0}
}

// Around returns the 80 points adjacent to the point
func (p Point4) Around() []Point4 {
	return []Point4{
		{p[0] - 1, p[1] - 1, p[2] - 1, p[3] - 1},
		{p[0] - 1, p[1] - 1, p[2] - 1, p[3]},
		{p[0] - 1, p[1] - 1, p[2] - 1, p[3] + 1},
		{p[0] - 1, p[1] - 1, p[2], p[3] - 1},
		{p[0] - 1, p[1] - 1, p[2], p[3]},
		{p[0] - 1, p[1] - 1, p[2], p[3] + 1},
		{p[0] - 1, p[1] - 1, p[2] + 1, p[3] - 1},
		{p[0] - 1, p[1] - 1, p[2] + 1, p[3]},
		{p[0] - 1, p[1] - 1, p[2] + 1, p[3] + 1},

		{p[0] - 1, p[1], p[2] - 1, p[3] - 1},
		{p[0] - 1, p[1], p[2] - 1, p[3]},
		{p[0] - 1, p[1], p[2] - 1, p[3] + 1},
		{p[0] - 1, p[1], p[2], p[3] - 1},
		{p[0] - 1, p[1], p[2], p[3]},
		{p[0] - 1, p[1], p[2], p[3] + 1},
		{p[0] - 1, p[1], p[2] + 1, p[3] - 1},
		{p[0] - 1, p[1], p[2] + 1, p[3]},
		{p[0] - 1, p[1], p[2] + 1, p[3] + 1},

		{p[0] - 1, p[1] + 1, p[2] - 1, p[3] - 1},
		{p[0] - 1, p[1] + 1, p[2] - 1, p[3]},
		{p[0] - 1, p[1] + 1, p[2] - 1, p[3] + 1},
		{p[0] - 1, p[1] + 1, p[2], p[3] - 1},
		{p[0] - 1, p[1] + 1, p[2], p[3]},
		{p[0] - 1, p[1] + 1, p[2], p[3] + 1},
		{p[0] - 1, p[1] + 1, p[2] + 1, p[3] - 1},
		{p[0] - 1, p[1] + 1, p[2] + 1, p[3]},
		{p[0] - 1, p[1] + 1, p[2] + 1, p[3] + 1},

		{p[0], p[1] - 1, p[2] - 1, p[3] - 1},
		{p[0], p[1] - 1, p[2] - 1, p[3]},
		{p[0], p[1] - 1, p[2] - 1, p[3] + 1},
		{p[0], p[1] - 1, p[2], p[3] - 1},
		{p[0], p[1] - 1, p[2], p[3]},
		{p[0], p[1] - 1, p[2], p[3] + 1},
		{p[0], p[1] - 1, p[2] + 1, p[3] - 1},
		{p[0], p[1] - 1, p[2] + 1, p[3]},
		{p[0], p[1] - 1, p[2] + 1, p[3] + 1},

		{p[0], p[1], p[2] - 1, p[3] - 1},
		{p[0], p[1], p[2] - 1, p[3]},
		{p[0], p[1], p[2] - 1, p[3] + 1},
		{p[0], p[1], p[2], p[3] - 1},
		// {p[0], p[1], p[2], p[3]},
		{p[0], p[1], p[2], p[3] + 1},
		{p[0], p[1], p[2] + 1, p[3] - 1},
		{p[0], p[1], p[2] + 1, p[3]},
		{p[0], p[1], p[2] + 1, p[3] + 1},

		{p[0], p[1] + 1, p[2] - 1, p[3] - 1},
		{p[0], p[1] + 1, p[2] - 1, p[3]},
		{p[0], p[1] + 1, p[2] - 1, p[3] + 1},
		{p[0], p[1] + 1, p[2], p[3] - 1},
		{p[0], p[1] + 1, p[2], p[3]},
		{p[0], p[1] + 1, p[2], p[3] + 1},
		{p[0], p[1] + 1, p[2] + 1, p[3] - 1},
		{p[0], p[1] + 1, p[2] + 1, p[3]},
		{p[0], p[1] + 1, p[2] + 1, p[3] + 1},

		{p[0] + 1, p[1] - 1, p[2] - 1, p[3] - 1},
		{p[0] + 1, p[1] - 1, p[2] - 1, p[3]},
		{p[0] + 1, p[1] - 1, p[2] - 1, p[3] + 1},
		{p[0] + 1, p[1] - 1, p[2], p[3] - 1},
		{p[0] + 1, p[1] - 1, p[2], p[3]},
		{p[0] + 1, p[1] - 1, p[2], p[3] + 1},
		{p[0] + 1, p[1] - 1, p[2] + 1, p[3] - 1},
		{p[0] + 1, p[1] - 1, p[2] + 1, p[3]},
		{p[0] + 1, p[1] - 1, p[2] + 1, p[3] + 1},

		{p[0] + 1, p[1], p[2] - 1, p[3] - 1},
		{p[0] + 1, p[1], p[2] - 1, p[3]},
		{p[0] + 1, p[1], p[2] - 1, p[3] + 1},
		{p[0] + 1, p[1], p[2], p[3] - 1},
		{p[0] + 1, p[1], p[2], p[3]},
		{p[0] + 1, p[1], p[2], p[3] + 1},
		{p[0] + 1, p[1], p[2] + 1, p[3] - 1},
		{p[0] + 1, p[1], p[2] + 1, p[3]},
		{p[0] + 1, p[1], p[2] + 1, p[3] + 1},

		{p[0] + 1, p[1] + 1, p[2] - 1, p[3] - 1},
		{p[0] + 1, p[1] + 1, p[2] - 1, p[3]},
		{p[0] + 1, p[1] + 1, p[2] - 1, p[3] + 1},
		{p[0] + 1, p[1] + 1, p[2], p[3] - 1},
		{p[0] + 1, p[1] + 1, p[2], p[3]},
		{p[0] + 1, p[1] + 1, p[2], p[3] + 1},
		{p[0] + 1, p[1] + 1, p[2] + 1, p[3] - 1},
		{p[0] + 1, p[1] + 1, p[2] + 1, p[3]},
		{p[0] + 1, p[1] + 1, p[2] + 1, p[3] + 1},
	}
}

// Scale returns a point with each coordinate multiplied by the specified factor.
func (p Point4) Scale(f int) Point4 {
	return Point4{p[0] * f, p[1] * f, p[2] * f, p[3] * f}
}

// Offset returns a point resulting from the addition of each point's coordinates.
func (p Point4) Offset(o Point4) Point4 {
	return Point4{p[0] + o[0], p[1] + o[1], p[2] + o[2], p[3] + o[3]}
}

// OffsetCoords returns a point resulting from adding the coordinate offsets to the point.
func (p Point4) OffsetCoords(w, x, y, z int) Point4 {
	return Point4{p[0] + w, p[1] + x, p[2] + y, p[3] + z}
}

// IntGrid4 is a four-dimensional grid of integers.
type IntGrid4 map[Point4]int

// Contains returns true if there is a value in the grid at the position.
func (g IntGrid4) Contains(p Point4) bool {
	_, found := g[p]
	return found
}

// ContainsCoords returns true if there is a value in the grid at the coordinates.
func (g IntGrid4) ContainsCoords(w, x, y, z int) bool {
	return g.Contains(Point4{w, x, y, z})
}

// Get returns the value at the position, or zero if not set.
func (g IntGrid4) Get(p Point4) int {
	return g[p]
}

// GetCoords returns the value at the coordinates, or zero if not set.
func (g IntGrid4) GetCoords(w, x, y, z int) int {
	return g.Get(Point4{w, x, y, z})
}

// Set sets the value in the grid at the position.
func (g IntGrid4) Set(p Point4, val int) {
	g[p] = val
}

// SetCoords sets the value in the grid at the coordinates.
func (g IntGrid4) SetCoords(w, x, y, z, val int) {
	g.Set(Point4{w, x, y, z}, val)
}

// IntGrid4EachFunc is the signature for the function passed to Each.
type IntGrid4EachFunc func(p Point4, x int)

// Each calls the specified function for each cell in the grid.
func (g IntGrid4) Each(eachFunc IntGrid4EachFunc) {
	for p, x := range g {
		eachFunc(p, x)
	}
}
