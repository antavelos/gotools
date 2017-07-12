package spiral

import (
	"github.com/antavelos/gotools/geometry"
)

func Spiral(p geometry.Point, n int, step float64) chan geometry.Point {
	c := make(chan geometry.Point)
	go func() {
		c <- p
		for i := 0; n > 1; i++ {
			switch i % 2 {
			case 0:
				// east
				for j := 0; j <= i && n > 1; j++ {
					p.Step(step, 0)
					n--
					c <- p
				}
				// north
				for j := 0; j <= i && n > 1; j++ {
					p.Step(0, step)
					n--
					c <- p
				}
			case 1:
				// west
				for j := 0; j <= i && n > 1; j++ {
					p.Step(-step, 0)
					n--
					c <- p
				}
				// south
				for j := 0; j <= i && n > 1; j++ {
					p.Step(0, -step)
					n--
					c <- p
				}
			}
		}
		close(c)
	}()
	return c
}
