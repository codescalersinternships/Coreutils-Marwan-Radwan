package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/codescalersinternships/Coreutils-Marwan-Radwan/internal/wc"
)

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

	filePath := flag.Arg(0)

	err := wc.Count(filePath, numLines, numWords, numChars)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

}
