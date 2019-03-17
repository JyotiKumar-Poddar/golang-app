package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	osName := runtime.GOOS
	osAddress := &osName
	fmt.Println("The current time is", osName, "\nMemory address is", osAddress)
	fmt.Println("Environment variable is", os.Environ())
	changeOsName(&osName) // pass by reference
	fmt.Print("Changed os name ", osName)
}

func changeOsName(osName *string) {
	*osName = "Mac OS"
}
