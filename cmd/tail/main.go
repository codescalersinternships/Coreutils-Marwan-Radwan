package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var numLines uint
	flag.UintVar(&numLines, "n", 10, "print the last n lines instead of the last 10")

	flag.Usage = func() {
		fmt.Println("Usage: tail [options] [file]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	filePath := flag.Arg(0)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error happened: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var startLineIndex uint
	totalLines := uint(len(lines))
	if totalLines < uint(numLines) {
		startLineIndex = 0
	} else {
		startLineIndex = totalLines - uint(numLines)
	}

	for _, line := range lines[startLineIndex:] {
		fmt.Println(line)
	}
}
