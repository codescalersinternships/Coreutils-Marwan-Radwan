package cat

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Cat reads from the given file path and prints the content to the standard output.
func Cat(filePath string, numberLines bool) error {
	var reader io.Reader

	if filePath == "" {
		reader = os.Stdin
	} else {
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
		reader = file
	}

	scanner := bufio.NewScanner(reader)
	lineCount := 1

	for scanner.Scan() {
		if numberLines {
			fmt.Printf("%d ", lineCount)
		}
		fmt.Println(scanner.Text())
		lineCount++
	}

	return scanner.Err()
}
