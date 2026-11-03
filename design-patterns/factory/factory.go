package main

import "fmt"

func MakeComputer(computerType string) (IComputer, error) {
	switch computerType {
	case "MacOS":
		return MakeMacOS(), nil
	case "Windows":
		return MakeWindows(), nil
	case "Linux":
		return MakeLinux(), nil
	default:
		return nil, fmt.Errorf("Computer type does not exist")
	}
}
