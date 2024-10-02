package main

import (
	"flag"
	"fmt"
	"os"
)

func reverse(s string) string {
	buffer := make([]byte, 0, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		buffer = append(buffer, s[i])
	}
	return string(buffer)
}

func main() {
	var inputString string
	

	printInput := flag.Bool("source", false, "print original string")
	flag.Parse()


	if len(flag.Args()) < 1 || flag.Args()[0] == "" {
		fmt.Println("Please provide a non-empty string. Usage: <command> <string> [--source]")
		os.Exit(1)
	}
	fmt.Println("Input is provided correctly. Processing...")
	
	inputString = flag.Args()[0]

	reversed := reverse(inputString)
	fmt.Println("Reversed string:", reversed)

	if *printInput {
		fmt.Println("Original string:", inputString)
	}

	fmt.Println("Program has exited successfully")
}
