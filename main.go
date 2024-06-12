package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Printf("Hello bruh\n")
	res, err := pathExists("~/.config/alacritty2")
	if err != nil {
		fmt.Printf("err in pathexists: %s\n", err)
	}
	fmt.Printf("first exist result: %t\n", res)

	res, err = pathExists("~/.config/alacritty")
	if err != nil {
		fmt.Printf("err in pathexists: %s\n", err)
	}
	fmt.Printf("second exist result: %t\n", res)
}

func pathExists(pathName string) (bool, error) {
	path, pathErr := expandPath(pathName)
	if pathErr != nil {
		return false, pathErr
	}
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func expandPath(pathName string) (string, error) {
	if strings.HasPrefix(pathName, "~") {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}
		pathName = filepath.Join(usr.HomeDir, pathName[1:])
	}
	return pathName, nil
}
