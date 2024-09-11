package false

import "os"

// False exits the program with a status code of 1.
func False() {
	os.Exit(1)
}
