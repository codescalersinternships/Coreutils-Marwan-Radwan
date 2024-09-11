package wc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Count counts the number of lines, words, and characters in a file and prints them in Standard-Output.
func Count(filePath string, numLines bool, numWords bool, numChars bool) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineCount, wordCount, charCount := 0, 0, 0

	for scanner.Scan() {
		lineCount++
		charCount += len(scanner.Text())
		wordCount += len(strings.Fields(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return err
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
	if !numLines && !numWords && !numChars {
		fmt.Printf("%d %d %d", lineCount, wordCount, charCount)
	}

	fmt.Printf("%s\n", filePath)

	return nil
}
