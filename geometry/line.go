package geometry

import (
	"fmt"
)

type Line struct {
	A, B Point
}

// A Line type is repsesented as a tuple of its Points i.e. (1.4, 5.6)
func (l Line) String() string {
	return fmt.Sprintf("(%v, %v)", l.A, l.B)
}

// Transform moves a Line on the plane x and y units on the xx and yy axes
// respectively
func (l *Line) Transform(x float64, y float64) {
	l.A.X += x
	l.A.Y += y
	l.B.X += x
	l.B.Y += y
}
