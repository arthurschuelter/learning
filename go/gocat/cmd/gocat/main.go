package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Error: No file selected")
		os.Exit(1)
	}

	path := args[1]

	f, closer, err := OpenFile(path)
	CheckError(err)

	data := make([]byte, 1)

	for {
		line, err := f.Read(data)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		fmt.Printf("%s", data[:line])
	}

	defer closer()
}

func OpenFile(path string) (*os.File, func(), error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, nil, err
	}

	return file, func() {
		file.Close()
	}, nil
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: no such file or directory")
		os.Exit(1)
	}
}
