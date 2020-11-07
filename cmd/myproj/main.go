package main

import (
	"os"

	myproj "github.com/k28/go-myproj"
)

func main() {
	os.Exit(Run(os.Args))
}

// Run main
func Run(args []string) int {
	target := parseArgs()
	myproj.Hello(target)
	return 0
}
