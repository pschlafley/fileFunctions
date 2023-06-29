package fileFunctions

import (
	"fmt"
	"os"
)

func CreateFile(fileName, path string) {
	if _, pathErr := os.Stat(path); pathErr != nil {
		if os.IsNotExist(pathErr) {
			fmt.Print(pathErr)
			os.Mkdir("examples", 0700)
		} else {
			fmt.Print(pathErr)
			os.Chmod("examples", 0700)
		}
	}

	os.Create(path + "/" + fileName)
}
