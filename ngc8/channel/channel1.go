package channel

import (
	"fmt"
	"time"
)

func Channel1() {
	numChannel := make(chan int)

	go func() {
		defer close(numChannel)
		for i := 0; i < 100; i++ {
			numChannel <- i
		}
	}()

	go func() {
		for number := range numChannel {
			result := fizzBuzz(number)
			fmt.Println(result)
		}
	}()

	time.Sleep(2 * time.Second)
}

func fizzBuzz(number int) string {
	result := ""

	if number%15 == 0 {
		result = fmt.Sprintf("%dFizzBuzz", number)
	} else if number%3 == 0 {
		result = fmt.Sprintf("%dFizz", number)
	} else if number%5 == 0 {
		result = fmt.Sprintf("%dBuzz", number)
	} else {
		result = fmt.Sprintf("%d", number)
	}

	return result
}
