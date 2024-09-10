package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
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

	lineCnt := 1
	for scanner.Scan() {
		if nFlag {
			fmt.Printf("%d ", lineCnt)
		}
		fmt.Println(scanner.Text())
		lineCnt++
	}

}
