package functions

import (
	"fmt"
	"os"
)

func CreateFile(FilePath string, content string) error {
	file, err := os.Create(FilePath)
	if err != nil {
		return fmt.Errorf("error in file creation: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("error in writing content: %w", err)
	}

	return nil
}
