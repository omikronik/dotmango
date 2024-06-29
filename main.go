package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type Config struct {
	Configuration []ConfigSettings `json:"files"`
}

type ConfigSettings struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Source string `json:"source"`
	Target string `json:"target"`
}

func main() {
	Start()
}

func Start() {
	var configData Config
	// get config
	configData, err := GetConfig()
	if err != nil {
		fmt.Printf("err in reading config: %s\n", err)
		return
	}

	for _, element := range configData.Configuration {
		fmt.Printf("%+v\n", element)
		fullTargetPath, err := ExpandPath(element.Target)
		if err != nil {
			fmt.Printf("err in expand target path, target value: %s,\nerr: %s\n", element.Target, err)
			return
		}
		var res bool
		res, err = PathExists(fullTargetPath)
		if err != nil {
			fmt.Printf("err in pathexists: %s\n", err)
		}
		if res == false {
		}
	}

	var res bool
	res, err = PathExists("~/.config/alacritty2")
	if err != nil {
		fmt.Printf("err in pathexists: %s\n", err)
	}
	fmt.Printf("first exist result: %t\n", res)

	res, err = PathExists("~/.config/alacritty")
	if err != nil {
		fmt.Printf("err in pathexists: %s\n", err)
	}
	fmt.Printf("second exist result: %t\n", res)
}

func GetConfig() (Config, error) {
	var configData Config

	f, err := os.ReadFile("./dotmango.json")
	if err != nil {
		return configData, err
	}

	err = json.Unmarshal(f, &configData)
	if err != nil {
		return configData, err
	}

	return configData, nil

}

func DeleteIfExists(source string) (bool, error) {
	return false, nil
}

func CreateFilepath(dest string) error {
	err := os.MkdirAll(dest, os.FileMode(755))
	if err != nil {
		return err
	}
	return nil
}

func Symlinkify(source string, dest string) error {
	err := os.Symlink(source, dest)
	if err != nil {
		return err
	}

	return nil
}

func PathExists(pathName string) (bool, error) {
	path, pathErr := ExpandPath(pathName)
	if pathErr != nil {
		return false, pathErr
	}

	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}

	return false, err
}

func ExpandPath(pathName string) (string, error) {
	if strings.HasPrefix(pathName, "~") {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}
		pathName = filepath.Join(usr.HomeDir, pathName[1:])
	}
	return pathName, nil
}
