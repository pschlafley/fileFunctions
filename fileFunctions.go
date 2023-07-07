package fileFunctions

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func FindFile(fileName, path string) {
	//root := "/Users/peyton.schlafley/Code/go-repos/go_terminal"
	var fileSystem fs.FS = os.DirFS(path)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(path)

		if d.Name() == fileName {
			fmt.Printf("Found the file named: %v \n Path: %v", fileName, path)
		}

		return nil
	})
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
