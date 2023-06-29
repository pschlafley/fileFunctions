package fileFunctions

import (
	"fmt"
	"os"
)

func CreateFile(fileName, path string) {
	if _, pathErr := os.Stat(path); pathErr != nil {
		if os.IsNotExist(pathErr) {
			os.Mkdir(path, 0700)
			fmt.Print(pathErr)
		} else {
			os.Chmod(path, 0700)
			fmt.Print(pathErr)
		}
	}

	os.Create(path + "/" + fileName)
}
