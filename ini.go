package main

import (
	"path/filepath"

	"gopkg.in/ini.v1"
)

// This file covers writing, reading etc from and to INI files used in the project

func initializeConfigINI(path string, usr *User) error {
	inidata := ini.Empty()
	sec, err := inidata.NewSection("user")
	if err != nil {
		return err
	}

	_, err = sec.NewKey("first_name", usr.firstName)
	if err != nil {
		return err
	}

	_, err = sec.NewKey("email", usr.email)
	if err != nil {
		return err
	}

	err = inidata.SaveTo(filepath.Join(path, "usrconfig.ini"))
	if err != nil {
		return err
	}

	return nil
}
