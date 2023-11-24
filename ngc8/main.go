package main

import (
	"fmt"
	"ngc/ngc8/channel"
	"sync"
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
	// input := []channel.Shape{
	// 	{ShapeType: "RECTANGLE", Length: 5},
	// 	{ShapeType: "CIRCLE", Length: 3},
	// 	{ShapeType: "TRIANGLE", Length: 5},
	// 	{ShapeType: "RECTANGLE", Length: 15},
	// 	{ShapeType: "CIRCLE", Length: 5},
	// }

	// channel.ProcessShapes(input)

	// THEMES 1
	// numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// sum := channel.CalculateSumOfSquares(numbers)

	// fmt.Printf("Sum of squares: %d\n", sum)

	// THEMES 2
	filenames := []string{"file1.txt", "file2.txt", "file3.txt"}
	var wg sync.WaitGroup
	resultChan := make(chan string, len(filenames))

	for _, filename := range filenames {
		wg.Add(1)
		go channel.CountWords(filename, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		fmt.Println(result)
	}
}
