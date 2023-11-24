package main

import (
	"fmt"
	"ngc8/channel"
)

func main() {
	// CHANNEL 1
	// channel.Channel1()

	// CHANNEL 2
	// sumSquareResult := make(chan int)
	// squareSumResult := make(chan int)

	// go channel.SumSquare(100, sumSquareResult)
	// go channel.SquareSum(100, squareSumResult)

	// sumSquare := <-sumSquareResult
	// squareSum := <-squareSumResult

	// fmt.Printf("Perbedaan kalkulasi:\n(a) Jumlah kuadrat: %d\n(b) Kuadrat jumlah: %d\n", sumSquare, squareSum)

	// CHANNEL 3
	rectChannel := make(chan channel.Shape)
	circleChannel := make(chan channel.Shape)
	triangleChannel := make(chan channel.Shape)

	input := []channel.Shape{
		{ShapeType: channel.RECTANGLE, Length: 5},
		{ShapeType: channel.CIRCLE, Length: 3},
		{ShapeType: channel.TRIANGLE, Length: 5},
		{ShapeType: channel.RECTANGLE, Length: 15},
		{ShapeType: channel.CIRCLE, Length: 5},
	}

	go func() {
		for _, shape := range input {
			switch shape.ShapeType {
			case channel.RECTANGLE:
				rectChannel <- shape
			case channel.CIRCLE:
				rectChannel <- shape
			case channel.TRIANGLE:
				rectChannel <- shape
			}
		}

		close(rectChannel)
		close(circleChannel)
		close(triangleChannel)
	}()

	go channel.CalculateArea(rectChannel)
	go channel.CalculateArea(circleChannel)
	go channel.CalculateArea(triangleChannel)

	for i := 0; i < len(input); i++ {
		select {
		case rect := <-rectChannel:
			fmt.Printf("Rectangle with length %d has area: %f\n", rect.Length, rect.Area)
		case circle := <-circleChannel:
			fmt.Printf("Circle with radius %d has area: %f\n", circle.Length, circle.Area)
		case triangle := <-triangleChannel:
			fmt.Printf("Triangle with base %d has area: %f\n", triangle.Length, triangle.Area)
		}
	}

}
