package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func apiFileList(c *gin.Context) {
	files, err := ioutil.ReadDir(documentDir)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err,
		})
	} else {
		fmt.Println(FileEntry{"hey", 2})
		c.JSON(200, gin.H{
			"files": createFileList(files),
		})
	}
}
