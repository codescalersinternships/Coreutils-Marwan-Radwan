package main

import (
	"flag"
	"fmt"
)

func main() {
	omitNewline := flag.Bool("n", false, "do not output the trailing newline")

	flag.Usage = func() {
		fmt.Println("Usage: echo [options] [text...]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	text := flag.Arg(0)

	if *omitNewline {
		fmt.Print(text)
	} else {
		fmt.Println(text)
	}
}
