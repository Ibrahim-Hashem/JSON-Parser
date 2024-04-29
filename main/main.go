package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}

}

// read file from the command line
func read_file() {
	
}

func main() {
	// read fil path argument from the command line
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path")
		os.Exit(1)
	}

	// read file path from the command line
	file_path := os.Args[1]
	
	println(file_path)
}