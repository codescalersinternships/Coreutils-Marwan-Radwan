package main

import (
	"flag"
	"fmt"
	"os"
)

func tree(path string, curDepth int, maxDepth int, numDirs *int, numFiles *int) {
	if curDepth > maxDepth {
		return
	}

	file, err := os.ReadDir(path)
	if err != nil {
		return
	}

	for _, f := range file {
		for i := 1; i < curDepth; i++ {
			fmt.Print("│   ")
		}
		if f.IsDir() {
			*numDirs++
			fmt.Println("├──", f.Name())
			tree(path+"/"+f.Name(), curDepth+1, maxDepth, numDirs, numFiles)
		} else {
			*numFiles++
			fmt.Println("└──", f.Name())
		}
	}

}

func main() {
	depth := flag.Int("l", 2, "specify the depth level")

	flag.Usage = func() {
		fmt.Println("Usage: tree [options] [file]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	path := flag.Arg(0)

	numFiles := 0
	numDirs := 0
	tree(path, 0, *depth, &numDirs, &numFiles)

	fmt.Printf("%d directories, %d files\n", numDirs, numFiles)
}
