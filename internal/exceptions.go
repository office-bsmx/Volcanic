package internal

import (
	"log"
	"os"
	"strconv"
)

// Throws an exception with specified code and terminates a program
func Throw(Code int) {
	log.Fatal("Exception raised with code " + strconv.Itoa(Code) + ", terminating...")
	os.Exit(Code)
}
