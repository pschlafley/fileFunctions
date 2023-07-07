package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	FindFile2("text.txt", ".")
}

func FindFile2(fileName, path string) {
	//root := "/Users/peyton.schlafley/Code/go-repos/go_terminal"
	var fileSystem fs.FS = os.DirFS(path)

	var data []fs.DirEntry

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		data = append(data, d)

		return nil
	})

	for i := 0; i < len(data); i++ {
		if data[i].Name() == fileName {
			fmt.Printf("fileName: %v \n data[i]: %v", fileName, data[i].Name())
		} else if data[i].Name() != fileName {
			fmt.Printf("fileName not found")
		}
	}
}

func FindFile(fileName, path string) (string, string) {
	//root := "/Users/peyton.schlafley/Code/go-repos/go_terminal"
	var fileSystem fs.FS = os.DirFS(path)

	var data []fs.DirEntry

	var fileWasFound bool

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		data = append(data, d)

		if d.Name() == fileName {
			fmt.Printf("Found the file named: %v \n Path: %v", fileName, path)
			fileWasFound = true
		} else {
			fmt.Printf("The file: %v was not found in this path: %v", fileName, path)
			fileWasFound = false
		}

		return nil
	})

	print(data)

	if fileWasFound {
		return fileName, path
	} else {
		return fmt.Sprintf("FileName: %v was not found", fileName), ""
	}

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
		} else {
			os.Chmod(path, 0700)
			os.Create(path + "/" + fileName)
			fmt.Print(pathErr)
		}
	}
}
