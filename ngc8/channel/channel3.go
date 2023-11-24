package channel

import (
	"fmt"
	"math"
	"sync"
)

const (
	RECTANGLE = "RECTANGLE"
	CIRCLE    = "CIRCLE"
	TRIANGLE  = "TRIANGLE"
)

type Shape struct {
	ShapeType string // Circle/Rectangle
	Length    float32
	Area      float32
}

func CalculateArea(shape *Shape, resultChan chan<- Shape, wg *sync.WaitGroup) {
	defer wg.Done()
	switch shape.ShapeType {
	case RECTANGLE:
		shape.Area = shape.Length * shape.Length
	case CIRCLE:
		shape.Area = math.Pi * shape.Length * shape.Length
	case TRIANGLE:
		shape.Area = 0.5 * shape.Length * shape.Length
	}
	resultChan <- *shape
}

func ProcessShapes(input []Shape) {
	resultChan := make(chan Shape, len(input)) // Menggunakan buffered channel

	var wg sync.WaitGroup

	for _, shape := range input {
		wg.Add(1)
		switch shape.ShapeType {
		case RECTANGLE:
			go CalculateArea(&shape, resultChan, &wg)
		case CIRCLE:
			go CalculateArea(&shape, resultChan, &wg)
		case TRIANGLE:
			go CalculateArea(&shape, resultChan, &wg)
		}
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for s := range resultChan {
		fmt.Printf("Shape: %s, Area: %f\n", s.ShapeType, s.Area)
	}
}
