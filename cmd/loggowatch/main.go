package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/NikitaCherkashin4/loggowatch/analyzer"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <logfile>")
		os.Exit(1)
	}

	filename := os.Args[1]
	lines, err := readLogFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	stats := analyzer.AnalyzeLogs(lines)

	fmt.Printf("Log Analysis Results:\n")
	fmt.Printf("Total Lines: %d\n", stats.TotalLines)
	fmt.Printf("\nLog Level Counts:\n")
	for level, count := range stats.Counts {
		fmt.Printf("  %s: %d\n", level, count)
	}
}

func readLogFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
