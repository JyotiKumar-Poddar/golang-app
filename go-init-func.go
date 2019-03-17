package main

import (
	"fmt"
	"runtime"
)

var message string

// program entry point required only for executables
func main() {
	message = "inside main function\n"
	print(message)
	fmt.Printf("Running in OS %v\n", runtime.GOOS)
}

// init function executes before function main()
func init() {
	message = "Inside init\n"
	print("Starts before method\n")
	print(message)
}

//TODO skip unwanted  Function return value
