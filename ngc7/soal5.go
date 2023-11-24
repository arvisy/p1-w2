package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

// downloadFile mengunduh file dari URL yang diberikan dan menyimpannya ke tujuan yang ditentukan.
func downloadFile(url string, destination string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Dapatkan data dari URL
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error saat mengunduh file dari %s: %v\n", url, err)
		return
	}
	defer response.Body.Close()

	// Buat file
	file, err := os.Create(destination)
	if err != nil {
		fmt.Printf("Error saat membuat file %s: %v\n", destination, err)
		return
	}
	defer file.Close()

	// Tulis body ke file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Printf("Error saat menulis ke file %s: %v\n", destination, err)
		return
	}

	fmt.Printf("Diunduh: %s ke %s\n", url, destination)
}
