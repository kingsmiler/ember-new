package main

import (
	"fmt"
	"os"
	"strings"
	"io/ioutil"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:\t ember-new [name]")
		os.Exit(1)
	}

	emberNewHome := getParentDirectory(getCurrentDirectory())
	userWd, _ := os.Getwd()

	fmt.Println(emberNewHome)

	fmt.Println(userWd)

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
