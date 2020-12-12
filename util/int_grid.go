package util

// IntGrid is a basic two dimensional grid of integers.
type IntGrid struct {
	height, width int
	cells         []int
}

// NewIntGrid creates a new IntGrid of the specified size.
func NewIntGrid(height, width int) IntGrid {
	g := IntGrid{height: height, width: width}
	g.cells = make([]int, height*width)
	return g
}

// Height returns the grid's height.
func (g IntGrid) Height() int {
	return g.height
}

// Width returns the grid's height.
func (g IntGrid) Width() int {
	return g.width
}

// Contains returns true when the specified coordinates are inside the grid's bounds.
func (g IntGrid) Contains(i, j int) bool {
	return i >= 0 && i < g.height && j >= 0 && j < g.width
}

// Get returns the value at the specified coordinates in the grid.
func (g IntGrid) Get(i, j int) int {
	return g.cells[i*g.width+j]
}

// Set sets the value at the specified coordinates in the grid.
func (g IntGrid) Set(i, j, x int) {
	g.cells[i*g.width+j] = x
}
