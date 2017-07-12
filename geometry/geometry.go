// Package geometry implements geometry functions
package geometry

type Shape interface {
	Transform(float64, float64)
	Rotate(Point)
}
