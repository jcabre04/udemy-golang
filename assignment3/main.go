package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	// Validate the command format is correct
	if len(os.Args) != 2 {
		fmt.Println("Incorrect cli command format. Correct ex:\n"+
			"\tgo run main.go myfile.txt")
		os.Exit(1)
	}

	// Open the file
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(2)
	}

	io.Copy(os.Stdout, f)
}