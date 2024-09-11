package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/codescalersinternships/Coreutils-Marwan-Radwan/internal/cat"
)

func main() {
	var nFlag bool
	flag.BoolVar(&nFlag, "n", false, "to number output lines")

	flag.Usage = func() {
		fmt.Println("Usage: cat [options] [file]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	filePath := ""
	if flag.NArg() > 0 {
		filePath = flag.Arg(0)
	}

	err := cat.Cat(filePath, nFlag)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
