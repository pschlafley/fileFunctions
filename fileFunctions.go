package fileFunctions

import (
	"fmt"
	"os"
)

func CreateFile(fileName, path string) {
	if _, pathErr := os.Stat(path); pathErr != nil {
		if os.IsNotExist(pathErr) {
			fmt.Print(pathErr)
			os.Mkdir(path, 0700)
		} else {
			fmt.Print(pathErr)
			os.Chmod(path, 0700)
		}
	}

	os.Create(path + "/" + fileName)
}
