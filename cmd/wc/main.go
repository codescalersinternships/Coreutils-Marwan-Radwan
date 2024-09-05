package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	filePath := os.Args[len(os.Args)-1]

	// -l flag : print the number of lines
	numLines := flag.Bool("l", false, "print the number of lines")
	// -w flag : print the number of words
	numWords := flag.Bool("w", false, "print the number of words")
	// -c flag : print the number of characters
	numChars := flag.Bool("c", false, "print the number of characters")

	// -h flag : print help message
	flag.Usage = func() {
		fmt.Println("Usage: wc [options] [file]")
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

	linesCnt := 0
	wordsCnt := 0
	charsCnt := 0

	for scanner.Scan() {
		linesCnt++
		charsCnt += len(scanner.Text())
		wordsCnt += len(strings.Split(scanner.Text(), " "))
	}

	if *numLines {
		fmt.Printf("%d ", linesCnt)
	}
	if *numWords {
		fmt.Printf("%d ", wordsCnt)
	}
	if *numChars {
		fmt.Printf("%d ", charsCnt)
	}
	if !*numLines && !*numWords && !*numChars {
		fmt.Printf("%d %d %d", linesCnt, wordsCnt, charsCnt)
	}

	fmt.Printf("%s\n", filePath)
}
