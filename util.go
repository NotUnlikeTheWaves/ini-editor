package main

import (
	"os"
)

func createFileList(files []os.FileInfo) []FileEntry {
	size := len(files)
	entryList := make([]FileEntry, size)
	for i, v := range files {
		entryList[i] = FileEntry{
			Path:    v.Name(),
			Size:    v.Size(),
			Lastmod: v.ModTime(),
		}
	}
	return entryList
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
