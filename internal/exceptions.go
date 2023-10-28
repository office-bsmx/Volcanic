package internal

import (
	"fmt"
	"os"
	"strconv"
)

// Throws an exception with specified code and terminates a program
func Throw(Code int) {
	fmt.Println("Exception raised with code " + strconv.Itoa(Code) + ", terminating...")
	os.Exit(Code)
}
