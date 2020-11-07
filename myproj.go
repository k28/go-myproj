package myproj

import "fmt"

// Sum a, b
func Sum(a int, b int) int {
	return a + b
}

// Hello say hello
func Hello(target string) string {
	return fmt.Sprintf("Hello %s", target)
}
