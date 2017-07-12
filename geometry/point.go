package geometry

import (
	"fmt"
)

type Point struct {
	X, Y float64
}

func (p Point) String() string {
	return fmt.Sprintf("(%v, %v)", p.X, p.Y)
}

func (p *Point) Step(x float64, y float64) {
	p.X += x
	p.Y += y
}
