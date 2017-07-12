package geometry

import (
	"testing"
)

func TestLineString(t *testing.T) {
	for _, c := range []struct {
		in   Line
		want string
	}{
		{Line{Point{1, 2}, Point{3, 4}}, "((1, 2), (3, 4))"},
	} {
		got := c.in.String()
		if got != c.want {
			t.Errorf("Line{%v, %v} prints %v, want %v", c.in.A, c.in.B, got, c.want)
		}
	}
}
