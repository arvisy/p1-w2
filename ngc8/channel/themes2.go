package channel

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func CountWords(filename string, wg *sync.WaitGroup, resultChan chan<- string) {
	defer wg.Done()

	file, err := os.Open(filename)
	if err != nil {
		resultChan <- fmt.Sprintf("Error opening file %s: %v", filename, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	wordCount := 0

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		wordCount += len(words)
	}

	if err := scanner.Err(); err != nil {
		resultChan <- fmt.Sprintf("Error reading file %s: %v", filename, err)
		return
	}

	resultChan <- fmt.Sprintf("%s : %d words", filename, wordCount)
}
