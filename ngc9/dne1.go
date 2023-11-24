package main

func Factorial(ch <-chan int) (int, chan int) {
	n := <-ch

	result := make(chan int)
	go func() {
		defer close(result)
		factorial := 1
		for i := 1; i <= n; i++ {
			factorial *= i
		}
		result <- factorial
	}()
	return n, result
}
