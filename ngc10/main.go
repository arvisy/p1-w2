package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// SOAL 1
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the first word: ")
	scanner.Scan()
	word1 := scanner.Text()

	fmt.Print("Enter the second word: ")
	scanner.Scan()
	word2 := scanner.Text()

	checkAnagram(word1, word2)

	// SOAL 2
	calculator()
}
