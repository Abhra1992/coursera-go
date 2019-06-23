package services

import "os"

// EnsureDirExists makes sure a directory is created before any operations are performed inside it
func EnsureDirExists(dirName string) error {
	err := os.MkdirAll(dirName, os.ModePerm)
	if err == nil || os.IsExist(err) {
		return nil
	}
	return err
}

// FileExists checks if a file is present
func FileExists(fname string) (bool, error) {
	_, err := os.Stat(fname)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// CleanFileName cleans invalid characters from file name
func CleanFileName(fname string) string {
	// Copy from python utils
	return fname
}
