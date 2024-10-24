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

// Todo: Finalize this
const (
	HELPMESSAGE = `Welcome to Magikit

To initialize magikit use:
	first

To initialize a project use:
	bla 
	bla
	bla

To tweak existing projects use:
	reset

To download/upload projects use:
	cast [url]
	draw [url]
`
	DO_NOT_EDIT_THIS_FOLDER_CONTENTS = "Do not manually change anything in this folder as it may result in file corruption or bugs in the version control."
)

func firstTimeSetup() error {
	usr := User{}

	if _, err := os.Stat(".magikit"); !os.IsNotExist(err) {
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

	err := os.Mkdir(".magikit", 0755)
	if err != nil {
		return err
	}

	err = os.Mkdir(".magikit/usrconfig", 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(".magikit/usrconfig/config", []byte(usr.email), 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(".magikit/DO_NOT_EDIT_THIS_FOLDER_CONTENTS", []byte(DO_NOT_EDIT_THIS_FOLDER_CONTENTS), 0755)
	if err != nil {
		return err
	}

	if runtime.GOOS == "windows" {
		filenameW, err := syscall.UTF16PtrFromString(".magikit")
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
