package geometry

import (
	"fmt"
)

type Point struct {
	x, y float64
}

func (p Point) String() string {
	return fmt.Sprintf("(%v, %v)", p.x, p.y)
}

func (p *Point) Step(x float64, y float64) {
	p.x += x
	p.y += y
}
