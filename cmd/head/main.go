package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	filePath := os.Args[len(os.Args)-1]

	// -n flag : to specify the number of lines to print
	numLines := flag.Int("n", 10, "print the first n lines instead of the first 10")

	// -h flag : print help message
	flag.Usage = func() {
		fmt.Println("Usage: head [options] [file]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("No such file or directory")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineCnt := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		lineCnt++
		if lineCnt == *numLines {
			break
		}
	}
}
