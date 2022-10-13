package controller

import (
	"gee"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func TextController(c *gee.Context) {
	var json struct {
		Raw string `json:"raw"`
	}
	exe, err := os.Executable()
	if err != nil {
		log.Fatal("find exe dir error, ", err)
	}
	dir := filepath.Dir(exe)
	fileName := uuid.New().String()
	dirName := "uploads"
	uploads := filepath.Join(dir, dirName)
	err = os.MkdirAll(uploads, os.ModePerm)
	if err != nil {
		log.Fatal("mkdir error, ", err)
	}
	fullPath := path.Join(dirName, fileName+".txt")

	err = ioutil.WriteFile(filepath.Join(dir, fullPath), []byte(json.Raw), 0644)
	if err != nil {
		log.Fatal("write txt error, ", err)
	}
	c.JSON(http.StatusOK, gee.H{"url": "/" + fullPath})
}
