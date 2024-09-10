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
	flag.UintVar(&numLines, "n", 10, "print the first n lines instead of the first 10")

	flag.Usage = func() {
		fmt.Println("Usage: head [options] [file]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	var file *os.File

	if flag.NArg() == 0 {
		file = os.Stdin
	} else {
		filePath := flag.Arg(0)

		file, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("Error happened: %v", err)
		}
		defer file.Close()
	}

	scanner := bufio.NewScanner(file)

	lineCnt := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		lineCnt++
		if lineCnt == int(numLines) {
			break
		}
	}
}
