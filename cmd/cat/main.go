package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	filePath := os.Args[len(os.Args)-1]

	// -n flag : to number the output lines
	nFlag := flag.Bool("n", false, "number of lines to print")

	// -h flag : print help message
	flag.Usage = func() {
		fmt.Println("Usage: cat [options] [file]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("No such file or directory")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineCnt := 1
	for scanner.Scan() {
		if *nFlag {
			fmt.Printf("%d ", lineCnt)
		}
		fmt.Println(scanner.Text())
		lineCnt++
	}

}
