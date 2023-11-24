package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// SOAL 1
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Error:", r)
	// 	}
	// }()

	// scanner := bufio.NewScanner(os.Stdin)

	// fmt.Print("Enter the first word: ")
	// scanner.Scan()
	// word1 := scanner.Text()

	// fmt.Print("Enter the second word: ")
	// scanner.Scan()
	// word2 := scanner.Text()

	// checkAnagram(word1, word2)

	// SOAL 2
	// calculator()

	// THEMES
	// Parse command-line flags
	filePath := flag.String("file", "", "Path to the log file")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Please provide the path to the log file using the -file flag.")
		return
	}

	logAnalyzer := NewLogAnalyzer()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic occurred:", r)
		}
	}()

	logAnalyzer.AnalyzeLogs(*filePath)
	logAnalyzer.DisplayResults()

}

// LogEntry adalah struktur yang merepresentasikan satu baris log dengan levelnya.
type LogEntry struct {
	Level string
}

// LogAnalyzer adalah alat analisis log yang dapat menghitung jumlah kemunculan setiap level log.
type LogAnalyzer struct {
	LogLevelCounts map[string]int
}

// NewLogAnalyzer membuat sebuah instansi baru dari LogAnalyzer dengan map LogLevelCounts yang sudah diinisialisasi.
func NewLogAnalyzer() *LogAnalyzer {
	return &LogAnalyzer{
		LogLevelCounts: make(map[string]int),
	}
}

// AnalyzeLogs membaca dan menganalisis file log. Setiap baris log diekstrak levelnya, dan jumlahnya dihitung.
func (la *LogAnalyzer) AnalyzeLogs(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		level := extractLogLevel(line)
		la.LogLevelCounts[level]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
}

// DisplayResults menampilkan hasil analisis log, mencetak jumlah kemunculan setiap level log.
func (la *LogAnalyzer) DisplayResults() {
	fmt.Println("Log Analysis Results:")
	for level, count := range la.LogLevelCounts {
		fmt.Printf("[%s] Level: %d occurrences\n", level, count)
	}
}

// extractLogLevel mengambil level log dari sebuah baris log menggunakan ekspresi regular.
func extractLogLevel(line string) string {
	re := regexp.MustCompile(`\[(.*?)\]`)
	matches := re.FindStringSubmatch(line)
	if len(matches) > 1 {
		return strings.ToUpper(matches[1])
	}
	return "UNKNOWN"
}
