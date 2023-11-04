package internal

import (
	"log"
)

// Throws an exception with specified code and terminates a program
func Throw(Code int, Err ...error) {
	log.Fatal("Exception raised with code ", Code, ", terminating...\n", Err)
}
