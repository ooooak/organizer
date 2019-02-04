package org

import (
	"io"
	"os"
)

// IsEmptyDir directory
func IsEmptyDir(absDir string) bool {
	f, err := os.Open(absDir)
	if err != nil {
		return false
	}

	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true
	}

	return false
}

// IsDir !
func IsDir(absPath string) bool {
	fi, err := os.Stat(absPath)
	return (err == nil && fi.IsDir())
}

// CreateDir !
func CreateDir(absDir string) error {
	return os.Mkdir(absDir, os.ModePerm)
}
