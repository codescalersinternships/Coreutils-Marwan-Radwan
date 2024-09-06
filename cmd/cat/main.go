package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	nFlag := flag.Bool("n", false, "to number output lines")

	flag.Usage = func() {
		fmt.Println("Usage: cat [options] [file]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	filePath := flag.Arg(0)

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
