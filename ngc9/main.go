package main

import (
	"fmt"
	"sync"
)

func main() {
	// DNE 1
	// randomNumberChannel := make(chan int)

	// go func() {
	// 	randomNumber := 4
	// 	randomNumberChannel <- randomNumber
	// 	close(randomNumberChannel)
	// }()

	// receivedNumber, factorialResultChannel := Factorial(randomNumberChannel)

	// fmt.Println("Received random number: ", receivedNumber)

	// factorialResult := <-factorialResultChannel
	// fmt.Printf("Factorial of %d is: %d\n", receivedNumber, factorialResult)

	// DNE 2
	diameters := []float64{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	resultChan := make(chan Circle, len(diameters))

	for _, diameter := range diameters {
		wg.Add(1)
		go calculateCircleProperties(diameter, resultChan, &wg)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		fmt.Printf("Diameter: %.2f\n", result.Area)
		fmt.Printf("Luas: %.2f\n", result.Area)
		fmt.Printf("Keliling: %.2f\n", result.Circumference)
		fmt.Println("==========================")
	}
}
