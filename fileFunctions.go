package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

func main() {
	fileWasFound, fileName, path, _ := FindFile("test.txt", "./test")

	if fileWasFound {
		DeleteFile(fileName, path)
	}
}

func FindFile(fileName, path string) (bool, string, string, []string) {
	var fileSystem fs.FS = os.DirFS(path)
	var fileWasFound bool

	var data []string
	var errors []string

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			_, after, _ := strings.Cut(err.Error(), ".:")

			errors = append(errors, after)

			return nil
		}

		data = append(data, d.Name())

		return nil
	})

	for i := 0; i < len(data); i++ {
		if data[i] == fileName {
			fileWasFound = true
		} else if data[i] != fileName {
			fileWasFound = false
		}
	}

	if fileWasFound {
		fmt.Printf("%v was found at %v", fileName, path)
	}

	return fileWasFound, fileName, path, errors
}

func EditFile(fileName, path string) {
	files, err := os.ReadDir(path)

	if files != nil {
		CreateFile(fileName, path)
	} else if err != nil {
		log.Fatal(err)
	}
}

func CreateFile(fileName, path string) (string, string, error) {
	if _, pathErr := os.Stat(path); pathErr != nil {
		if os.IsNotExist(pathErr) {
			return fileName, path, pathErr
		}
	}

	os.Chmod(path, 0700)
	os.Create(path + "/" + fileName)
	return fileName, path, nil
}

func CreateDirectory(path string) (string, error) {
	if _, pathErr := os.Stat(path); pathErr != nil {
		if os.IsNotExist(pathErr) {
			os.Mkdir(path, 0700)
		}
		return path, pathErr
	}
	return path, nil
}

func DeleteFile(fileName, path string) (string, error) {
	os.Remove(path + "/" + fileName)
	return "", nil
}
