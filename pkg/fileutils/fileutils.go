package fileutils

import (
	"os"
)

func FileExists(filename string) (bool, error) {
	_, err := os.Stat(filename)
	if err == nil {
		return true, nil
	}
	if os.IsExist(err) {
		return true, err
	}

	return false, nil
}
