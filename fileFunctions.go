package fileFunctions

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

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

func CreateFile(fileName, path string) {
	if _, pathErr := os.Stat(path); pathErr != nil {
		if os.IsNotExist(pathErr) {
			fmt.Printf("Error: %s", pathErr)
		}
	} else {
		os.Chmod(path, 0700)
		os.Create(path + "/" + fileName)
	}
}

func CreateDirectory(path string) {
	if _, pathErr := os.Stat(path); pathErr != nil {
		if os.IsNotExist(pathErr) {
			fmt.Printf("Error: %s", pathErr)
			os.Mkdir(path, 0700)
		}
	}
}
