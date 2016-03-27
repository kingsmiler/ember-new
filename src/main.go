package main

import (
	"fmt"
	"os"
	"strings"
	"io/ioutil"
)

const fileList []string  = []string{
	"bower.json",
	"app/index.html",
	"README.md",
	"package.json",
	"environment.js",
	"tests/index.html",
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:\t ember-new [name]")
		os.Exit(1)
	}

	emberNewHome := getParentDirectory(getCurrentDirectory())
	userWd, _ := os.Getwd()
	newProjectName := os.Args[1];

	fmt.Println("emberNewHome=" + emberNewHome)

	fmt.Println("userWd=" + userWd)

	copyDir(emberNewHome+"/ember-blueprint", userWd+"/"+newProjectName)
}

func modifyFlag(targetFile string, newProjectName string) {
	userWd, _ := os.Getwd()
	read, err := ioutil.ReadFile(userWd + "/aa.txt")
	if err != nil {
		panic(err)
	}

	newContents := strings.Replace(string(read), "old", "new", -1)

	fmt.Println(newContents)

	err = ioutil.WriteFile(userWd + "/aa.txt", []byte(newContents), 0)
	if err != nil {
		panic(err)
	}
}
