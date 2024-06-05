package functions

import (
	"fmt"
	"os"
)

func CreateFile(FilePath string, content string) {
	file, err := os.Create(FilePath)
	if err != nil {
		fmt.Println("os error in file ceration")
		os.Exit(1)
	}
	defer file.Close()
	_, er := file.WriteString(content)
	if er != nil {
		fmt.Println("err in writing output content")
		os.Exit(1)
	}
}
