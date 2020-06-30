package main

import (
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func apiFileList(c *gin.Context) {
	files, err := ioutil.ReadDir(documentDir)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err,
		})
	} else {
		c.JSON(200, gin.H{
			"files": createFileList(files),
		})
	}
}

func apiReadFile(c *gin.Context) {
	fileName := c.Param("path")
	filePath := filepath.Join(documentDir, fileName) + ".ini"
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		c.String(http.StatusNotFound, "No file found named %s.", filePath)
	} else {
		c.String(http.StatusOK, string(file))
	}
}
