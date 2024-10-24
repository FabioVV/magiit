package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 || len(os.Args) < 2 {
		fmt.Println("Incorrect usage. Type help to see available commands")
		return
	}

	switch os.Args[1] {
	case "first":
		fmt.Println("First time setup ...")
		err := firstTimeSetup()

		if err != nil {
			return
		}

	case "updatesetup":

	case "help":
		fmt.Println(HELPMESSAGE)
	default:
		fmt.Println("Incorrect usage: Command not recognised. Type help to see available commands ")
	}

}
