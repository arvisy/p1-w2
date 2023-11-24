package main

import (
	"sync"
	"time"
)

func Fibonacci(n int, wg *sync.WaitGroup, ch chan int) {
	time.Sleep(100 * time.Millisecond)
	defer wg.Done()

	if n <= 1 {
		ch <- n
		return
	}

	wg.Add(2)
	go Fibonacci(n-1, wg, ch)
	go Fibonacci(n-2, wg, ch)

	result1 := <-ch
	result2 := <-ch
	ch <- result1 + result2

}
