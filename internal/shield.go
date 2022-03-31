package internal

import "fmt"

func printString() string {
	return "Test shields.io"
}

func PrintToStd() {
	fmt.Println(printString())
}
