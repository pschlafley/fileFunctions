package fileFunctions

import (
	"fmt"
	"log"
	"os"
)

func WriteFile(fileName, path string) {
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
