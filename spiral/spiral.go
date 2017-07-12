// Package spiral implements spiral generation funcions based on
// geometrical shapes
package spiral

import (
	"github.com/antavelos/gotools/geometry"
)

// Spiral generates a sequence of points that consist a spiral given a
// starting point, a desired length of the sequence and the distance
// between the points
func Spiral(p geometry.Point, n int, d float64) chan geometry.Point {
	c := make(chan geometry.Point)
	go func() {
		c <- p
		for i := 0; n > 1; i++ {
			switch i % 2 {
			case 0:
				// east
				for j := 0; j <= i && n > 1; j++ {
					p.Transform(d, 0)
					n--
					c <- p
				}
				// north
				for j := 0; j <= i && n > 1; j++ {
					p.Transform(0, d)
					n--
					c <- p
				}
			case 1:
				// west
				for j := 0; j <= i && n > 1; j++ {
					p.Transform(-d, 0)
					n--
					c <- p
				}
				// south
				for j := 0; j <= i && n > 1; j++ {
					p.Transform(0, -d)
					n--
					c <- p
				}
			}
		}
		close(c)
	}()
	return c
}
