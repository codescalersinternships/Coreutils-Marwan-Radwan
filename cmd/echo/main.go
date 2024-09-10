package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var omitNewline bool
	flag.BoolVar(&omitNewline, "n", false, "do not output the trailing newline")

	flag.Usage = func() {
		fmt.Println("Usage: echo [options] [text...]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	args := flag.Args()
	text := strings.Join(args, " ")

	if omitNewline {
		fmt.Print(text)
	} else {
		fmt.Println(text)
	}
}
