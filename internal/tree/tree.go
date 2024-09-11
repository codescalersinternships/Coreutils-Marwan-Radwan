package tree

import (
	"fmt"
	"os"
	"path/filepath"
)

// Tree traverses the directory tree rooted at the given path and prints the number of directories and files.
func Tree(path string, maxDepth int) error {
	numFiles := 0
	numDirs := 0

	if err := treeTraverse(path, maxDepth, 0, &numDirs, &numFiles); err != nil {
		return err
	}

	fmt.Printf("%d directories, %d files\n", numDirs, numFiles)

	return nil
}

func treeTraverse(path string, maxDepth int, curDepth int, numDirs *int, numFiles *int) error {
	if curDepth > maxDepth {
		return nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, f := range files {
		for i := 1; i < curDepth; i++ {
			fmt.Print("│   ")
		}
		if f.IsDir() {
			*numDirs++
			fmt.Println("├──", f.Name())
			if err := treeTraverse(filepath.Join(path, f.Name()), curDepth+1, maxDepth, numDirs, numFiles); err != nil {
				return err
			}
		} else {
			*numFiles++
			fmt.Println("└──", f.Name())
		}
	}

	return nil
}
