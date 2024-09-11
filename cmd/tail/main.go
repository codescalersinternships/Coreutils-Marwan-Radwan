package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/codescalersinternships/Coreutils-Marwan-Radwan/internal/tail"
)

func main() {
	var numLines uint
	flag.UintVar(&numLines, "n", 10, "print the last n lines instead of the last 10")

	flag.Usage = func() {
		fmt.Println("Usage: tail [options] [file]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	filePath := flag.Arg(0)
	err := tail.Tail(filePath, int(numLines))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
