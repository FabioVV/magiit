package main

import (
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"
)

func HashFile(filecontent []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(filecontent))
}

func addFile(file string) error {
	filecontent, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	hash := HashFile(filecontent)

	filePathToAdd := filepath.Join(ROOT, "objs", hash)

	err = os.WriteFile(filePathToAdd, filecontent, 0755)
	if err != nil {
		return err
	}

	fmt.Println("Added", file, "to the staging area")
	return nil
}
