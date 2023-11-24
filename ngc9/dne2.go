package main

import (
	"math"
	"sync"
)

type Circle struct {
	Area          float64
	Circumference float64
}

func calculateCircleProperties(diameter float64, resultChan chan<- Circle, wg *sync.WaitGroup) {
	defer wg.Done()

	radius := diameter / 2
	area := math.Pi * radius * radius
	circumference := math.Pi * diameter

	result := Circle{
		Area:          area,
		Circumference: circumference,
	}

	resultChan <- result
}
