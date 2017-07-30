// Package spiral implements spiral generation funcions based on
// geometrical shapes
package spiral

import (
	"math"

	g "github.com/antavelos/gotools/geometry"
)

// Spiral generates a sequence of points that consist a spiral given a
// starting point, a desired length of the sequence and the distance
// between the points
func Spiral(p g.Point, n int, d float64) chan g.Point {
	c := make(chan g.Point)
	go func() {
		defer close(c)
		c <- p
		var i int
		for {
			f := math.Pow(-1, float64(i))
			for j := 0; j <= ((2 * i) + 1); j++ {
				if j <= i {
					p.MoveX(d * f)
				} else {
					p.MoveY(d * f)
				}
				c <- p

				n--
				if n == 1 {
					return
				}
			}
			i++
		}
	}()
	return c
}
