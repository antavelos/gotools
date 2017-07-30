package geometry

import (
	"fmt"
)

// A Point type represents a point in the plane and consists
// of two coordinates
type Point struct {
	X, Y float64
}

// A Point type is repsesented as a tuple of its coordinates i.e. (1.4, 5.6)
func (p Point) String() string {
	return fmt.Sprintf("(%v, %v)", p.X, p.Y)
}

// Step moves a Point on the plane x and y units on the xx and yy axes
// respectively
func (p *Point) Transform(x float64, y float64) {
	p.X += x
	p.Y += y
}

func (p *Point) MoveY(d float64) {
	p.Transform(0, d)
}

func (p *Point) MoveX(d float64) {
	p.Transform(d, 0)
}
