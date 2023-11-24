package main

import (
	"fmt"
)

func main() {
	randomNumberChannel := make(chan int)

	go func() {
		randomNumber := 4
		randomNumberChannel <- randomNumber
		close(randomNumberChannel)
	}()

	receivedNumber, factorialResultChannel := Factorial(randomNumberChannel)

	fmt.Println("Received random number: ", receivedNumber)

	factorialResult := <-factorialResultChannel
	fmt.Printf("Factorial of %d is: %d\n", receivedNumber, factorialResult)
}
