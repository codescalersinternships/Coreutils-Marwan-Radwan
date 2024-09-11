package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/codescalersinternships/Coreutils-Marwan-Radwan/internal/tree"
)

func main() {
	var depth uint
	flag.UintVar(&depth, "l", 2, "specify the depth level")

	flag.Usage = func() {
		fmt.Println("Usage: tree [options] [file]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	path := flag.Arg(0)
	if path == "" {
		path = "."
	}

	err := tree.Tree(path, int(depth))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
