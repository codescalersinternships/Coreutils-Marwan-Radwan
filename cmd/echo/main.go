package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	startIndex := 1

	// -n flag : suppress trailing newline
	omitNewline := flag.Bool("n", false, "do not output the trailing newline")

	// -h flag : print help message
	flag.Usage = func() {
		fmt.Println("Usage: echo [options] [text...]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *omitNewline {
		startIndex = 2
	}

	args := os.Args[startIndex:]
	text := strings.Join(args, " ")

	if *omitNewline {
		fmt.Print(text)
	} else {
		fmt.Println(text)
	}
}
