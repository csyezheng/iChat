package utils

import (
	"os/user"
	"path/filepath"
)

// Returns full path; ~ replaced with actual home directory
func ExpandedFilename(filename string) string {
	if filename == "" {
		panic("filename was empty")
	}

	if len(filename) > 2 && filename[:2] == "~/" {
		if usr, err := user.Current(); err == nil {
			filename = filepath.Join(usr.HomeDir, filename[2:])
		}
	}

	result, err := filepath.Abs(filename)

	if err != nil {
		panic(err)
	}

	return result
}