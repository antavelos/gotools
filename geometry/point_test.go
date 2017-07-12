package geometry

import (
	"testing"
)

func TestPointString(t *testing.T) {
	for _, c := range []struct {
		in   Point
		want string
	}{
		{Point{1, 2}, "(1, 2)"},
		{Point{0.2, 0.6}, "(0.2, 0.6)"},
	} {
		got := c.in.String()
		if got != c.want {
			t.Errorf("Point{%v, %v} prints %v, want %v", c.in.X, c.in.Y, got, c.want)
		}
	}
}
