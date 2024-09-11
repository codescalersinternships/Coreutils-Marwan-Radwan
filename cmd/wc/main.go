package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func countStats(reader io.Reader) (int, int, int, error) {
	scanner := bufio.NewScanner(reader)
	lineCount, wordCount, charCount := 0, 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		lineCount++
		charCount += utf8.RuneCountInString(line)
		wordCount += len(strings.Fields(line))
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, 0, fmt.Errorf("error reading input: %v", err)
	}

	return lineCount, wordCount, charCount, nil
}

func main() {
	var numLines, numWords, numChars bool
	flag.BoolVar(&numLines, "l", false, "print the number of lines")
	flag.BoolVar(&numWords, "w", false, "print the number of words")
	flag.BoolVar(&numChars, "c", false, "print the number of characters")

	flag.Usage = func() {
		fmt.Println("Usage: wc [options] [file]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	var reader io.Reader
	var fileName string

	if flag.NArg() > 0 {
		fileName = flag.Arg(0)
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		defer file.Close()
		reader = file
	} else {
		reader = os.Stdin
	}

	lineCount, wordCount, charCount, err := countStats(reader)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	if !numLines && !numWords && !numChars {
		numLines, numWords, numChars = true, true, true
	}

	if numLines {
		fmt.Printf("%d ", lineCount)
	}
	if numWords {
		fmt.Printf("%d ", wordCount)
	}
	if numChars {
		fmt.Printf("%d ", charCount)
	}
	fmt.Printf("%s\n", fileName)
}
