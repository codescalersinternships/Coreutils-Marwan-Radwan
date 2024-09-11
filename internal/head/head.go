package head

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Head reads the contents of a file or standard input and prints the specified number of lines.
func Head(filePath string, numLines int) error {
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

	for i := 0; i < numLines && scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}

	return scanner.Err()
}
