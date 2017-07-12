package spiral

import (
	"bytes"
	"github.com/antavelos/gotools/geometry"
	"testing"
)

func TestSpiral(t *testing.T) {
	for _, c := range []struct {
		p    geometry.Point
		n    int
		step float64
		want string
	}{
		{geometry.Point{}, 5, 1.0, "(0, 0)(1, 0)(1, 1)(0, 1)(-1, 1)"},
		{geometry.Point{}, 10, 0.1, "(0, 0)(0.1, 0)(0.1, 0.1)(0, 0.1)(-0.1, 0.1)(-0.1, 0)(-0.1, -0.1)(0, -0.1)(0.1, -0.1)(0.2, -0.1)"},
		{geometry.Point{5, 5}, 5, 0.1, "(5, 5)(5.1, 5)(5.1, 5.1)(5, 5.1)(4.9, 5.1)"},
	} {

		var buf bytes.Buffer
		for p := range Spiral(c.p, c.n, c.step) {
			buf.WriteString(p.String())
		}
		got := buf.String()
		if got != c.want {
			t.Errorf("Spiral(%v, %v, %v) == %q, want %q", c.p, c.n, c.step, got, c.want)
		}
	}
}
