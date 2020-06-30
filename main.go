package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const documentDirName string = "documents"

var documentDir string

func main() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	documentDir = filepath.Join(path, documentDirName)
	err = initDocumentDirectory(documentDir)
	if err != nil {
		fmt.Printf("Failed to init document directory. Error: %s", err)
		os.Exit(1)
	}
	fmt.Println("Successfully inited document directory.")

	// Set up Gin stuff
	r := gin.Default()

	r.GET("/api/v1/filelist", apiFileList)
	r.GET("/api/v1/file/:path")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func initDocumentDirectory(path string) error {
	// TODO: Expand with permissions check.
	fmt.Printf("INI Editor: Using document directory %s\n", path)

	exists, err := exists(path)
	if err != nil {
		return err
	}
	if exists == false {
		fmt.Println("Document directory not found, attempting to create...")
		err := os.Mkdir(path, os.ModePerm)
		return err
	}
	return nil
}
