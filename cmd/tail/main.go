package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	numLines := flag.Int("n", 10, "print the last n lines instead of the last 10")

	flag.Usage = func() {
		fmt.Println("Usage: tail [options] [file]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	filePath := flag.Arg(0)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("No such file or directory")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, line := range lines[len(lines)-*numLines:] {
		fmt.Println(line)
	}
}
