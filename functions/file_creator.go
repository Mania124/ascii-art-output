package functions

import (
	"fmt"
	"os"
)

func CreateFile(FilePath string, content string) error {
	// creating a file
	file, err := os.Create(FilePath)
	if err != nil {
		return fmt.Errorf("error in file creation: %w", err)
	}
	defer file.Close()
	// writing into the file created
	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("error in writing content: %w", err)
	}

	return nil
}
