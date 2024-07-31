package utils

import (
	"fmt"
	"os"
)

func CreateDirIfNotExists(dirPath string) error {
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("Directory creation failed: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("Directory existance check failed: %w", err)
	}
	return nil
}
