package utils

import (
	"fmt"
	"os"
)

func CreateFile(path string) {
	// check file in path
	_, err := os.Stat(path)

	// create is not exist
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		IsError(err)
		defer file.Close()

		// add initial array for json
		_, err = file.WriteString("[]")
		IsError(err)

		// commit action
		err = file.Sync()
		IsError(err)
	} else {
		fmt.Println("File already there")
	}
}

func JoinPathFile(path, fileName string) string {
	return path + "/" + fileName
}
