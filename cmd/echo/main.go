package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/codescalersinternships/Coreutils-Marwan-Radwan/internal/echo"
)

func main() {
	var omitNewline bool
	flag.BoolVar(&omitNewline, "n", false, "do not output the trailing newline")

	flag.Usage = func() {
		fmt.Println("Usage: echo [options] [text...]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	args := flag.Args()
	text := strings.Join(args, " ")

	echo.Echo(text, omitNewline)
}
