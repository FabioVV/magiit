package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 || len(os.Args) < 2 {
		fmt.Println("Incorrect usage")
		return
	}

	switch os.Args[1] {
	case "first":
		fmt.Println("First time setup....")
		firstTimeSetup()

	default:
		fmt.Println("Incorrect usage. Command not recognised")
	}

}
