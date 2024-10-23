package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"syscall"
)

type User struct {
	firstName string
	email     string
}

func firstTimeSetup() error {
	usr := User{}

	if _, err := os.Stat(".magiit"); !os.IsNotExist(err) {
		fmt.Println("Setup has already been run. If you wish to make changes locate you .magiit file and open on a text editor or run the updatesetup command")
		return nil
	}

	fmt.Print("First name: ")
	r := bufio.NewReader(os.Stdin)
	n, err := r.ReadString('\n')

	if err != nil {
		fmt.Print("Error reading input")
		return err
	}

	fmt.Print("Email: ")
	r = bufio.NewReader(os.Stdin)
	e, err := r.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input")
		return err
	}

	usr.firstName = n
	usr.email = e

	err = generateUsrConfig(&usr)
	if err != nil {
		fmt.Println("Error creating user file:")
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func generateUsrConfig(usr *User) error {
	err := os.WriteFile(".magiit", []byte(usr.email), 0755)

	if err != nil {
		return err
	}

	if runtime.GOOS == "windows" {
		filenameW, err := syscall.UTF16PtrFromString(".magiit")
		if err != nil {
			return err
		}

		err = syscall.SetFileAttributes(filenameW, syscall.FILE_ATTRIBUTE_HIDDEN)
		if err != nil {
			return err
		}
	}

	return nil
}
