package echo

import (
	"fmt"
)

// Echo prints the given text to the standard output.
func Echo(text string, newLineFlag bool) {
	if newLineFlag {
		fmt.Print(text)
	} else {
		fmt.Println(text)
	}
}
