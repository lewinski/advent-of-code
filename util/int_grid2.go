package util

// Point2 is a two-dimensional point.
type Point2 [2]int

// Origin2 returns the point (0, 0).
func Origin2() Point2 {
	return Point2{0, 0}
}

// Around returns the 8 points adjacent to the point.
func (p Point2) Around() []Point2 {
	return []Point2{
		{p[0] - 1, p[1] - 1},
		{p[0] - 1, p[1]},
		{p[0] - 1, p[1] + 1},

		{p[0], p[1] - 1},
		// {p[0], p[1]},
		{p[0], p[1] + 1},

		{p[0] + 1, p[1] - 1},
		{p[0] + 1, p[1]},
		{p[0] + 1, p[1] + 1},
	}
}

// Scale returns a point with each coordinate multiplied by the specified factor.
func (p Point2) Scale(f int) Point2 {
	return Point2{p[0] * f, p[1] * f}
}

// Offset returns a point resulting from the addition of each point's coordinates.
func (p Point2) Offset(o Point2) Point2 {
	return Point2{p[0] + o[0], p[1] + o[1]}
}

// OffsetCoords returns a point resulting from adding the coordinate offsets to the point.
func (p Point2) OffsetCoords(x, y int) Point2 {
	return Point2{p[0] + x, p[1] + y}
}

// IntGrid2 is a two-dimensional grid of integers.
type IntGrid2 map[Point2]int

// Contains returns true if there is a value in the grid at the position.
func (g IntGrid2) Contains(p Point2) bool {
	_, found := g[p]
	return found
}

// ContainsCoords returns true if there is a value in the grid at the coordinates.
func (g IntGrid2) ContainsCoords(x, y int) bool {
	return g.Contains(Point2{x, y})
}

// Get returns the value at the position, or zero if not set.
func (g IntGrid2) Get(p Point2) int {
	return g[p]
}

// GetCoords returns the value at the coordinates, or zero if not set.
func (g IntGrid2) GetCoords(x, y int) int {
	return g.Get(Point2{x, y})
}

// Set sets the value in the grid at the position.
func (g IntGrid2) Set(p Point2, val int) {
	g[p] = val
}

// SetCoords sets the value in the grid at the coordinates.
func (g IntGrid2) SetCoords(x, y, val int) {
	g.Set(Point2{x, y}, val)
}

// IntGrid2EachFunc is the signature for the function passed to Each.
type IntGrid2EachFunc func(p Point2, x int)

// Each calls the specified function for each cell in the grid.
func (g IntGrid2) Each(eachFunc IntGrid2EachFunc) {
	for p, x := range g {
		eachFunc(p, x)
	}
}
