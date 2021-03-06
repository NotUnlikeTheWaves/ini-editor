package main

import (
	"fmt"
	"os"
	"path/filepath"
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

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func getDocumentDir() string {
	documentDirName := "documents"
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	documentDir := filepath.Join(path, documentDirName)
	return documentDir
}

func initDocumentDirectory(path string) error {
	// TODO: Expand with permissions check.
	fmt.Printf("INI Editor: Using document directory %s\n", path)

	exists, err := fileExists(path)
	if err != nil {
		fmt.Printf("Failed to init document directory. Error: %s", err)
		os.Exit(1)
	}
	if !exists {
		fmt.Println("Document directory not found, attempting to create...")
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Printf("Failed to created document directory. Error: %s", err)
		}
	}
	fmt.Println("Successfully inited document directory.")
	return nil
}
