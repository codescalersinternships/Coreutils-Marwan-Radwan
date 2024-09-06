package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	numLines := flag.Int("n", 10, "print the first n lines instead of the first 10")

	flag.Usage = func() {
		fmt.Println("Usage: head [options] [file]")
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

	lineCnt := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		lineCnt++
		if lineCnt == *numLines {
			break
		}
	}
}
