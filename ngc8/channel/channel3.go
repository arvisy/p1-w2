package channel

import "math"

const (
	RECTANGLE = "rectangle"
	CIRCLE    = "circle"
	TRIANGLE  = "triangle"
)

type Shape struct {
	ShapeType string
	Length    int
	Area      float32
}

func CalculateArea(ch chan Shape) {
	for shape := range ch {
		switch shape.ShapeType {
		case RECTANGLE:
			shape.Area = float32(shape.Length) * float32(shape.Length)
		case CIRCLE:
			shape.Area = math.Pi * float32(shape.Length) * float32(shape.Length)
		case TRIANGLE:
			shape.Area = 0.5 * float32(shape.Length) * float32(shape.Length)
		}
		ch <- shape
	}
}
