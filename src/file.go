package main

import (
	"strings"
	"path/filepath"
	"os"
	"log"
	"fmt"
	"io"
)

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}


func getParentDirectory(directory string) string {
	return substr(directory, 0, strings.LastIndex(directory, "/"))
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func copyFile(source string, target string) (err error) {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourceFile.Close()

	targetFile, err := os.Create(target)
	if err != nil {
		return err
	}

	defer targetFile.Close()

	_, err = io.Copy(targetFile, sourceFile)
	if err == nil {
		sourceInfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(target, sourceInfo.Mode())
		}

	}

	return
}

func copyDir(source string, target string) (err error) {

	// get properties of source dir
	sourceInfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir

	err = os.MkdirAll(target, sourceInfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {
		sourceFilePointer := source + "/" + obj.Name()
		destinationFilePointer := target + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = copyDir(sourceFilePointer, destinationFilePointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform copy
			err = copyFile(sourceFilePointer, destinationFilePointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

