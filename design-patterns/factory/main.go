package main

import "fmt"

func main() {
	fmt.Println("=== Computer Factory ===")

	macOs, _ := MakeComputer(MacOS)
	macOs.print()

	windows, _ := MakeComputer(Windows)
	windows.print()

	linux, _ := MakeComputer(Linux)
	linux.print()
}
