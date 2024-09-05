package main

import "os"

func main() {
	env := os.Environ()
	for _, e := range env {
		println(e)
	}
}
