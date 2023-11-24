package main

import (
	"fmt"
	"sync"
)

func main() {
	// SOAL 1
	// var wg sync.WaitGroup
	// ch := make(chan int, 2)

	// for i := 1; i <= 10; i++ {
	// 	wg.Add(1)
	// 	go Fibonacci(i, &wg, ch)
	// 	wg.Wait()
	// 	result := <-ch
	// 	fmt.Printf("bilangan fibonacci ke-%d adalah %d\n", i, result)
	// }
	// close(ch)

	// // SOAL 2
	// istilahs := []Istilah{
	// 	{Nama: "As Soon As Possible"},
	// 	{Nama: "Liquid-crystal display"},
	// 	{Nama: "Thank George It's Friday!"},
	// }

	// Akronim(istilahs)

	// SOAL 3
	// words := []string{"exampleword", "scrabble", "goroutine", "asynchronous"}

	// var wg sync.WaitGroup

	// for _, word := range words {
	// 	wg.Add(1)
	// 	go func(w string) {
	// 		defer wg.Done()
	// 		score := CalculateScrabbleScore(w)
	// 		fmt.Printf("%s | scrabble score %d\n", w, score)
	// 	}(word)
	// }

	// wg.Wait()

	// SOAL 4
	// var wg sync.WaitGroup

	// Task 1
	// wg.Add(1)
	// ScheduleTask(func() {
	// 	defer wg.Done()
	// 	fmt.Println("Task 1 is being executed")
	// 	time.Sleep(2 * time.Second)
	// })

	// Task 2
	// wg.Add(1)
	// ScheduleTask(func() {
	// 	defer wg.Done()
	// 	fmt.Println("Task 2 is being executed")
	// 	time.Sleep(1 * time.Second)
	// })

	// Task 3
	// wg.Add(1)
	// ScheduleTask(func() {
	// 	defer wg.Done()
	// 	fmt.Println("Task 3 is being executed")
	// 	time.Sleep(3 * time.Second)
	// })

	// fmt.Println("Main application continues...")

	// wg.Wait()

	// fmt.Println("All tasks completed")

	// SOAL 5
	// Daftar URL untuk diunduh
	urls := []string{
		"https://sample-videos.com/video123/mp4/720/big_buck_bunny_720p_1mb.mp4",
		"https://sample-videos.com/video123/mp4/720/big_buck_bunny_720p_1mb.mp4",
		"https://sample-videos.com/video123/mp4/720/big_buck_bunny_720p_1mb.mp4",
	}

	// Buat WaitGroup untuk menunggu semua unduhan selesai
	var wg sync.WaitGroup

	// Iterasi melalui URL dan unduh setiap file secara bersamaan
	for i, url := range urls {
		wg.Add(1)
		go downloadFile(url, fmt.Sprintf("video%d.mp4", i+1), &wg)
	}

	// Tunggu semua Goroutines selesai
	wg.Wait()

	// Cetak pesan yang menunjukkan bahwa semua file telah diunduh
	fmt.Println("Semua file telah diunduh")
}
