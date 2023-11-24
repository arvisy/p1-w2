package channel

import "sync"

func CalculateSumOfSquares(numbers []int) int {
	var wg sync.WaitGroup
	resultChan := make(chan int)

	// Membagi tugas perhitungan ke beberapa Goroutines
	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			result := n * n
			resultChan <- result
		}(num)
	}

	// Menutup channel setelah semua Goroutines selesai
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Mengumpulkan hasil dari channel
	sum := 0
	for res := range resultChan {
		sum += res
	}

	return sum
}
