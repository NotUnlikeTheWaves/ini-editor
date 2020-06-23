package main

import (
	"fmt"
	"io/ioutil"
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
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"pwd":     documentDir,
		})
	})

	r.GET("/nojson", func(c *gin.Context) {
		c.String(200, "hey!")
	})

	r.GET("/filelist", func(c *gin.Context) {
		files, err := ioutil.ReadDir(documentDir)
		if err != nil {
			c.JSON(400, gin.H{
				"msg": err,
			})
		} else {
			c.JSON(200, gin.H{
				"files": files,
			})
		}
	})
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
