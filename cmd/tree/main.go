package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func tree(path string, curDepth int, maxDepth int, numDirs *int, numFiles *int) error {
	if curDepth > maxDepth {
		return nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("error happened: %v", err)
	}

	for _, f := range files {
		for i := 1; i < curDepth; i++ {
			fmt.Print("│   ")
		}
		if f.IsDir() {
			*numDirs++
			fmt.Println("├──", f.Name())
			if err := tree(filepath.Join(path, f.Name()), curDepth+1, maxDepth, numDirs, numFiles); err != nil {
				return err
			}
		} else {
			*numFiles++
			fmt.Println("└──", f.Name())
		}
	}

	return nil
}

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

	numFiles := 0
	numDirs := 0
	err := tree(path, 0, int(depth), &numDirs, &numFiles)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%d directories, %d files\n", numDirs, numFiles)
}
