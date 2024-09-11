package env

import (
	"fmt"
	"os"
)

// Env prints the environment variables to the standard output.
func Env() {
	env := os.Environ()
	for _, e := range env {
		fmt.Println(e)
	}
}
