package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	fileWasFound, fileName, path := FindFile("casey.txt", "./test")

	if !fileWasFound {
		CreateFile(fileName, path)
	}
}

func FindFile(fileName, path string) (bool, string, string) {
	//root := "/Users/peyton.schlafley/Code/go-repos/go_terminal"
	var fileSystem fs.FS = os.DirFS(path)
	var fileWasFound bool

	var data []string

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
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
		fmt.Printf("file found!\n FileName: %v \n Path: %v \n", fileName, path)
	} else if !fileWasFound {
		fmt.Printf("The file: (%v) was not found in the given directory (%v) \n", fileName, path)
	}
	return fileWasFound, fileName, path
}

func EditFile(fileName, path string) {
	files, err := os.ReadDir(path)

	if files != nil {
		CreateFile(fileName, path)
	} else if err != nil {
		log.Fatal(err)
	}
}

func CreateFile(fileName, path string) {
	if _, pathErr := os.Stat(path); pathErr != nil {
		if os.IsNotExist(pathErr) {
			os.Mkdir(path, 0700)
			os.Create(path + "/" + fileName)
			fmt.Printf("Error: %s\n Creating it now...", pathErr)
		}
	}
	os.Chmod(path, 0700)
	os.Create(path + "/" + fileName)
}
