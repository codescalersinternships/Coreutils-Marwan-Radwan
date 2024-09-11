package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/codescalersinternships/Coreutils-Marwan-Radwan/internal/head"
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

	filePath := ""
	if flag.NArg() > 0 {
		filePath = flag.Arg(0)
	}

	err := head.Head(filePath, int(numLines))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
