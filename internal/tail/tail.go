package tail

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Tail reads the last n lines from a file and prints them to the standard output.
func Tail(filePath string, numLines int) error {
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

	return nil
}
