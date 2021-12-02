package io

import "os"

// FileExists checks if a given filepath is an existing file and not a directory
func FileExists(filepath string) bool {

	fileinfo, err := os.Stat(filepath)

	if os.IsNotExist(err) {
		return false
	}

	return !fileinfo.IsDir()
}
