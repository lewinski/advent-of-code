package util

// Point3 is a three-dimensional point.
type Point3 [3]int

// Origin3 returns the point (0, 0, 0).
func Origin3() Point3 {
	return Point3{0, 0, 0}
}

// Around returns the 26 points adjacent to the point
func (p Point3) Around() []Point3 {
	return []Point3{
		{p[0] - 1, p[1] - 1, p[2] - 1},
		{p[0] - 1, p[1] - 1, p[2]},
		{p[0] - 1, p[1] - 1, p[2] + 1},
		{p[0] - 1, p[1], p[2] - 1},
		{p[0] - 1, p[1], p[2]},
		{p[0] - 1, p[1], p[2] + 1},
		{p[0] - 1, p[1] + 1, p[2] - 1},
		{p[0] - 1, p[1] + 1, p[2]},
		{p[0] - 1, p[1] + 1, p[2] + 1},

		{p[0], p[1] - 1, p[2] - 1},
		{p[0], p[1] - 1, p[2]},
		{p[0], p[1] - 1, p[2] + 1},
		{p[0], p[1], p[2] - 1},
		// {p[0], p[1], p[2]},
		{p[0], p[1], p[2] + 1},
		{p[0], p[1] + 1, p[2] - 1},
		{p[0], p[1] + 1, p[2]},
		{p[0], p[1] + 1, p[2] + 1},

		{p[0] + 1, p[1] - 1, p[2] - 1},
		{p[0] + 1, p[1] - 1, p[2]},
		{p[0] + 1, p[1] - 1, p[2] + 1},
		{p[0] + 1, p[1], p[2] - 1},
		{p[0] + 1, p[1], p[2]},
		{p[0] + 1, p[1], p[2] + 1},
		{p[0] + 1, p[1] + 1, p[2] - 1},
		{p[0] + 1, p[1] + 1, p[2]},
		{p[0] + 1, p[1] + 1, p[2] + 1},
	}
}

// Scale returns a point with each coordinate multiplied by the specified factor.
func (p Point3) Scale(f int) Point3 {
	return Point3{p[0] * f, p[1] * f, p[2] * f}
}

// Offset returns a point resulting from the addition of each point's coordinates.
func (p Point3) Offset(o Point3) Point3 {
	return Point3{p[0] + o[0], p[1] + o[1], p[2] + o[2]}
}

// OffsetCoords returns a point resulting from adding the coordinate offsets to the point.
func (p Point3) OffsetCoords(x, y, z int) Point3 {
	return Point3{p[0] + x, p[1] + y, p[2] + z}
}

// IntGrid3 is a three-dimensional grid of integers.
type IntGrid3 map[Point3]int

// Contains returns true if there is a value in the grid at the position.
func (g IntGrid3) Contains(p Point3) bool {
	_, found := g[p]
	return found
}

// ContainsCoords returns true if there is a value in the grid at the coordinates.
func (g IntGrid3) ContainsCoords(x, y, z int) bool {
	return g.Contains(Point3{x, y, z})
}

// Get returns the value at the position, or zero if not set.
func (g IntGrid3) Get(p Point3) int {
	return g[p]
}

// GetCoords returns the value at the coordinates, or zero if not set.
func (g IntGrid3) GetCoords(x, y, z int) int {
	return g.Get(Point3{x, y, z})
}

// Set sets the value in the grid at the position.
func (g IntGrid3) Set(p Point3, val int) {
	g[p] = val
}

// SetCoords sets the value in the grid at the coordinates.
func (g IntGrid3) SetCoords(x, y, z, val int) {
	g.Set(Point3{x, y, z}, val)
}

// IntGrid3EachFunc is the signature for the function passed to Each.
type IntGrid3EachFunc func(p Point3, x int)

// Each calls the specified function for each cell in the grid.
func (g IntGrid3) Each(eachFunc IntGrid3EachFunc) {
	for p, x := range g {
		eachFunc(p, x)
	}
}
